package lights_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"test.com/lights"
)

func TestConfiguration(t *testing.T) {
	c := lights.Configuration{
		Yellow: 5,
		Red:    10,
		Green:  15,
	}

	assert.Equal(t, c.Yellow, 5)
	assert.Equal(t, c.Red, 10)
	assert.Equal(t, c.Green, 15)
}

func TestOnOff(t *testing.T) {
	c := lights.Configuration{
		Yellow: 5,
		Red:    10,
		Green:  15,
	}

	tl := lights.NewTrafficLight(context.Background(), c)

	assert.Equal(t, tl.Seconds(), 0)

	tl.On()
	time.Sleep(5 * time.Second)
	tl.Off()

	assert.Equal(t, tl.Seconds(), 5)
}

func TestLights(t *testing.T) {
	c := lights.Configuration{
		Yellow: 2,
		Green:  4,
		Red:    3,
	}

	tl := lights.NewTrafficLight(context.Background(), c)

	tl.On()
	assert.Equal(t, lights.Yellow, tl.Light())
	time.Sleep(3 * time.Second)
	assert.Equal(t, lights.Green, tl.Light())
	time.Sleep(5 * time.Second)
	assert.Equal(t, lights.Red, tl.Light())
	time.Sleep(2 * time.Second)
	assert.Equal(t, lights.Yellow, tl.Light())
	time.Sleep(3 * time.Second)
	assert.Equal(t, lights.Green, tl.Light())
	time.Sleep(5 * time.Second)
	assert.Equal(t, lights.Red, tl.Light())
	tl.Off()
}
