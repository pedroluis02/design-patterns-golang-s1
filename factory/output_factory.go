package factory

type OutputStreamType int

const (
	FileType OutputStreamType = iota
	ConsoleType
	MemoryType
)

type OutputStreamFactory struct {
}

func (c *OutputStreamFactory) GetOutput(t OutputStreamType) OutputStream {
	var output OutputStream
	switch t {
	case FileType:
		output = &File{}
	case ConsoleType:
		output = &Console{}
	case MemoryType:
		output = &Memory{}
	default:
		output = nil
	}

	if output != nil {
		output.Init()
	}

	return output
}
