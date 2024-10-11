package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/zkryaev/taskwb-L0/models"
)

func AddDelivery(tx *sql.Tx, delivery models.Delivery, OrderUID string) error {
	query := `INSERT INTO deliveries ("name", "phone", "zip", "city", "address", "region", "email", "order_uid") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := tx.Exec(
		query,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
		OrderUID,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetDelivery(db *sql.DB, OrderUID string) (*models.Delivery, error) {
	query := "SELECT * FROM deliveries WHERE order_uid = $1"

	row := db.QueryRow(query, OrderUID)

	var delivery models.Delivery
	var uid string
	err := row.Scan(
		&uid,
		&delivery.Name,
		&delivery.Phone,
		&delivery.Zip,
		&delivery.City,
		&delivery.Address,
		&delivery.Region,
		&delivery.Email,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("get delivery failed: %w", err)
		}
		return nil, fmt.Errorf("get delivery failed: %w", err)
	}

	return &delivery, nil
}

/*
func UpdateDelivery(db *sql.DB, delivery models.Delivery, OrderUID string) error {

}*/
