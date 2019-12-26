package alog

import "fmt"

type passmock struct {
	msg chan []byte
}

func (pm *passmock) Write(p []byte) (n int, err error) {

	n = len(p)
	fmt.Println("IN")
	pm.msg <- p

	fmt.Println("OUT")

	return n, err
}

type writemock struct{}

func (w *writemock) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))

	return n, err
}
