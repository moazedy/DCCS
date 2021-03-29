package models

import "github.com/google/uuid"

type ComparingObject struct {
	ContractId  uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	SerialValue string    `json:"serial_value"`
	IntValue    int64     `json:"int_value"`
}

type CompareResult struct {
	EqualityPercent float64   `json:"equality_percent"`
	Title           string    `json:"title"`
	ContractId      uuid.UUID `json:"id"`
}
