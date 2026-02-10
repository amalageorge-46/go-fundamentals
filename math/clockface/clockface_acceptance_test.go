package clockface_test

import (
	"testing"
	"time"

	"example.com/hello/math/clockface"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	// You MUST use the package prefix because you are in clockface_test
	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

//func TestSecondHandAt30Seconds(t *testing.T) {
//	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

//	want := clockface.Point{X: 150, Y: 150 + 90}
//	got := clockface.SecondHand(tm)

//	if got != want {
//		t.Errorf("Got %v, wanted %v", got, want)
//	}
//}
