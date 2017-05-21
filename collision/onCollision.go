package collision

import (
	"errors"

	"bitbucket.org/oakmoundstudio/oak/event"
)

// A Phase is a struct that other structs who want to use PhaseCollision
// should be composed of
type Phase struct {
	OnCollisionS *Space
	// If allocating maps becomes an issue
	// we can have two constant maps that we
	// switch between on alternating frames
	Touching map[int]bool
}

func (cp *Phase) getCollisionPhase() *Phase {
	return cp
}

type collisionPhase interface {
	getCollisionPhase() *Phase
}

// PhaseCollision binds to the entity behind the space's CID so that it will
// recieve CollisionStart and CollisionStop events, appropriately when
// entities begin to collide or stop colliding with the space.
func PhaseCollision(s *Space) error {
	switch t := event.GetEntity(int(s.CID)).(type) {
	case collisionPhase:
		oc := t.getCollisionPhase()
		oc.OnCollisionS = s
		s.CID.Bind(phaseCollisionEnter, "EnterFrame")
		return nil
	}
	return errors.New("This space's entity does not implement OnCollisionyThing")
}

func phaseCollisionEnter(id int, nothing interface{}) int {
	e := event.GetEntity(id).(collisionPhase)
	oc := e.getCollisionPhase()

	// check hits
	hits := Hits(oc.OnCollisionS)
	newTouching := map[int]bool{}

	// if any are new, trigger on collision
	for _, h := range hits {
		l := h.Label
		if _, ok := oc.Touching[l]; !ok {
			event.CID(id).Trigger("CollisionStart", l)
		}
		newTouching[l] = true
	}

	// if we lost any, trigger off collision
	for l := range oc.Touching {
		if _, ok := newTouching[l]; !ok {
			event.CID(id).Trigger("CollisionStop", l)
		}
	}

	oc.Touching = newTouching

	return 0
}
