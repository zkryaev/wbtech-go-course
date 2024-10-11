package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/zkryaev/taskwb-L0/models"
	"github.com/zkryaev/taskwb-L0/repository/config"
	"github.com/zkryaev/taskwb-L0/repository/database"
)

type OrdersRepo struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*OrdersRepo, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	ordersDB := &OrdersRepo{
		DB: db,
	}
	return ordersDB, nil
}

func (o *OrdersRepo) AddOrder(order models.Order) error {
	tx, err := o.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	query := `INSERT INTO orders("order_uid", "track_number","entry", "locale", "internal_signature", "customer_id", "delivery_service", "shardkey", "sm_id", "date_created", "oof_shard") 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err = o.DB.Exec(
		query,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert order: %w", err)
	}

	err = database.AddPayment(tx, order.Payment, order.OrderUID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert payment: %w", err)
	}

	err = database.AddItems(tx, order.Items, order.OrderUID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert items: %w", err)
	}

	err = database.AddDelivery(tx, order.Delivery, order.OrderUID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert delivery: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (o *OrdersRepo) GetOrder(OrderUID string) (*models.Order, error) {
	query := "SELECT * FROM orders WHERE order_uid = $1"
	row := o.DB.QueryRow(query, OrderUID)
	var order models.Order
	err := row.Scan(
		&order.OrderUID,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerID,
		&order.DeliveryService,
		&order.Shardkey,
		&order.SmID,
		&order.DateCreated,
		&order.OofShard,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}
	delivery, err := database.GetDelivery(o.DB, OrderUID)
	if err != nil {
		return nil, err
	}
	order.Delivery = *delivery

	payment, err := database.GetPayment(o.DB, OrderUID)
	if err != nil {
		return nil, err
	}
	order.Payment = *payment

	items, err := database.GetItems(o.DB, OrderUID)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return &order, nil
}

func (o *OrdersRepo) GetOrders() ([]models.Order, error) {
	query := "SELECT * FROM orders"
	rows, err := o.DB.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShard)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order row: %w", err)
		}

		delivery, err := database.GetDelivery(o.DB, order.OrderUID)
		if err != nil {
			return nil, fmt.Errorf("failed to get delivery for order %s: %w", order.OrderUID, err)
		}
		order.Delivery = *delivery

		payment, err := database.GetPayment(o.DB, order.OrderUID)
		if err != nil {
			return nil, fmt.Errorf("failed to get payment for order %s: %w", order.OrderUID, err)
		}
		order.Payment = *payment

		items, err := database.GetItems(o.DB, order.OrderUID)
		if err != nil {
			return nil, fmt.Errorf("failed to get items for order %s: %w", order.OrderUID, err)
		}
		order.Items = items

		orders = append(orders, order)
	}

	return orders, nil
}
