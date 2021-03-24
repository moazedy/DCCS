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

	return &con, nil
}
