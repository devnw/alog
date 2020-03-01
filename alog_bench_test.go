package alog

import "testing"

func Benchmark_Println(b *testing.B) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Println(nil)
		if _, ok := <-mock.msg; !ok {
			return
		}
	}
}

func Benchmark_Printf(b *testing.B) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Println("%s", "hello")
		if _, ok := <-mock.msg; !ok {
			return
		}
	}
}

func Benchmark_Print(b *testing.B) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Print(nil)
		if _, ok := <-mock.msg; !ok {
			return
		}
	}
}

func Benchmark_Debugln(b *testing.B) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Debugln(nil, nil)
		if _, ok := <-mock.msg; !ok {
			return
		}
	}
}

func BenchmarkDebugf(b *testing.B) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Debugln(nil, "%s", "hello")
		if _, ok := <-mock.msg; !ok {
			return
		}
	}
}

func Benchmark_Debug(b *testing.B) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {

		Debug(nil, nil)
		if _, ok := <-mock.msg; !ok {
			return
		}
	}
}
