//+build !js

package oak

import (
	"bytes"
	"testing"

	"github.com/oakmound/oak/v2/collision"
	"github.com/oakmound/oak/v2/event"
	"github.com/oakmound/oak/v2/mouse"
	"github.com/oakmound/oak/v2/render"
)

type ent struct{}

func (e ent) Init() event.CID {
	return 0
}

func TestDebugConsole(t *testing.T) {
	c1 := NewController()
	conf.LoadBuiltinCommands = true
	triggered := false
	err := c1.AddCommand("test", func([]string) {
		triggered = true
	})
	if err != nil {
		t.Fatalf("failed to add test command")
	}

	render.UpdateDebugMap("r", render.EmptyRenderable())

	event.NextID(ent{})

	rCh := make(chan bool)
	sCh := make(chan bool)
	r := bytes.NewBufferString(
		"test\n" +
			"nothing\n" +
			"viewport unlock\n" +
			"viewport unlock\n" +
			"viewport lock\n" +
			"viewport lock\n" +
			"viewport nothing\n" +
			"fade nothing\n" +
			"fade nothing 100\n" +
			"fade r\n" +
			"skip nothing\n" +
			"print nothing\n" +
			"print 2\n" +
			"print 1\n" +
			"mouse nothing\n" +
			"mouse details\n" +
			"garbage input\n" +
			"\n" +
			"skip scene\n")
	go c1.debugConsole(rCh, sCh, r)
	rCh <- true
	sleep()
	sleep()
	if !triggered {
		t.Fatalf("debug console did not trigger test command")
	}
	<-sCh
}

func TestMouseDetails(t *testing.T) {
	c1 := NewController()

	c1.mouseDetails(0, mouse.NewZeroEvent(0, 0))
	s := collision.NewUnassignedSpace(-1, -1, 2, 2)
	collision.Add(s)
	c1.mouseDetails(0, mouse.NewZeroEvent(0, 0))
	collision.Remove(s)

	// This should spew this nothing entity, but it doesn't.
	id := event.NextID(ent{})
	s = collision.NewSpace(-1, -1, 2, 2, id)
	c1.mouseDetails(0, mouse.NewZeroEvent(0, 0))
	collision.Remove(s)
}
