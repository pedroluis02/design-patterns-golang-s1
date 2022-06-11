package main

import (
	"fmt"

	"github.com/pedroluis02/design-patterns-golang-s1/singleton"
)

func main() {
	testSingleton()
}

func testSingleton() {
	connection := singleton.GetDatabaseConnection()
	fmt.Println(fmt.Sprintf("DB Status: %s", connection.GetStatus()))
	connection.Execute("select CURDATE()")
	connection.Disconnect()
}
