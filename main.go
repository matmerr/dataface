package dataface

import (
	"fmt"
)

func main() {

	// Redis
	db, _ := InitializeDatabase("redis", "localhost", 6379, "", "")
	db.Put("redis test", []byte("redis response"))
	response, _ := db.Get("redis response")
	fmt.Println(string(response[:])) // redis reponse

	// MongoDB
	db, _ = InitializeDatabase("mongo", "localhost", 27017, "", "")
	db.Put("mongo test", []byte("mongo response"))
	response, _ = db.Get("mongo test")
	fmt.Println(string(response[:])) // mongo response
}
