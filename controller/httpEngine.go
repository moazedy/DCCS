package controller

import (
	"DCCS/domain/datastore"
	"DCCS/logic"
	"DCCS/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func Run(port string) {

	session, err := datastore.NewCouchbaseSession()
	if err != nil {
		log.Fatal(err)
	}

	contractRepo := repository.NewContractRepo(session)
	// comparingObjectRepo := repository.NewComparingObjectRepo(session)

	contractLogic := logic.NewContractLogic(contractRepo)
	contractController := NewContractController(contractLogic)

	engine := gin.Default()
	con := engine.Group("contract")
	con.POST("/save", contractController.Save)
	con.GET("/read/id/:id", contractController.ReadById)
	con.GET("read/title/:title", contractController.ReadByTitle)
	con.POST("/search/exact", contractController.ReadByIntValue)
	con.POST("/search/similars", contractController.GetContractSimilars)
	con.POST("/search/main", contractController.ContractMainSearch)

	engine.Run(":" + port)

}
