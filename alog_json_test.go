package alog

import (
	"context"
	"testing"
)

func jsonhelper(t *testing.T, tfunc func(mock *passmock)) {
	mock := &passmock{make(chan []byte)}

	dest := &Destination{
		lvls,
		JSON,
		mock,
	}

	err := testg(dest, nil)
	if err != nil {
		t.Fatal(err)
	}

	tfunc(mock)

	Wait(true)
}

func jsoncheck(t *testing.T, mock *passmock, expected jsonlogtest) {
	log, ok := <-mock.msg
	if !ok {
		t.Fatal("closed channel")
	}

	err := checkJSON(log, expected)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_alog_global_defaults_JSON(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := Destination{
		lvls,
		JSON,
		mock,
	}

	err := Global(
		context.Background(),
		"PREFIX",
		"",
		nil,
		-1,
		dest,
	)

	if err != nil {
		t.Fatal(err)
	}

	for _, test := range prefixlogsJSON {
		if test.lvl&INFO > 0 {
			Println(test.text)
		} else if test.lvl&CUSTOM > 0 {
			Customln("CUSTOM", test.err, test.text)
		} else {
			stdlns[test.lvl](test.err, test.text)
		}

		jsoncheck(t, mock, test.expected)
	}

	Close()
}

func Test_alog_ln_JSON(t *testing.T) {
	jsonhelper(t, func(mock *passmock) {
		for _, test := range logsJSON {
			if test.lvl&INFO > 0 {
				Println(test.text)
			} else if test.lvl&CUSTOM > 0 {
				Customln("CUSTOM", test.err, test.text)
			} else {
				stdlns[test.lvl](test.err, test.text)
			}

			jsoncheck(t, mock, test.expected)
		}
	})
}

func Test_alog_ln_multi_JSON(t *testing.T) {
	jsonhelper(t, func(mock *passmock) {
		for _, test := range logsJSON {
			if test.lvl&INFO > 0 {
				Println(test.text, test.text)
			} else if test.lvl&CUSTOM > 0 {
				Customln("CUSTOM", test.err, test.text, test.text)
			} else {
				stdlns[test.lvl](test.err, test.text, test.text)
			}

			// Loop twice since there should be two lines for each of these
			for i := 0; i < 2; i++ {
				jsoncheck(t, mock, test.expected)
			}
		}
	})
}

func Test_alog_normal_JSON(t *testing.T) {
	jsonhelper(t, func(mock *passmock) {
		for _, test := range logsJSON {
			if test.lvl&INFO > 0 {
				Print(test.text)
			} else if test.lvl&CUSTOM > 0 {
				Custom("CUSTOM", test.err, test.text)
			} else {
				std[test.lvl](test.err, test.text)
			}

			jsoncheck(t, mock, test.expected)
		}
	})
}

func Test_alog_multi_JSON(t *testing.T) {
	jsonhelper(t, func(mock *passmock) {
		for _, test := range multiJSON {
			if test.lvl&INFO > 0 {
				Print(test.text, test.text)
			} else if test.lvl&CUSTOM > 0 {
				Custom("CUSTOM", test.err, test.text, test.text)
			} else {
				std[test.lvl](test.err, test.text, test.text)
			}

			jsoncheck(t, mock, test.expected)
		}
	})
}

func Test_alog_normalf_JSON(t *testing.T) {
	jsonhelper(t, func(mock *passmock) {
		for _, test := range flogsJSON {
			if test.lvl&INFO > 0 {
				Printf("%s *%s*", test.text, test.text)
			} else if test.lvl&CUSTOM > 0 {
				Customf("CUSTOM", test.err, "%s *%s*", test.text, test.text)
			} else {
				stdfs[test.lvl](test.err, "%s *%s*", test.text, test.text)
			}

			jsoncheck(t, mock, test.expected)
		}
	})
}

func Test_alog_chan_JSON(t *testing.T) {
	jsonhelper(t, func(mock *passmock) {
		for _, test := range clogsJSON {
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

				jsoncheck(t, mock, test.expected)
			}()
		}
	})
}

func Test_alog_chan_err_JSON(t *testing.T) {
	jsonhelper(t, func(mock *passmock) {
		for _, test := range cerrlogsJSON {
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

					jsoncheck(t, mock, test.expected)
				}()
			}
		}
	})
}
