package alog

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_alog_global_defaults(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		lvls,
		STD,
		mock,
	}

	if err := Global(
		context.Background(),
		"PREFIX",
		"",
		nil,
		-1,
		dest,
	); err == nil {

		for _, test := range prefixlogs {
			if test.lvl&INFO > 0 {
				Println(test.text)
			} else if test.lvl&CUSTOM > 0 {
				Customln("CUSTOM", test.err, test.text)
			} else {
				stdlns[test.lvl](test.err, test.text)
			}

			if log, ok := <-mock.msg; ok {
				if err := check(log, test.expected); err != nil {
					t.Error(err)
				}
			} else {
				return
			}
		}

		Close()
	}
}

func Test_alog_Global(t *testing.T) {
	if err := Global(
		context.Background(),
		"PREFIX",
		"",
		nil,
		-1,
	); err == nil {

		t.Error("expected error but got success")
	}
}

func Test_alog_setGlobal(t *testing.T) {
	if err := setGlobal(nil); err == nil {

		t.Error("expected error but got success")
	}
}

func Test_alog_ln(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	for _, test := range logs {
		if test.lvl&INFO > 0 {
			Println(test.text)
		} else if test.lvl&CUSTOM > 0 {
			Customln("CUSTOM", test.err, test.text)
		} else {
			stdlns[test.lvl](test.err, test.text)
		}

		if log, ok := <-mock.msg; ok {
			if err := check(log, test.expected); err != nil {
				t.Error(err)
			}
		} else {
			return
		}
	}

	Wait(true)
}

func Test_alog_stringer(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	lg := newjsonlog(
		"",
		"INFO",
		"INFO",
		"INFO",
	)

	Println(lg)

	if log, ok := <-mock.msg; ok {
		if err := check(log, fmt.Sprintf("[INFO] %s", lg.String())); err != nil {
			t.Error(err)
		}
	} else {
		return
	}

	Close()
}

func Test_alog_interface(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	lg := abnorm{"INFO"}

	Println(lg)

	if log, ok := <-mock.msg; ok {
		if err := check(log, "[INFO] {INFO}"); err != nil {
			t.Error(err)
		}
	} else {
		return
	}

	Close()
}

func Test_alog_nested(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	lg := iny{iny{iny{}}}

	Println(lg)

	if log, ok := <-mock.msg; ok {
		if err := check(log, "[INFO] {{{<nil>}}}"); err != nil {
			t.Error(err)
		}
	} else {
		return
	}

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
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	lg := []iny{
		{abnorm{"INFO"}},
		{abnorm{"INFO"}},
		{abnorm{"INFO"}},
		{abnorm{"INFO"}},
		{abnorm{"INFO"}},
	}

	Println(lg)

	if log, ok := <-mock.msg; ok {
		if err := check(log, "[INFO] [{{INFO}} {{INFO}} {{INFO}} {{INFO}} {{INFO}}]"); err != nil {
			t.Error(err)
		}
	} else {
		return
	}

	Close()
}

func Test_alog_sliced_interface(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	lg := []interface{}{[]interface{}{[]interface{}{"HELLOWORLD"}}}

	Println(lg)

	if log, ok := <-mock.msg; ok {
		if err := check(log, "[INFO] HELLOWORLD"); err != nil {
			t.Error(err)
		}
	} else {
		return
	}

	Close()
}

func Test_alog_sliced_string(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	lg := []string{"HELLO", "WORLD"}

	Println(lg)

	if log, ok := <-mock.msg; ok {
		if err := check(log, "[INFO] HELLO,WORLD"); err != nil {
			t.Error(err)
		}
	} else {
		return
	}

	Close()
}

func Test_alog_empty(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	Print()
	checkunable(t, mock, INFO)

	Debug(nil)
	checkunable(t, mock, DEBUG)

	Trace(nil)
	checkunable(t, mock, TRACE)

	Warn(nil)
	checkunable(t, mock, WARN)

	Error(nil)
	checkunable(t, mock, ERROR)

	Crit(nil)
	checkunable(t, mock, CRIT)

	Fatal(nil)
	checkunable(t, mock, FATAL)

	Custom("CUSTOM", nil)
	checkunable(t, mock, CUSTOM)

	Close()
}

func checkunable(t *testing.T, mock *passmock, level LogLevel) {

	if log, ok := <-mock.msg; ok {
		if err := check(log, fmt.Sprintf("[%s]unable to create log string, empty message and error", types[level])); err != nil {
			t.Error(err)
		}
	} else {
		return
	}
}

func Test_alog_empty_ln(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	Println()
	checkwarn(t, mock)

	Debugln(nil)
	checkwarn(t, mock)

	Traceln(nil)
	checkwarn(t, mock)

	Warnln(nil)
	checkwarn(t, mock)

	Errorln(nil)
	checkwarn(t, mock)

	Critln(nil)
	checkwarn(t, mock)

	Fatalln(nil)
	checkwarn(t, mock)

	Customln("CUSTOM", nil)
	checkwarn(t, mock)

	Close()
}

func checkwarn(t *testing.T, mock *passmock) {

	if log, ok := <-mock.msg; ok {
		if err := check(log, "[WARN] empty log value passed"); err != nil {
			t.Error(err)
		}
	} else {
		return
	}
}

func Test_alog_ln_multi(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

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

			if log, ok := <-mock.msg; ok {
				if err := check(log, test.expected); err != nil {
					t.Error(err)
				}
			} else {
				return
			}
		}
	}

	Wait(true)
}

func Test_alog_normal(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	for _, test := range logs {
		if test.lvl&INFO > 0 {
			Print(test.text)
		} else if test.lvl&CUSTOM > 0 {
			Custom("CUSTOM", test.err, test.text)
		} else {
			std[test.lvl](test.err, test.text)
		}

		if log, ok := <-mock.msg; ok {
			if err := check(log, test.expected); err != nil {
				t.Error(err)
			}
		} else {
			return
		}
	}

	Wait(true)
}

func Test_alog_multi(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	for _, test := range multi {
		if test.lvl&INFO > 0 {
			Print(test.text, test.text)
		} else if test.lvl&CUSTOM > 0 {
			Custom("CUSTOM", test.err, test.text, test.text)
		} else {
			std[test.lvl](test.err, test.text, test.text)
		}

		if log, ok := <-mock.msg; ok {
			if err := check(log, test.expected); err != nil {
				t.Error(err)
			}
		} else {
			return
		}
	}

	Wait(true)
}

func Test_alog_normalf(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

	for _, test := range flogs {
		if test.lvl&INFO > 0 {
			Printf("%s *%s*", test.text, test.text)
		} else if test.lvl&CUSTOM > 0 {
			Customf("CUSTOM", test.err, "%s *%s*", test.text, test.text)
		} else {
			stdfs[test.lvl](test.err, "%s *%s*", test.text, test.text)
		}

		if log, ok := <-mock.msg; ok {
			if err := check(log, test.expected); err != nil {
				t.Error(err)
			}
		} else {
			return
		}
	}

	Wait(true)
}

func Test_alog_chan(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

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

			if log, ok := <-mock.msg; ok {
				if err := check(log, test.expected); err != nil {
					t.Error(err)
				}
			} else {
				return
			}
		}()
	}

	Wait(true)
}

func Test_alog_chan_err(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	if err := testg(nil, mock); err != nil {
		t.Error(err)
		return
	}

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

				if log, ok := <-mock.msg; ok {
					if err := check(log, test.expected); err != nil {
						t.Error(err)
					}
				} else {
					return
				}
			}()
		}
	}

	Wait(true)
}
