package collision

import (
	"testing"
	"time"

	"github.com/oakmound/oak/v3/event"
)

type cphase struct {
	Phase
	callers *event.CallerMap
}

func (cp *cphase) Init() event.CID {
	return cp.callers.NextID(cp)
}

func TestCollisionPhase(t *testing.T) {
	callers := event.NewCallerMap()
	bus := event.NewBus(callers)
	go bus.ResolveChanges()
	go func() {
		for {
			<-time.After(5 * time.Millisecond)
			<-bus.TriggerBack(event.Enter, nil)
		}
	}()
	cp := cphase{
		callers: callers,
	}
	cid := cp.Init()
	s := NewSpace(10, 10, 10, 10, cid)
	tree := NewTree()
	err := PhaseCollisionWithBus(s, tree, bus, callers)
	if err != nil {
		t.Fatalf("phase collision failed: %v", err)
	}
	var active bool
	bus.Bind("CollisionStart", cid, func(event.CID, interface{}) int {
		active = true
		return 0
	})
	bus.Bind("CollisionStop", cid, func(event.CID, interface{}) int {
		active = false
		return 0
	})

	s2 := NewLabeledSpace(15, 15, 10, 10, 5)
	tree.Add(s2)
	time.Sleep(200 * time.Millisecond)
	if !active {
		t.Fatalf("collision should be active")
	}

	tree.Remove(s2)
	time.Sleep(200 * time.Millisecond)
	if active {
		t.Fatalf("collision should be inactive")
	}

	s3 := NewSpace(10, 10, 10, 10, 5)
	err = PhaseCollision(s3, nil)
	if err == nil {
		t.Fatalf("phase collision should have failed")
	}
}
