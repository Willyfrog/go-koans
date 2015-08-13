package go_koans

import (
	"bytes"
	"fmt"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/
		var (
			n   int
			err error
			p   []byte
		)
		p = make([]byte, 8)
		n = 1
		for n > 0 {
			n, err = in.Read(p)
			if err == nil && n > 0 {
				n, err = out.Write(p[:n])
			}
			if err != nil {
				if err.Error() != "EOF" {
					fmt.Printf("error: %v", err)
				}
				n = -1
			}

		}
		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		var (
			err   error
			p     []byte
			total int
		)
		p = make([]byte, 1)
		n := 1
		for n > 0 && total < 5 {
			n, err = in.Read(p)
			if err == nil && n > 0 {
				n, err = out.Write(p[:n])
			}
			total += n
			if err != nil {
				if err.Error() != "EOF" {
					fmt.Printf("error: %v", err)
				}
				n = -1
			}
		}
		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
