package lights

import (
	"fmt"
	"sync"
	"time"
)

const (
	Yellow = "Y"
	Green  = "G"
	Red    = "R"
)

type TrafficLight interface {
	On()
	Seconds() int
	Light() string
	Off()
}

type Configuration struct {
	Yellow int
	Red    int
	Green  int
}

type trafficLight struct {
	on      bool
	seconds int
	period  int
	config  Configuration
	light   string
	mu      sync.Mutex
}

func NewTrafficLight(c Configuration) TrafficLight {
	return &trafficLight{
		on:      false,
		config:  c,
		seconds: 0,
		period:  0,
		light:   Yellow,
	}
}

func (t *trafficLight) On() {
	if !t.on {
		t.on = true

		go func() {
			for {

				t.mu.Lock()
				t.seconds += 1
				t.mu.Unlock()

				time.Sleep(1 * time.Second)

				if !t.on {
					break
				}
			}
		}()

		go func() {
			for {
				t.mu.Lock()
				t.period += 1
				t.mu.Unlock()
				time.Sleep(1 * time.Second)
				fmt.Println(t.light, " ", t.period, " seconds")
				if t.light == Yellow && t.config.Yellow == t.period {
					t.mu.Lock()
					t.period = 0
					t.mu.Unlock()
					t.light = Green
				}

				if t.light == Green && t.config.Green == t.period {
					t.mu.Lock()
					t.period = 0
					t.mu.Unlock()
					t.light = Red
				}

				if t.light == Red && t.config.Red == t.period {
					t.mu.Lock()
					t.period = 0
					t.mu.Unlock()
					t.light = Yellow
				}

			}
		}()
	}
}

func (t *trafficLight) Off() {
	if t.on {
		t.on = false
	}
}

func (t *trafficLight) Seconds() int {
	return t.seconds
}

func (t *trafficLight) Light() string {
	return t.light
}
