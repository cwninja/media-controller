package tv

import "testing"
import "math/rand"

func TestTimeFoematting(t * testing.T) {
  for i:= 0; i < 1000; i++ {
    time := rand.Intn(10000000)
    formatted := formatTime(time)
    parsed := parseTime(formatted)
    if time != parsed {
      t.Error("Put in", time, "and got out", parsed)
    }
  }
}

func TestShortTimes(t * testing.T) {
  time := 10
  formatted := formatTime(time)
  if formatted != "00:00:10" {
    t.Error("Put in", time, "and got out", formatted)
  }
}
