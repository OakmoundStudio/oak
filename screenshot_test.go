package oak

import (
	"image"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/oakmound/oak/v3/scene"
)

func TestScreenShot(t *testing.T) {
	c1 := NewController()
	err := c1.SceneMap.AddScene("blank", scene.Scene{})
	if err != nil {
		t.Fatalf("Scene Add failed: %v", err)
	}
	go c1.Init("blank")
	time.Sleep(2 * time.Second)
	MatchScreenShot(t, c1, filepath.Join("testdata", "screenshot.png"))
}

func MatchScreenShot(t *testing.T, c *Controller, path string) {
	t.Helper()
	rgba := c.ScreenShot()
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open screenshot file: %v", err)
	}
	testRGBA, _, err := image.Decode(f)
	if err != nil {
		t.Fatalf("failed to decode screenshot file: %v", err)
	}
	bds := rgba.Bounds()
	if testRGBA.Bounds() != bds {
		t.Fatalf("mismatch screenshot size: got %v expected %v", bds, testRGBA.Bounds())
	}
	for x := bds.Min.X; x < bds.Max.X; x++ {
		for y := bds.Min.Y; y < bds.Max.Y; y++ {
			got := rgba.RGBAAt(x, y)
			gotR, gotG, gotB, gotA := got.RGBA()
			testGot := testRGBA.At(x, y)
			testR, testG, testB, testA := testGot.RGBA()
			if gotR != testR ||
				gotG != testG ||
				gotB != testB ||
				gotA != testA {
				t.Fatalf("pixel mismatch (%d,%d)", x, y)
			}
		}
	}
}