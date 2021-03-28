//+build !js

package mouse

import (
	"testing"
	"time"

	"github.com/oakmound/oak/v2/alg/floatgeom"
	"github.com/oakmound/oak/v2/collision"
	"github.com/oakmound/oak/v2/event"
)

type cphase struct {
	CollisionPhase
}

func (cp *cphase) Init() event.CID {
	return event.NextID(cp)
}

func TestCollisionPhase(t *testing.T) {
	go event.ResolvePending()
	go func() {
		for {
			<-time.After(5 * time.Millisecond)
			<-event.TriggerBack(event.Enter, nil)
		}
	}()
	cp := cphase{}
	cid := cp.Init()
	s := collision.NewSpace(10, 10, 10, 10, cid)
	if PhaseCollision(s) != nil {
		t.Fatalf("phase collision errored")
	}
	var active bool
	cid.Bind("MouseCollisionStart", func(event.CID, interface{}) int {
		active = true
		return 0
	})
	cid.Bind("MouseCollisionStop", func(event.CID, interface{}) int {
		active = false
		return 0
	})
	time.Sleep(200 * time.Millisecond)
	LastEvent = Event{floatgeom.Point2{10, 10}, ButtonNone, ""}
	time.Sleep(200 * time.Millisecond)
	if !active {
		t.Fatalf("phase collision did not trigger")
	}
	LastEvent = Event{floatgeom.Point2{21, 21}, ButtonNone, ""}
	time.Sleep(200 * time.Millisecond)
	if active {
		t.Fatalf("phase collision triggered innapropriately")
	}
	s = collision.NewSpace(10, 10, 10, 10, 5)
	if PhaseCollision(s) == nil {
		t.Fatalf("phase collision did not error on invalid space")
	}
}
