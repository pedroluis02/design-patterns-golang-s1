package factory

import (
	"fmt"
	"log"
	"os"
	"time"
)

type File struct {
	filename string
	status   bool
}

func (f *File) writeData(d string) {
	fo, err := os.OpenFile(f.filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	} else {
		n, err := fo.Write([]byte(d))
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(fmt.Sprintf("write on file=%s: OK, data-len=%d", f.filename, n))
		}
	}

	defer fo.Close()
}

func (f *File) Init() {
	f.filename = "log-tmp.txt"
	fo, err := os.OpenFile(f.filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		f.status = false
		log.Fatal(err)
	} else {
		f.status = true
		fmt.Println(fmt.Sprintf("file=%s created: OK", f.filename))

		currentTime := time.Now()
		str := currentTime.String() + "\n"

		fo.WriteString(str)
	}

	defer fo.Close()
}

func (f *File) WriteString(d string) {
	f.writeData(d)
}

func (f *File) WriteNewLine() {
	f.writeData("\n")
}
