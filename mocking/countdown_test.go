package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(d time.Duration) {
	s.durationSlept = d
}

type SpyCountdownOperations struct {
	Calls []string
}

func (sco *SpyCountdownOperations) Sleep() {
	sco.Calls = append(sco.Calls, sleep)
}

func (sco *SpyCountdownOperations) Write(_ []byte) (n int, err error) {
	sco.Calls = append(sco.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

func TestCountdown(t *testing.T) {
	t.Run("prints liftoff", func(t *testing.T) {
		b := bytes.Buffer{}

		Countdown(&b, &SpyCountdownOperations{})

		got := b.String()
		wantString := `3
2
1
Go!`

		if got != wantString {
			t.Errorf("got %q, want %q", got, wantString)
		}
	})

	t.Run("order matters", func(t *testing.T) {
		sc := SpyCountdownOperations{}

		Countdown(&sc, &sc)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(sc.Calls, want) {
			t.Errorf("wanted calls %v but got %v", want, sc.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	d := 5 * time.Second
	spyTime := SpyTime{}
	cs := ConfigurableSleeper{d, spyTime.Sleep}
	cs.Sleep()

	got := spyTime.durationSlept
	want := d

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
