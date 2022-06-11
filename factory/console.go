package factory

import "fmt"

type Console struct {
}

func (c *Console) Init() {}

func (c *Console) WriteString(d string) {
	fmt.Print(d)
}

func (c *Console) WriteNewLine() {
	fmt.Println()
}
