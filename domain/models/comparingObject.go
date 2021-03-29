package models

import "github.com/google/uuid"

type ComparingObject struct {
	ContractId  uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	SerialValue string    `json:"serial_value"`
	IntValue    int64     `json:"int_value"`
}
