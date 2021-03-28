//+build !js

package audio

import (
	"testing"
	"time"

	"github.com/200sc/klangsynthese/audio/filter"
	"github.com/200sc/klangsynthese/synth"
)

func TestAudioFuncs(t *testing.T) {
	kla, err := synth.Int16.Sin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	a := New(DefFont, kla.(Data))
	err = <-a.Play()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	time.Sleep(a.PlayLength())
	// Assert audio is playing
	<-a.Play()
	err = a.Stop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	time.Sleep(a.PlayLength())
	// Assert audio is not playing
	kla, err = a.Copy()
	a = kla.(*Audio)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	a.Play()
	time.Sleep(a.PlayLength())
	// Assert audio is playing
	a = a.MustCopy().(*Audio)
	if a.Xp() != nil {
		t.Fatalf("audio without position had x pointer")
	}
	if a.Yp() != nil {
		t.Fatalf("audio without position had y pointer")
	}
	kla, err = a.Filter(filter.Volume(.5))
	a = kla.(*Audio)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	a.Play()
	time.Sleep(a.PlayLength())
	// Assert quieter audio is playing
	a = a.MustFilter(filter.Volume(.5)).(*Audio)
	a.Play()
	time.Sleep(a.PlayLength())
	// Assert yet quieter audio is playing
}
