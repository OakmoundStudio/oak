// +build !windows,!linux,!darwin

package joystick

import "errors"

func newOsJoystick() osJoystick {
	return osJoystick{}
}

type osJoystick struct {
}

func osinit() {}

func (j *Joystick) prepare() error {
	return errors.New("OS not supported")
}

func (j *Joystick) getState() (*State, error) {
	return nil, errors.New("OS not supported")
}

func (j *Joystick) vibrate(left, right uint16) error {
	return errors.New("OS not supported")
}

func (j *Joystick) close() error {
	return errors.New("OS not supported")
}

func getJoysticks() []*Joystick {
	return nil
}
