package main

import "RestauMana/database"

func main() {
	database.Initdb()
	Router()
}
