package models

import "time"

type Order struct {
	OrderID     int       `json:"order_id"`
	Product     Product   `json:"item"`
	Status      string    `json:"status"`
	DateOfOrder time.Time `json:"date_of_order"`
	User        User      `json:"user"`
}
