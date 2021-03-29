package utils

import (
	"DCCS/constants"
	"DCCS/domain/datastore"
	"DCCS/domain/models"
	"DCCS/repository"
	"sort"
)

func CompareReceivedSerialData(received string) ([]models.CompareResult, error) {
	session, err := datastore.NewCouchbaseSession()
	if err != nil {
		return nil, err
	}
	allComparingObjects, err := repository.NewComparingObjectRepo(session).GetAllComparingObjects()
	if err != nil {
		return nil, err
	}

	result := make([]models.CompareResult, len(allComparingObjects))

	for i, cObject := range allComparingObjects {
		var percent float64 = 0
		for k := 0; k < len(received); k++ {
			if received[k] == cObject.SerialValue[k] {
				percent = percent + constants.EachTermPercent
			}

		}
		result[i] = models.CompareResult{
			EqualityPercent: percent,
			ContractId:      cObject.ContractId,
			Title:           cObject.Title,
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].EqualityPercent < result[j].EqualityPercent
	})

	return result, nil
}
