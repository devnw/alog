package alog

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/pkg/errors"
)

type fakelog struct {
	tp   LogLevel
	text string
	err  error
}

func randtypes(ctx context.Context, count int) <-chan fakelog {
	rand.Seed(time.Now().UnixNano())
	rands := make(chan fakelog)

	go func(rands chan<- fakelog) {
		defer close(rands)

		i := 0
		for i < count {
			v := LogLevel(rand.Intn(int(CUSTOM)-int(INFO)+1) + int(INFO))

			if v <= CUSTOM {
				var e error
				if v&INFO == 0 {
					e = errors.New(randomdata.SillyName())
				}

				fl := fakelog{
					v,
					randomdata.SillyName(),
					e,
				}

				select {
				case <-ctx.Done():
					return
				case rands <- fl:
					i++
				}
			}
		}
	}(rands)

	return rands
}

func Test_alog_2global(t *testing.T) {

	mock := &writemock{}

	dest := Destination{
		INFO | DEBUG | TRACE | WARN | ERROR | CRIT | FATAL | CUSTOM,
		STD,
		mock,
	}
	ctx, cancel := context.WithCancel(context.Background())

	if err := Global(
		ctx,
		"",
		DEFAULTTIMEFORMAT,
		time.UTC,
		DEFAULTBUFFER,
		dest,
	); err == nil {
		ls := randtypes(ctx, 10)

		defer cancel()
		for {
			select {
			case <-ctx.Done():
				return
			case l, ok := <-ls:
				if ok {
					if l.tp&INFO > 0 {
						Println(l.text)
					} else if l.tp&DEBUG > 0 {

						Debugln(l.err, l.text)
					} else if l.tp&TRACE > 0 {

						Traceln(l.err, l.text)
					} else if l.tp&WARN > 0 {

						Warnln(l.err, l.text)
					} else if l.tp&ERROR > 0 {

						Errorln(l.err, l.text)
					} else if l.tp&CRIT > 0 {

						Critln(l.err, l.text)
					} else if l.tp&FATAL > 0 {

						Fatalln(l.err, l.text)
					} else if l.tp&CUSTOM > 0 {

						Customln(randomdata.SillyName(), l.err, l.text)
					}
				} else {
					return
				}
			}
		}

	} else {
		cancel()
		t.Error(err)
	}
}

func Test_alog_global(t *testing.T) {

	mock := &passmock{make(chan []byte)}

	dest := Destination{
		INFO | WARN | ERROR | CRIT | FATAL | CUSTOM,
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

		Critln(errors.New("TEST CRIT"), "HELLO WORLD")

		fmt.Println(string(<-mock.msg))

		Println("HELLO WORLD")

		fmt.Println(string(<-mock.msg))

		Debugln(errors.New("TEST DEBUG"), "HELLO WORLD")

		// NO DEBUG DEST fmt.Println(string(<-mock.msg))

		Critln(errors.New("TEST CRIT"), "HELLO WORLD")

		fmt.Println(string(<-mock.msg))

		Wait(true)

	} else {
		fmt.Println(err)
	}
}
