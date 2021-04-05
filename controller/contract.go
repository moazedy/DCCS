package controller

import (
	"DCCS/domain/models"
	"DCCS/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContractController interface {
	Save(ctx *gin.Context)
	ReadById(ctx *gin.Context)
	ReadByTitle(ctx *gin.Context)
	ReadByIntValue(ctx *gin.Context)
	//CheckContractExistance(ctx *gin.Context)
	GetContractSimilars(ctx *gin.Context)
	ContractMainSearch(ctx *gin.Context)
}

type contract struct {
	logic logic.ContractLogic
}

func NewContractController(logic logic.ContractLogic) ContractController {
	c := new(contract)
	c.logic = logic

	return c
}

func (c *contract) Save(ctx *gin.Context) {
	var request models.Contract

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message:": "seems you have sent imperfect request"})
		return
	}

	err = c.logic.Save(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message:": "error on saving contract",
			"error ": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message:": "contract saved to db"})
}

func (c *contract) ReadById(ctx *gin.Context) {
	Id := ctx.Param("id")
	con, err := c.logic.ReadById(Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message:": "error on reading contract"})
		return
	}

	ctx.JSON(http.StatusOK, con)
}

func (c *contract) ReadByTitle(ctx *gin.Context) {
	title := ctx.Param("title")
	con, err := c.logic.ReadByTitle(title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message:": "error on reading contract", "error: ": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, con)
}

func (c *contract) GetContractSimilars(ctx *gin.Context) {
	var request models.ContractSerialDataTemplate

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message:": "seems you have sent imperfect request"})
		return
	}

	res, err := c.logic.GetSimilarContracts(request, 3) // TODO : result size should be received from client.
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message:": "error on system functions",
			"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *contract) ReadByIntValue(ctx *gin.Context) {
	var request models.ContractSerialDataTemplate

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message:": "seems you have sent imperfect request"})
		return
	}

	res, err := c.logic.ReadByIntValue(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message:": "error on system functions",
			"error": err.Error()})
		return
	}

	if res.Title != "" {
		ctx.JSON(http.StatusOK, res)
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "no matches found !"})

}

func (c *contract) ContractMainSearch(ctx *gin.Context) {
	var request models.ContractSerialDataTemplate

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message:": "seems you have sent imperfect request"})
		return
	}

	exact, err := c.logic.ReadByIntValue(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message:": "error on system functions",
			"error": err.Error()})
		return
	}
	if exact != nil && exact.Title != "" {
		resultMap := make(map[string]interface{})
		resultMap["title"] = exact.Title
		resultMap["id"] = exact.Id
		resultMap["int value"] = exact.IntValue
		resultMap["serial value"] = exact.SerialValue
		ctx.JSON(http.StatusOK, gin.H{"your sent ditails are match with this type of contract: ": resultMap})
		return
	}

	similars, err := c.logic.GetSimilarContracts(request, 3)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message:": "error on system functions",
			"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"there are no exact match contracts with your sent ditails, but there are similars : ": similars})

}
