package factory

import "container/list"

type Memory struct {
	data *list.List
}

func (m *Memory) Init() {
	if m.data == nil {
		m.data = list.New()
	}
}

func (m *Memory) WriteString(d string) {
	m.data.PushBack(&list.Element{
		Value: d,
	})
}

func (m *Memory) WriteNewLine() {
	m.data.PushBack(&list.Element{
		Value: "\n",
	})
}
