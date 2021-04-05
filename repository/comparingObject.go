package repository

import (
	"DCCS/domain/models"
	"DCCS/repository/queries"
	"log"

	"github.com/couchbase/gocb/v2"
)

type ComparingObjectRepo interface {
	GetAllComparingObjects() ([]models.ComparingObject, error)
}

type comparingObject struct {
	session *gocb.Cluster
}

func NewComparingObjectRepo(session *gocb.Cluster) ComparingObjectRepo {
	c := new(comparingObject)
	c.session = session

	return c
}

func (c comparingObject) GetAllComparingObjects() ([]models.ComparingObject, error) {
	res, err := c.session.Query(
		queries.GetAllComparingObjectsQuery,
		nil,
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var cObjects []models.ComparingObject
	for res.Next() {
		var cObject models.ComparingObject
		err = res.Row(&cObject)
		if err != nil {
			if err == gocb.ErrNoResult {
				return cObjects, nil
			}

			log.Println(err.Error())
			return nil, err
		}

		cObjects = append(cObjects, cObject)
	}

	return cObjects, nil

}
