package main

import (
	"fmt"
	"testing"

	"github.com/pedroluis02/design-patterns-golang-s1/builder"
	"github.com/pedroluis02/design-patterns-golang-s1/singleton"
)

func TestSingleton(t *testing.T) {
	connection := singleton.GetDatabaseConnection()
	status := connection.GetStatus()
	want := singleton.Success

	if status != want {
		t.Errorf("got %s, wanted %s", status, want)
	}
}

func TestBuilder(t *testing.T) {
	product := &builder.Product{
		Id:    "1",
		Name:  "Macbook pro",
		Price: 3500,
		Stock: 5,
	}

	jsonBuilder := &builder.ProductJsonBuilder{}
	response := &builder.ProductResponse{}

	jsonStr := response.SetBuilder(jsonBuilder).BuildResponse(product)
	wantJson := fmt.Sprintf("{\"id\":\"%s\",\"name\":\"%s\",\"price\":%g,\"stock\":%d}", product.Id, product.Name, product.Price, product.Stock)

	fmt.Println(jsonStr)
	fmt.Println(wantJson)

	if jsonStr != wantJson {
		t.Errorf("json result error: %s", wantJson)
	}
}
