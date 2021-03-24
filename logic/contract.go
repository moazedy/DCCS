package logic

import (
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

	err := c.repo.Save(con)
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
