package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/pedroluis02/design-patterns-golang-s1/builder"
	"github.com/pedroluis02/design-patterns-golang-s1/factory"
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

	t.Log(jsonStr)
	t.Log(wantJson)

	if jsonStr != wantJson {
		t.Errorf("json result error: %s", wantJson)
	}
}

func TestFactory(t *testing.T) {
	outputFactory := &factory.OutputStreamFactory{}
	outputFactory.GetOutput(factory.FileType)

	want := "log-tmp.txt"
	_, err := os.Stat(want)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Logf("%s exists", want)
	}
}
