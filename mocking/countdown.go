package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const startCount = 3
const finalMessage = "Go!"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(d time.Duration)
}

func (cs ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

func main() {
	s := ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, s)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := startCount; i > 0; i-- {
		_, err := fmt.Fprintln(w, i)
		if err != nil {
			return
		}

		s.Sleep()
	}

	_, err := fmt.Fprintf(w, finalMessage)
	if err != nil {
		return
	}
}
