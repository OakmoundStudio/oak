//+build !js

package render

import (
	"path/filepath"
	"testing"

	"github.com/oakmound/oak/v2/fileutil"
)

var (
	imgPath1    = filepath.Join("16", "jeremy.png")
	badImgPath1 = filepath.Join("16", "invalid.png")
)

func TestBatchLoad(t *testing.T) {
	fileutil.BindataDir = AssetDir
	fileutil.BindataFn = Asset
	if BatchLoad(filepath.Join("assets", "images")) != nil {
		t.Fatalf("batch load failed")
	}
	sh, err := GetSheet(imgPath1)
	if err != nil {
		t.Fatalf("get sheet failed: %v", err)
	}
	if len(sh.ToSprites()) != 8 {
		t.Fatalf("sheet did not contain 8 sprites")
	}
	_, err = loadSprite("dir", "dummy.jpg", 0)
	if err == nil {
		t.Fatalf("load sprite should have failed")
	}
	sp, err := GetSprite("dummy.gif")
	if sp != nil {
		t.Fatalf("get sprite should be nil")
	}
	if err == nil {
		t.Fatalf("get sprite should have failed")
	}
	sp, err = GetSprite(imgPath1)
	if sp == nil {
		t.Fatalf("get sprite failed")
	}
	if err != nil {
		t.Fatalf("get sprite failed: %v", err)
	}
	UnloadAll()
}

func TestSetAssetPath(t *testing.T) {
	fileutil.BindataDir = AssetDir
	fileutil.BindataFn = Asset
	_, err := LoadSheet(dir, imgPath1, 16, 16, 0)
	if err != nil {
		t.Fatalf("load sheet failed: %v", err)
	}
	UnloadAll()
	SetAssetPaths(wd)
	_, err = LoadSheet(dir, imgPath1, 16, 16, 0)
	if err == nil {
		t.Fatalf("load sheet should have failed")
	}
	UnloadAll()
	SetAssetPaths(
		filepath.Join(
			wd,
			"assets",
			"images"),
	)
	_, err = LoadSheet(dir, imgPath1, 16, 16, 0)
	if err != nil {
		t.Fatalf("load sheet failed: %v", err)
	}
	UnloadAll()

}

func TestBadSheetParams(t *testing.T) {
	fileutil.BindataDir = AssetDir
	fileutil.BindataFn = Asset

	_, err := LoadSheet(dir, imgPath1, 0, 16, 0)
	if err == nil {
		t.Fatalf("load sheet should have failed")
	}
	_, err = LoadSheet(dir, imgPath1, 16, 0, 0)
	if err == nil {
		t.Fatalf("load sheet should have failed")
	}
	_, err = LoadSheet(dir, imgPath1, 16, 16, -1)
	if err == nil {
		t.Fatalf("load sheet should have failed")
	}
	_, err = LoadSheet(dir, imgPath1, 16, 16, 1000)
	if err == nil {
		t.Fatalf("load sheet should have failed")
	}
}

func TestSheetStorage(t *testing.T) {
	fileutil.BindataDir = AssetDir
	fileutil.BindataFn = Asset

	if SheetIsLoaded(imgPath1) {
		t.Fatalf("sheets should not be loaded at startup")
	}
	_, err := GetSheet(imgPath1)
	if err == nil {
		t.Fatalf("get sheet should have failed")
	}
	_, err = LoadSheet(dir, imgPath1, 16, 16, 0)
	if err != nil {
		t.Fatalf("load sheet failed: %v", err)
	}
	if !SheetIsLoaded(imgPath1) {
		t.Fatalf("sheet did not load")
	}
	_, err = GetSheet(imgPath1)
	if err != nil {
		t.Fatalf("get sheet failed: %v", err)
	}
	UnloadAll()
}

func TestSheetUtility(t *testing.T) {

	fileutil.BindataDir = AssetDir
	fileutil.BindataFn = Asset
	_, err := LoadSprites(dir, imgPath1, 16, 16, 0)
	if err != nil {
		t.Fatalf("load sprites failed: %v", err)
	}
	_, err = LoadSprites(dir, badImgPath1, 16, 16, 0)
	if err == nil {
		t.Fatalf("load sprites should have failed")
	}

	_, err = LoadSheetSequence(imgPath1, 16, 16, 0, 1, 0, 0)
	if err != nil {
		t.Fatalf("load sprite sequence failed: %v", err)
	}
	_, err = LoadSheetSequence(badImgPath1, 16, 16, 0, 1, 0, 0)
	if err == nil {
		t.Fatalf("load sprite sequence should have failed")
	}

}
