package logic

import (
	"DCCS/application/utils"
	"DCCS/constants"
	"DCCS/domain/models"
	"DCCS/repository"
	"errors"

	"github.com/google/uuid"
)

type ContractLogic interface {
	Save(con models.Contract) error
	ReadById(id string) (*models.Contract, error)
	ReadByTitle(title string) (*models.Contract, error)
	ReadByIntValue(con models.ContractSerialDataTemplate) (*models.Contract, error)
	GetSimilarContracts(con models.ContractSerialDataTemplate, resultSize int) ([]models.CompareResult, error)
}

type contract struct {
	repo repository.ContractRepo
}

func NewContractLogic(repo repository.ContractRepo) ContractLogic {
	c := new(contract)
	c.repo = repo

	return c
}

func (c *contract) Save(con models.Contract) error {
	con.Id = uuid.New()
	serialValue := utils.ContractToSerialCode(con)
	con.SerialValue = serialValue
	ContractIntValue, err := utils.SerialCodeToIntValue(serialValue)
	if err != nil {
		return err
	}
	con.IntValue = *ContractIntValue

	err = c.repo.Save(con)
	if err != nil {
		return err
	}

	return nil
}

func (c *contract) ReadById(id string) (*models.Contract, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New(constants.InvalidIdValue)
	}

	con, err := c.repo.ReadById(id)
	if err != nil {
		return nil, err
	}

	return con, nil
}

func (c *contract) ReadByTitle(title string) (*models.Contract, error) {
	if title == "" {
		return nil, errors.New(constants.InvalidTitleValue)
	}

	con, err := c.repo.ReadByTitle(title)
	if err != nil {
		return nil, err
	}

	return con, nil
}

func (c contract) ReadByIntValue(con models.ContractSerialDataTemplate) (*models.Contract, error) {
	requestSerialValue := utils.TemplateToSerialCode(con)
	requestIntValue, err := utils.SerialCodeToIntValue(requestSerialValue)
	if err != nil {
		return nil, err
	}

	res, err := c.repo.ReadByIntValue(*requestIntValue)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c contract) GetSimilarContracts(con models.ContractSerialDataTemplate, resultSize int) ([]models.CompareResult, error) {
	requestSerialValue := utils.TemplateToSerialCode(con)
	primaryResult, err := utils.CompareReceivedSerialData(requestSerialValue)
	if err != nil {
		return nil, err
	}

	result := make([]models.CompareResult, resultSize)
	lenth := len(primaryResult)
	for i := 1; i <= resultSize; i++ {
		result[i-1] = primaryResult[lenth-i]
	}

	return result, nil
}
