package queries

import "DCCS/constants"

const (
	GetAllComparingObjectsQuery1 = ` Select ` + constants.ContractsBucket + `.id, ` +
		constants.ContractsBucket + `.title, ` + constants.ContractsBucket + `.serial_value, ` +
		constants.ContractsBucket + `.int_value   FROM ` + constants.ContractsBucket

	GetAllComparingObjectsQuery = `select contracts.id, contracts.title, contracts.serial_value, contracts.int_value from contracts`
)
