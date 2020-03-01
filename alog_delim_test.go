package alog

import (
	"context"
	"testing"
)

func Test_alog_ln_DELIM(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := &Destination{
		lvls,
		DELIM,
		mock,
	}

	if err := testg(dest, mock); err != nil {
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
			t.Log(string(log))
			if err := check(log, test.expected); err != nil {
				t.Error(err)
			}
		} else {
			return
		}
	}

	Wait(true)
}

func Test_alog_ln_multi_DELIM(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := &Destination{
		lvls,
		DELIM,
		mock,
	}

	if err := testg(dest, mock); err != nil {
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

func Test_alog_normal_DELIM(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := &Destination{
		lvls,
		DELIM,
		mock,
	}

	if err := testg(dest, mock); err != nil {
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

func Test_alog_multi_DELIM(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := &Destination{
		lvls,
		DELIM,
		mock,
	}

	if err := testg(dest, mock); err != nil {
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

func Test_alog_normalf_DELIM(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := &Destination{
		lvls,
		DELIM,
		mock,
	}

	if err := testg(dest, mock); err != nil {
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

func Test_alog_chan_DELIM(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := &Destination{
		lvls,
		DELIM,
		mock,
	}

	if err := testg(dest, mock); err != nil {
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

func Test_alog_chan_err_DELIM(t *testing.T) {
	mock := &passmock{make(chan []byte)}

	dest := &Destination{
		lvls,
		DELIM,
		mock,
	}

	if err := testg(dest, mock); err != nil {
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
