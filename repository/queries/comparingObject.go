package queries

import "DCCS/constants"

const (
	GetAllComparingObjectsQuery = ` Select ` + constants.ContractsBucket + `.id, ` +
		constants.ContractsBucket + `.title, ` + constants.ContractsBucket + `.serial_value, ` +
		constants.ContractsBucket + `.int_value   FROM ` + constants.ContractsBucket
)
