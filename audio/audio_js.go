//+build js

package audio

import "errors"

var (
	loadedWavs = make(map[string]Audio, 0)
)

// We alias the winaudio package's interface here
// so game files don't need to import winaudio
type Audio interface {
	Pause() error
	Stop() error
	Resume() error
	Play() error
	SetLooping(loop bool)
	SetVolume(volume int32) error
	GetVolume() (int32, error)
	SetPan(pan int32) error
	GetPan() (int32, error)
	SetFrequency(freq uint32) error
	GetFrequency() (uint32, error)
}

type PosAudio struct {
	Audio
	X, Y *int
}

type PosAudioF struct {
	Audio
	X, Y *float64
}

func InitWinAudio() {
}

func GetSounds(fileNames ...string) ([]Audio, error) {
	sounds := make([]Audio, len(fileNames))
	return sounds, errors.New("Audio not supported on JS")
}

func GetWav(fileName string) (Audio, error) {
	return nil, errors.New("Audio not supported on JS")
}

func PlayWav(fileName string) error {
	return errors.New("Audio not supported on JS")
}

func LoadWav(directory, fileName string) (Audio, error) {
	return nil, errors.New("Audio not supported on JS")
}

func BatchLoad(baseFolder string) error {
	return errors.New("Audio not supported on JS")
}
func GetActivePosWavChannel(frequency, freqRand int, fileNames ...string) (chan [3]int, error) {
	soundCh := make(chan [3]int)
	return soundCh, errors.New("Audio not supported on JS")
}

func GetActiveWavChannel(frequency, freqRand int, fileNames ...string) (chan int, error) {

	soundCh := make(chan int)
	return soundCh, errors.New("Audio not supported on JS")
}

// Non-Active channels will attempt to steal most sends sent to the output
// audio channel. This will allow a game to constantly send on a channel and
// obtain an output rate of near the sent in frequency instead of locking
// or requiring buffered channel usage.
//
// An important example case-- walking around
// When a character walks, they have some frequency step speed and some
// set of potential fileName sounds that play, and the usage of a channel
// here will let the EnterFrame code which detects the walking status to
// send on the walking audio channel constantly without worrying about
// triggering too many sounds.

func GetPosWavChannel(frequency, freqRand int, fileNames ...string) (chan [3]int, error) {
	return GetActivePosWavChannel(frequency, freqRand, fileNames...)
}

func GetWavChannel(frequency, freqRand int, fileNames ...string) (chan int, error) {

	return GetActiveWavChannel(frequency, freqRand, fileNames...)
}

// For Pan and volume calculation
func SetEars(x, y *float64, panWidth float64, silentRadius float64) {
}

func CalculatePan(x2 float64) int32 {
	return 0
}

func CalculateVolume(x2, y2 float64) int32 {
	return 0
}

func PlayPositional(sound Audio, x, y float64) (err error) {
	return errors.New("Audio not supported on JS")
}
