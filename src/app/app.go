package app

import "fmt"

func init() {
	ReadSettings()
	ConnectDb()
	fmt.Println("Connected to DB")
}
