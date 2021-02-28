package alog

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func stdmock(t *testing.T) *passmock {
	mock := &passmock{make(chan []byte)}

	err := testg(nil, mock)
	if err != nil {
		t.Fatal(err)
	}

	return mock
}

func stdcheck(t *testing.T, mock *passmock, expected string) {
	log, ok := <-mock.msg
	if !ok {
		t.Fatal("closed channel")
	}

	err := check(log, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_alog_global_defaults(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		lvls,
		STD,
		mock,
	}

	err := Global(
		nil,
		"PREFIX",
		"",
		nil,
		-1,
		dest,
	)

	if err != nil {
		t.Fatal(err)
	}

	for _, test := range prefixlogs {
		if test.lvl&INFO > 0 {
			Println(test.text)
		} else if test.lvl&CUSTOM > 0 {
			Customln("CUSTOM", test.err, test.text)
		} else {
			stdlns[test.lvl](test.err, test.text)
		}

		stdcheck(t, mock, test.expected)
	}

	Close()
}

func Test_alog_Global(t *testing.T) {
	err := Global(
		context.Background(),
		"PREFIX",
		"",
		nil,
		-1,
	)

	if err == nil {
		t.Fatal("expected error but got success")
	}
}

func Test_alog_setGlobal(t *testing.T) {
	err := setGlobal(nil)

	if err == nil {
		t.Fatal("expected error but got success")
	}
}

func Test_alog_ln(t *testing.T) {

	mock := stdmock(t)
	for _, test := range logs {
		if test.lvl&INFO > 0 {
			Println(test.text)
		} else if test.lvl&CUSTOM > 0 {
			Customln("CUSTOM", test.err, test.text)
		} else {
			stdlns[test.lvl](test.err, test.text)
		}

		stdcheck(t, mock, test.expected)
	}

	Wait(true)
}

func Test_alog_stringer(t *testing.T) {
	mock := stdmock(t)

	lg := newjsonlog(
		"",
		"INFO",
		"INFO",
		"INFO",
	)

	Println(lg)

	stdcheck(t, mock, fmt.Sprintf("[INFO] %s", lg.String()))

	Close()
}

func Test_alog_interface(t *testing.T) {
	mock := stdmock(t)

	lg := abnorm{"INFO"}

	Println(lg)

	stdcheck(t, mock, "[INFO] {INFO}")

	Close()
}

func Test_alog_nested(t *testing.T) {
	mock := stdmock(t)

	lg := iny{iny{iny{}}}

	Println(lg)

	stdcheck(t, mock, "[INFO] {{{<nil>}}}")

	Close()
}

func Test_alog_ctx_cancel(t *testing.T) {
	mock := &passmock{make(chan []byte)}
	ctx, cancel := context.WithCancel(context.Background())

	dest := Destination{
		lvls,
		STD,
		mock,
	}

	_ = Global(
		ctx,
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	)

	defer func() {
		if r := recover(); r != nil {
			t.Error("panic recovered in test")
		}
	}()

	func() {
		defer cancel()
		Println("TEST")
	}()

	Close()
}

func Test_alog_write_close(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		lvls,
		STD,
		mock,
	}

	_ = Global(
		context.Background(),
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wchan := make(chan interface{})

	Printc(ctx, wchan)

	defer func() {
		if r := recover(); r != nil {
			t.Error("panic recovered in test")
		}
	}()

	close(wchan)

	Close()
}

func Test_alog_sliced_struct(t *testing.T) {
	mock := stdmock(t)

	lg := []iny{
		{abnorm{"INFO"}},
		{abnorm{"INFO"}},
		{abnorm{"INFO"}},
		{abnorm{"INFO"}},
		{abnorm{"INFO"}},
	}

	Println(lg)

	stdcheck(t, mock, "[INFO] [{{INFO}} {{INFO}} {{INFO}} {{INFO}} {{INFO}}]")

	Close()
}

func Test_alog_sliced_interface(t *testing.T) {
	mock := stdmock(t)

	lg := []interface{}{[]interface{}{[]interface{}{"HELLOWORLD"}}}

	Println(lg)

	stdcheck(t, mock, "[INFO] HELLOWORLD")

	Close()
}

func Test_alog_sliced_string(t *testing.T) {
	mock := stdmock(t)

	lg := []string{"HELLO", "WORLD"}

	Println(lg)

	stdcheck(t, mock, "[INFO] HELLO,WORLD")

	Close()
}

func Test_alog_ln_multi(t *testing.T) {
	mock := stdmock(t)

	for _, test := range logs {
		if test.lvl&INFO > 0 {
			Println(test.text, test.text)
		} else if test.lvl&CUSTOM > 0 {
			Customln("CUSTOM", test.err, test.text, test.text)
		} else {
			stdlns[test.lvl](test.err, test.text, test.text)
		}

		// Loop twice since there should be two lines for each of these
		for i := 0; i < 2; i++ {
			stdcheck(t, mock, test.expected)
		}
	}

	Wait(true)
}

func Test_alog_normal(t *testing.T) {
	mock := stdmock(t)

	for _, test := range logs {
		if test.lvl&INFO > 0 {
			Print(test.text)
		} else if test.lvl&CUSTOM > 0 {
			Custom("CUSTOM", test.err, test.text)
		} else {
			std[test.lvl](test.err, test.text)
		}

		stdcheck(t, mock, test.expected)
	}

	Wait(true)
}

func Test_alog_multi(t *testing.T) {
	mock := stdmock(t)

	for _, test := range multi {
		if test.lvl&INFO > 0 {
			Print(test.text, test.text)
		} else if test.lvl&CUSTOM > 0 {
			Custom("CUSTOM", test.err, test.text, test.text)
		} else {
			std[test.lvl](test.err, test.text, test.text)
		}

		stdcheck(t, mock, test.expected)
	}

	Wait(true)
}

func Test_alog_normalf(t *testing.T) {
	mock := stdmock(t)

	for _, test := range flogs {
		if test.lvl&INFO > 0 {
			Printf("%s *%s*", test.text, test.text)
		} else if test.lvl&CUSTOM > 0 {
			Customf("CUSTOM", test.err, "%s *%s*", test.text, test.text)
		} else {
			stdfs[test.lvl](test.err, "%s *%s*", test.text, test.text)
		}

		stdcheck(t, mock, test.expected)
	}

	Wait(true)
}

func Test_alog_chan(t *testing.T) {
	mock := stdmock(t)
	for _, test := range clogs {

		func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			wchan := make(chan interface{})

			if test.lvl&CUSTOM > 0 {
				Customc(ctx, wchan, "CUSTOM")
			} else {
				chanfs[test.lvl](ctx, wchan)
			}

			select {
			case <-ctx.Done():
				return
			case wchan <- test.text:
			}

			stdcheck(t, mock, test.expected)
		}()
	}

	Wait(true)
}

func Test_alog_chan_err(t *testing.T) {
	mock := stdmock(t)

	for _, test := range cerrlogs {

		// Skip info for this test
		if test.lvl&INFO == 0 {
			func() {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				wchan := make(chan interface{})

				if test.lvl&CUSTOM > 0 {
					Customc(ctx, wchan, "CUSTOM")
				} else {
					chanfs[test.lvl](ctx, wchan)
				}

				select {
				case <-ctx.Done():
					return
				case wchan <- test.err:
				}

				stdcheck(t, mock, test.expected)
			}()
		}
	}

	Wait(true)
}
