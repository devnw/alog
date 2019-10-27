package alog

type writemock struct {
	msg chan []byte
}

func (w *writemock) Write(p []byte) (n int, err error) {
	w.msg <- p

	return n, err
}
