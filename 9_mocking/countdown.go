package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const coundownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep (){
    c.sleep(c.duration)

}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleep Sleeper) {
	for i := coundownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleep.Sleep()
	}

	fmt.Fprint(out, finalWord)

}

func main() {

    sleeper := &ConfigurableSleeper{2 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
