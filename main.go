package main

import (
	"fmt"

	"github.com/pedroluis02/design-patterns-golang-s1/builder"
	"github.com/pedroluis02/design-patterns-golang-s1/singleton"
)

func main() {
	builder1()
}

func singleton1() {
	connection := singleton.GetDatabaseConnection()
	fmt.Println(fmt.Sprintf("DB Status: %s", connection.GetStatus()))
	connection.Execute("select CURDATE()")
	connection.Disconnect()
}

func builder1() {
	product := &builder.Product{
		Id:    "P1-000-111-222-333-444",
		Name:  "Macbook pro 2019 Intel",
		Price: 3500,
		Stock: 10,
	}

	jsonBuilder := &builder.ProductJsonBuilder{}
	xmlBuilder := &builder.ProductXmlBuilder{}
	response := &builder.ProductResponse{}

	jsonStr := response.SetBuilder(jsonBuilder).BuildResponse(product)
	fmt.Println(fmt.Sprintf("json: %s", jsonStr))

	xmlStr := response.SetBuilder(xmlBuilder).BuildResponse(product)
	fmt.Println(fmt.Sprintf("xml: %s", xmlStr))
}
