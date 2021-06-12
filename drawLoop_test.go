package oak

import (
	"testing"
	"time"

	"github.com/oakmound/oak/v3/scene"
)

func BenchmarkDrawLoop(b *testing.B) {
	c1 := NewController()
	c1.AddScene("draw", scene.Scene{})
	go c1.Init("draw")
	// give the engine some time to start
	time.Sleep(5 * time.Second)
	// We don't want any regular ticks getting through
	c1.DrawTicker.SetTick(100 * time.Hour)

	b.ResetTimer()
	// This sees how fast the draw ticker will accept forced steps,
	// which won't be accepted until the draw loop itself pulls
	// from the draw ticker, which it only does after having drawn
	// the screen for a frame. This way we push the draw loop
	// to draw as fast as possible and measure that speed.
	for i := 0; i < b.N; i++ {
		c1.DrawTicker.ForceStep()
	}
}