package queries

import "DCCS/constants"

const (
	SaveContractQuery = `INSERT INTO ` + constants.ContractsBucket + ` (KEY,VALUE) VALUES ($id, $contract) `

	ReadContractByIdQuery = `SELECT ` + constants.ContractsBucket + `.*  FROM  ` + constants.ContractsBucket +
		` WHERE id=$id `

	ReadContractByTitleQuery = `SELECT ` + constants.ContractsBucket + `.*  FROM  ` + constants.ContractsBucket +
		` WHERE title=$title `
)
