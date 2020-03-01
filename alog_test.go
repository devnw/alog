package alog

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func check(value []byte, expected string) (err error) {

	if len(value) > 0 {

		output := string(value)

		if strings.LastIndex(output, "\n") == len(output)-1 {
			i := strings.Index(output, "[")
			output = strings.TrimSpace(output[i:])

			if expected != output {
				err = errors.Errorf("expected result: '%s' != output: '%s'", expected, output)
			}

		} else {
			err = errors.Errorf("expected newline at end of log")
		}
	} else {
		err = errors.Errorf("value is empty")
	}

	return err
}

func Test_alog_global_ln(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		INFO | DEBUG | TRACE | WARN | ERROR | CRIT | FATAL | CUSTOM,
		STD,
		mock,
	}

	if err := Global(
		context.Background(),
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	); err == nil {
		for _, test := range logs {
			if test.lvl&INFO > 0 {
				Println(test.text)
			} else if test.lvl&DEBUG > 0 {

				Debugln(test.err, test.text)
			} else if test.lvl&TRACE > 0 {

				Traceln(test.err, test.text)
			} else if test.lvl&WARN > 0 {

				Warnln(test.err, test.text)
			} else if test.lvl&ERROR > 0 {

				Errorln(test.err, test.text)
			} else if test.lvl&CRIT > 0 {

				Critln(test.err, test.text)
			} else if test.lvl&FATAL > 0 {

				Fatalln(test.err, test.text)
			} else if test.lvl&CUSTOM > 0 {

				Customln("CUSTOM", test.err, test.text)
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

	} else {
		fmt.Println(err)
	}
}

func Test_alog_global_ln_multi(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		INFO | DEBUG | TRACE | WARN | ERROR | CRIT | FATAL | CUSTOM,
		STD,
		mock,
	}

	if err := Global(
		context.Background(),
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	); err == nil {
		for _, test := range logs {
			if test.lvl&INFO > 0 {
				Println(test.text, test.text)
			} else if test.lvl&DEBUG > 0 {

				Debugln(test.err, test.text, test.text)
			} else if test.lvl&TRACE > 0 {

				Traceln(test.err, test.text, test.text)
			} else if test.lvl&WARN > 0 {

				Warnln(test.err, test.text, test.text)
			} else if test.lvl&ERROR > 0 {

				Errorln(test.err, test.text, test.text)
			} else if test.lvl&CRIT > 0 {

				Critln(test.err, test.text, test.text)
			} else if test.lvl&FATAL > 0 {

				Fatalln(test.err, test.text, test.text)
			} else if test.lvl&CUSTOM > 0 {

				Customln("CUSTOM", test.err, test.text, test.text)
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

	} else {
		fmt.Println(err)
	}
}

func Test_alog_global_normal(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		INFO | DEBUG | TRACE | WARN | ERROR | CRIT | FATAL | CUSTOM,
		STD,
		mock,
	}

	if err := Global(
		context.Background(),
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	); err == nil {
		for _, test := range logs {
			if test.lvl&INFO > 0 {
				Print(test.text)
			} else if test.lvl&DEBUG > 0 {

				Debug(test.err, test.text)
			} else if test.lvl&TRACE > 0 {

				Trace(test.err, test.text)
			} else if test.lvl&WARN > 0 {

				Warn(test.err, test.text)
			} else if test.lvl&ERROR > 0 {

				Error(test.err, test.text)
			} else if test.lvl&CRIT > 0 {

				Crit(test.err, test.text)
			} else if test.lvl&FATAL > 0 {

				Fatal(test.err, test.text)
			} else if test.lvl&CUSTOM > 0 {

				Custom("CUSTOM", test.err, test.text)
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

	} else {
		fmt.Println(err)
	}
}

func Test_alog_global_multi(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		INFO | DEBUG | TRACE | WARN | ERROR | CRIT | FATAL | CUSTOM,
		STD,
		mock,
	}

	if err := Global(
		context.Background(),
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	); err == nil {
		for _, test := range multi {
			if test.lvl&INFO > 0 {
				Print(test.text, test.text)
			} else if test.lvl&DEBUG > 0 {

				Debug(test.err, test.text, test.text)
			} else if test.lvl&TRACE > 0 {

				Trace(test.err, test.text, test.text)
			} else if test.lvl&WARN > 0 {

				Warn(test.err, test.text, test.text)
			} else if test.lvl&ERROR > 0 {

				Error(test.err, test.text, test.text)
			} else if test.lvl&CRIT > 0 {

				Crit(test.err, test.text, test.text)
			} else if test.lvl&FATAL > 0 {

				Fatal(test.err, test.text, test.text)
			} else if test.lvl&CUSTOM > 0 {

				Custom("CUSTOM", test.err, test.text, test.text)
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

	} else {
		fmt.Println(err)
	}
}

func Test_alog_global_normalf(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		INFO | DEBUG | TRACE | WARN | ERROR | CRIT | FATAL | CUSTOM,
		STD,
		mock,
	}

	if err := Global(
		context.Background(),
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	); err == nil {
		for _, test := range flogs {
			if test.lvl&INFO > 0 {
				Printf("%s *%s*", test.text, test.text)
			} else if test.lvl&DEBUG > 0 {

				Debugf(test.err, "%s *%s*", test.text, test.text)
			} else if test.lvl&TRACE > 0 {

				Tracef(test.err, "%s *%s*", test.text, test.text)
			} else if test.lvl&WARN > 0 {

				Warnf(test.err, "%s *%s*", test.text, test.text)
			} else if test.lvl&ERROR > 0 {

				Errorf(test.err, "%s *%s*", test.text, test.text)
			} else if test.lvl&CRIT > 0 {

				Critf(test.err, "%s *%s*", test.text, test.text)
			} else if test.lvl&FATAL > 0 {

				Fatalf(test.err, "%s *%s*", test.text, test.text)
			} else if test.lvl&CUSTOM > 0 {

				Customf("CUSTOM", test.err, "%s *%s*", test.text, test.text)
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

	} else {
		fmt.Println(err)
	}
}

func Test_alog_global_chan(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		INFO | DEBUG | TRACE | WARN | ERROR | CRIT | FATAL | CUSTOM,
		STD,
		mock,
	}

	if err := Global(
		context.Background(),
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	); err == nil {
		for _, test := range clogs {

			func() {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				wchan := make(chan interface{})

				if test.lvl&INFO > 0 {
					Printc(ctx, wchan)

				} else if test.lvl&DEBUG > 0 {

					Debugc(ctx, wchan)
				} else if test.lvl&TRACE > 0 {

					Tracec(ctx, wchan)
				} else if test.lvl&WARN > 0 {

					Warnc(ctx, wchan)
				} else if test.lvl&ERROR > 0 {

					Errorc(ctx, wchan)
				} else if test.lvl&CRIT > 0 {

					Critc(ctx, wchan)
				} else if test.lvl&FATAL > 0 {

					Fatalc(ctx, wchan)
				} else if test.lvl&CUSTOM > 0 {

					Customc(ctx, wchan, "CUSTOM")
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

	} else {
		fmt.Println(err)
	}
}

func Test_alog_global_chan_err(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		INFO | DEBUG | TRACE | WARN | ERROR | CRIT | FATAL | CUSTOM,
		STD,
		mock,
	}

	if err := Global(
		context.Background(),
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	); err == nil {
		for _, test := range cerrlogs {

			// Skip info for this test
			if test.lvl&INFO == 0 {
				func() {
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()

					wchan := make(chan interface{})

					if test.lvl&DEBUG > 0 {

						Debugc(ctx, wchan)
					} else if test.lvl&TRACE > 0 {

						Tracec(ctx, wchan)
					} else if test.lvl&WARN > 0 {

						Warnc(ctx, wchan)
					} else if test.lvl&ERROR > 0 {

						Errorc(ctx, wchan)
					} else if test.lvl&CRIT > 0 {

						Critc(ctx, wchan)
					} else if test.lvl&FATAL > 0 {

						Fatalc(ctx, wchan)
					} else if test.lvl&CUSTOM > 0 {

						Customc(ctx, wchan, "CUSTOM")
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

	} else {
		fmt.Println(err)
	}
}
