package main

import (
	"github.com/Ckaiique/go-rest-api.git/database"
	"github.com/Ckaiique/go-rest-api.git/models"
	"github.com/Ckaiique/go-rest-api.git/routes"
)


func main() {
	
	models.Personalidades = []models.Personalidade{
		{Id:1, Nome:"Nome 1",Historia: "Historia 1"},
		{Id:2, Nome:"Nome 2",Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}