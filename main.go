package main

import (
	"test/wex/database"
	"test/wex/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
