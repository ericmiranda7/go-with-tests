package _select

import (
	"errors"
	"net/http"
	"time"
)

var tenSecondDefaultTimeout = 10 * time.Second

func Racer(u1, u2 string) (string, error) {
	return ConfigurableRacer(u1, u2, tenSecondDefaultTimeout)
}

func ConfigurableRacer(u1, u2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(u1):
		return u1, nil
	case <-ping(u2):
		return u2, nil
	case <-time.After(20 * time.Millisecond):
		return "", errors.New("timeout")
	}
}

func ping(u string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(u)
		close(ch)
	}()

	return ch
}
