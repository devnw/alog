package alog

import "fmt"

type writemock struct {
	//msg chan []byte
}

func (w *writemock) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))

	return n, err
}
