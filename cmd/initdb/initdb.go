package main

import "fmt"

// TODO: 修正
func main() {

	db, err := database.NewMyPostgreSQLDB(logger, true)
	if err != nil {
		fmt.Println("error:", err)
	}

	if err := seed.CreateDevSeedData(db); err != nil {
		fmt.Println("error:", err)
	}

}