package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("test output", func(t *testing.T) {
		buff := &bytes.Buffer{}
		// a reference to the struct
		spySleeper := &SpyCountdownOperations{}

		Countdown(buff, spySleeper)

		got := buff.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q ", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		// this is a pointer to a object thats created
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}

	})

	t.Run("Customizing sleep time ", func(t *testing.T) {
		sleepTime := 5 * time.Second

		// the reason we are passing around references is so we can
		// mutate the contents inside of that struct, not just a copy
		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()


		if spyTime.durationSlept != sleepTime{
			t.Errorf("slept %q, wanted %q", spyTime.durationSlept, sleepTime)
		}
	})
}
