package main

import (
	"DCCS/constants"
	"DCCS/controller"
	"DCCS/domain/datastore"
)

func main() {

	datastore.CouchbasePing()

	controller.Run(constants.HttpServingPort)

}
