// Package collision uses an rtree to track rectangles
// and their intersections.
package collision

import (
	"sync"

	"bitbucket.org/oakmoundstudio/oak/physics"

	"github.com/Sythe2o0/rtreego"
)

var (
	rt      *rtreego.Rtree
	addLock = sync.Mutex{}
)

// A Point is a specific point where
// collision occured and a zone to identify
// what was collided with.
type Point struct {
	physics.Vector
	Zone *Space
}

// NewPoint creates a new point
func NewPoint(s *Space, x, y float64) Point {
	return Point{physics.NewVector(x, y), s}
}

// IsNil returns whether the underlying zone of a Point is nil
func (cp Point) IsNil() bool {
	return cp.Zone == nil
}

// Init sets the package global rtree to a new rtree.
func Init() {
	rt = rtreego.NewTree(20, 40)
}

// Clear just calls init.
func Clear() {
	Init()
}

// Add adds a set of spaces to the rtree
func Add(sps ...*Space) {
	addLock.Lock()
	for _, sp := range sps {
		if sp != nil {
			rt.Insert(sp)
		}
	}
	addLock.Unlock()
}

// Remove removes a space from the rtree
func Remove(sp *Space) {
	rt.Delete(sp)
}

// UpdateSpace resets a space's location to a given
// rtreego.Rect.
// This is not an operation on a space because
// a space can exist in multiple rtrees.
func UpdateSpace(x, y, w, h float64, s *Space) {
	loc := NewRect(x, y, w, h)
	rt.Delete(s)
	s.Location = loc
	rt.Insert(s)
}

// Hits returns the set of spaces which are colliding
// with the passed in space.
func Hits(sp *Space) []*Space {
	results := rt.SearchIntersect(sp.Bounds())
	out := make([]*Space, len(results))
	for index, v := range results {
		out[index] = v.(*Space)
	}
	return out
}

// HitLabel acts like hits, but reutrns the first space within hits
// that matches one of the input labels
func HitLabel(sp *Space, labels ...int) *Space {
	results := rt.SearchIntersect(sp.Bounds())
	for _, v := range results {
		for _, label := range labels {
			if v.(*Space) != sp && v.(*Space).Label == label {
				return v.(*Space)
			}
		}
	}
	return nil
}
