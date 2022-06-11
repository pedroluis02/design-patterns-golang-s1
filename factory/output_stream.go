package factory

type OutputStream interface {
	Init()
	WriteString(d string)
	WriteNewLine()
}
