package main

import (
	"fmt"

	"github.com/pedroluis02/design-patterns-golang-s1/abstract_factory"
	"github.com/pedroluis02/design-patterns-golang-s1/builder"
	"github.com/pedroluis02/design-patterns-golang-s1/factory"
	"github.com/pedroluis02/design-patterns-golang-s1/observer"
	"github.com/pedroluis02/design-patterns-golang-s1/singleton"
)

func main() {
	observer1()
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

func factory1() {
	outputFactory := &factory.OutputStreamFactory{}
	output := outputFactory.GetOutput(factory.FileType)

	output.WriteString("Message 1")
	output.WriteNewLine()
	output.WriteString("Message 2")
}

func abstractFactory() {
	factoryMaker := abstract_factory.ServerFactoryMaker{}
	factory := factoryMaker.GetFactory(abstract_factory.Testing)

	socketServer := factory.CreateSocketServer()
	apiRestFul := factory.CreateApiRestFul()

	fmt.Println(socketServer.ToString())
	fmt.Println(apiRestFul.ToString())
}

func observer1() {
	product := observer.NewProduct("Macbook pro")
	fmt.Println(fmt.Sprintf("Product: %s", product.GetName()))

	observer.NewProductTax(product)
	observer.NewProductDiscount(product, 10)

	product.SetPrice(4000)
	product.SetPrice(2500)
}
