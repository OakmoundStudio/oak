package static

import (
	"os"

	oak "github.com/oakmound/oak/v3"
	"github.com/oakmound/oak/v3/event"
	"github.com/oakmound/oak/v3/render"
	"github.com/oakmound/oak/v3/scene"
)

type Slide struct {
	Rs          *render.CompositeR
	ContinueKey string
	PrevKey     string
	transition  scene.Transition
	cont        bool
	prev        bool
	OnClick     func()
}

func (ss *Slide) Init() {
	oak.SetFullScreen(true)
	render.Draw(ss.Rs, 0)
	event.GlobalBind("KeyUp"+ss.ContinueKey, func(event.CID, interface{}) int {
		ss.cont = true
		return 0
	})
	event.GlobalBind("KeyUp"+ss.PrevKey, func(event.CID, interface{}) int {
		ss.prev = true
		return 0
	})
	event.GlobalBind("KeyUpEscape", func(event.CID, interface{}) int {
		os.Exit(0)
		return 0
	})
	if ss.OnClick != nil {
		event.GlobalBind("MousePress", func(event.CID, interface{}) int {
			ss.OnClick()
			return 0
		})
	}
}

func (ss *Slide) Continue() bool {
	return !ss.cont && !ss.prev
}

func (ss *Slide) Prev() bool {
	ret := ss.prev
	ss.prev = false
	ss.cont = false
	return ret
}

func (ss *Slide) Append(rs ...render.Renderable) {
	for _, r := range rs {
		ss.Rs.Append(r)
	}
}

func (ss *Slide) Transition() scene.Transition {
	return ss.transition
}

func NewSlide(rs ...render.Renderable) *Slide {
	return &Slide{
		Rs:          render.NewCompositeR(rs...),
		ContinueKey: "RightArrow",
		PrevKey:     "LeftArrow",
	}
}

func Transition(trans scene.Transition) SlideOption {
	return func(s *Slide) *Slide {
		s.transition = trans
		return s
	}
}

func Background(r render.Modifiable) SlideOption {
	return func(s *Slide) *Slide {
		s.Rs.Prepend(r)
		return s
	}
}

func ControlKeys(cont, prev string) SlideOption {
	return func(s *Slide) *Slide {
		s.ContinueKey = cont
		s.PrevKey = prev
		return s
	}
}

type SlideOption func(*Slide) *Slide

func NewSlideSet(n int, opts ...SlideOption) []*Slide {
	slides := make([]*Slide, n)
	for i := range slides {
		slides[i] = NewSlide()
		for _, opt := range opts {
			slides[i] = opt(slides[i])
		}
	}
	return slides
}
