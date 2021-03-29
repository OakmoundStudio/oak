module github.com/oakmound/oak/v2

require (
	github.com/200sc/go-dist v1.0.0
	github.com/200sc/klangsynthese v0.2.2-0.20201022002431-a0e14a8c862b
	github.com/BurntSushi/toml v0.3.1
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/disintegration/gift v1.2.0
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/hajimehoshi/go-mp3 v0.3.1 // indirect
	github.com/oakmound/libudev v0.2.1
	github.com/oakmound/shiny v0.4.3-0.20210328180751-4c942f7e9c15
	github.com/oakmound/w32 v1.0.1-0.20210323130255-ae527b9640fd
	github.com/oov/directsound-go v0.0.0-20141101201356-e53e59c700bf // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/yobert/alsa v0.0.0-20200618200352-d079056f5370 // indirect
	golang.org/x/image v0.0.0-20201208152932-35266b937fa6
	golang.org/x/mobile v0.0.0-20190719004257-d2bd2a29d028
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
)

go 1.16

replace github.com/oakmound/shiny => ../shiny
