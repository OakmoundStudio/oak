package audio

import (
	"fmt"

	"github.com/oakmound/oak/v3/audio/font"
	"github.com/oakmound/oak/v3/audio/klang"
	"github.com/oakmound/oak/v3/oakerr"
)

// Audio is a struct of some audio data and the variables
// required to filter it through a sound font.
type Audio struct {
	*font.Audio
	toStop    klang.Audio
	X, Y      *float64
	setVolume int32
}

// New returns an audio from a font, some audio data, and optional
// positional coordinates
func New(f *font.Font, d Data, coords ...*float64) *Audio {
	a := new(Audio)
	a.Audio = font.NewAudio(f, d)
	if len(coords) > 0 {
		a.X = coords[0]
		if len(coords) > 1 {
			a.Y = coords[1]
		}
	}
	return a
}

// SetVolume attempts to set the volume of the underlying OS audio.
func (a *Audio) SetVolume(v int32) error {
	a.setVolume = v
	if a.toStop != nil {
		return a.toStop.SetVolume(v)
	}
	return nil
}

// Play begin's an audio's playback
func (a *Audio) Play() <-chan error {
	a2, err := a.Copy()
	if err != nil {
		return errChannel(err)
	}
	a3, err := a2.Filter(a.Font.Filters...)
	if err != nil {
		return errChannel(err)
	}
	a4, err := a3.(*Audio).FullAudio.Copy()
	if err != nil {
		return errChannel(err)
	}
	a.toStop = a4
	a.toStop.SetVolume(a.setVolume)
	return a4.Play()
}

func errChannel(err error) <-chan error {
	ch := make(chan error)
	go func() {
		ch <- err
	}()
	return ch
}

// Stop stops an audio's playback
func (a *Audio) Stop() error {
	if a == nil || a.toStop == nil {
		return oakerr.NilInput{InputName: "Audio"}
	}
	return a.toStop.Stop()
}

// Copy returns a copy of the audio
func (a *Audio) Copy() (klang.Audio, error) {
	a2, err := a.Audio.Copy()
	if err != nil {
		return nil, err
	}
	return New(a.Audio.Font, a2.(klang.FullAudio), a.X, a.Y), nil
}

// MustCopy acts like Copy, but panics on an error.
func (a *Audio) MustCopy() klang.Audio {
	return New(a.Audio.Font, a.Audio.MustCopy().(klang.FullAudio), a.X, a.Y)
}

// Filter returns the audio with some set of filters applied to it.
func (a *Audio) Filter(fs ...klang.Filter) (klang.Audio, error) {
	var ad klang.Audio = a
	var err, consErr error
	for _, f := range fs {
		ad, err = f.Apply(ad)
		if err != nil {
			if consErr == nil {
				consErr = err
			} else {
				consErr = fmt.Errorf("%w, %v", err, consErr)
			}
		}
	}
	return ad, consErr
}

// MustFilter acts like Filter but ignores errors.
func (a *Audio) MustFilter(fs ...klang.Filter) klang.Audio {
	ad, _ := a.Filter(fs...)
	return ad
}

// Xp returns a pointer to the x position of this audio, if it has one.
// It has no position, this returns nil.
func (a *Audio) Xp() *float64 {
	return a.X
}

// Yp returns a pointer to the y position of this audio, if it has one.
// It has no position, this returns nil. If This is not nil, Xp will not be nil.
func (a *Audio) Yp() *float64 {
	return a.Y
}

var (
	// Guarantee that Audio can have positional filters applied to it
	_ SupportsPos = &Audio{}
)
