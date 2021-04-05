package repository

import (
	"DCCS/domain/models"
	"DCCS/repository/queries"

	"errors"
	"log"

	"github.com/couchbase/gocb/v2"
)

type ContractRepo interface {
	Save(con models.Contract) error
	ReadById(id string) (*models.Contract, error)
	ReadByTitle(title string) (*models.Contract, error)
	ReadByIntValue(intValue int64) (*models.Contract, error)
	ContractExists(title string) (*bool, error)
}

type contract struct {
	session *gocb.Cluster
}

func NewContractRepo(session *gocb.Cluster) ContractRepo {
	c := new(contract)
	c.session = session

	return c
}

func (c *contract) Save(con models.Contract) error {
	_, err := c.session.Query(
		queries.SaveContractQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"id":       con.Id,
			"contract": con,
		}},
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (c *contract) ReadById(id string) (*models.Contract, error) {
	res, err := c.session.Query(
		queries.ReadContractByIdQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"id": id,
		}},
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var con models.Contract
	for res.Next() {
		err = res.Row(&con)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("contract does not exist !")
			}

			log.Println(err.Error())
			return nil, err
		}
	}

	if con.Title == "" {
		return nil, errors.New("contract does not exist !")
	}
	return &con, nil
}

func (c *contract) ReadByTitle(title string) (*models.Contract, error) {
	res, err := c.session.Query(
		queries.ReadContractByTitleQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"title": title,
		}},
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var con models.Contract
	for res.Next() {
		err = res.Row(&con)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("contract does not exist !")
			}

			log.Println(err.Error())
			return nil, err
		}
	}

	if con.Title == "" {
		return nil, errors.New("contract does not exist !")
	}

	return &con, nil
}

func (c *contract) ReadByIntValue(intValue int64) (*models.Contract, error) {
	res, err := c.session.Query(
		queries.ReadContractByIntValue,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"intval": intValue,
		}},
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var con models.Contract
	for res.Next() {
		err = res.Row(&con)
		if err != nil {
			if err == gocb.ErrNoResult {
				return nil, errors.New("contract does not exist !")
			}

			log.Println(err.Error())
			return nil, err
		}
	}

	return &con, nil

}

func (c *contract) ContractExists(title string) (*bool, error) {
	var boolResult bool
	res, err := c.session.Query(
		queries.ReadContractByTitleQuery,
		&gocb.QueryOptions{NamedParameters: map[string]interface{}{
			"title": title,
		}},
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var con models.Contract
	for res.Next() {
		err = res.Row(&con)
		if err != nil {
			if err == gocb.ErrNoResult {
				return &boolResult, errors.New("contract does not exist !")
			}

			log.Println(err.Error())
			return nil, err
		}
	}

	if con.Title == "" {

		return &boolResult, nil
	}
	boolResult = true

	return &boolResult, nil
}
