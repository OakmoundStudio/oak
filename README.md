# Oak

## A Pure Go game engine

[![Go Reference](https://pkg.go.dev/badge/github.com/oakmound/oak/v3.svg)](https://pkg.go.dev/github.com/oakmound/oak/v3)
[![Code Coverage](https://codecov.io/gh/oakmound/oak/branch/develop/graph/badge.svg)](https://codecov.io/gh/oakmound/oak)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go)

## Table of Contents

1. [Installation](#installation)

1. [Motivation](#motivation)

1. [Features](#features)

1. [Support](#support)

1. [Quick Start](#quick-start)

1. [Implementation and Examples](#examples)

1. [Finished Games](#finished-games)

***

## Installation <a name="installation"/>

`go get -u github.com/oakmound/oak/v3/`

## Motivation <a name="motivation"/>

The initial version of oak was made to support Oakmound Studio's game,
[Agent Blue](https://github.com/oakmound/AgentRelease), and was developed in parallel.

Because Oak wants to have as few non-Go dependencies as possible, on as many platforms as possible, Oak does not by default use bindings involving CGo like OpenGL or GLFW.

### On Pure Go

Oak has recently brought in dependencies that include C code, but we still describe the engine as a Pure Go engine, which at face value seems contradictory. Oak's goal is that, by default, a user can pull down the engine and create a fully functional game or GUI application on a machine with no C compiler installed, so when we say Pure Go we mean that, by default, the library is configured so no C compilation is required, and that no major features are locked behind C compliation.  

We anticipate in the immediate future needing to introduce alternate drivers that include C dependencies for performance improvements in some scenarios, and currently we have no OSX solution that lacks objective C code.

## Features and Systems <a name="features"></a>

1. Window Management
    - Windows and key events forked from [shiny](https://pkg.go.dev/golang.org/x/exp/shiny)
    - Logical frame rate distinct from Draw rate
    - Fullscreen, Window Positioning support, etc
    - Auto-scaling for screen size changes (or dynamic content sizing)
1. [Image Rendering](https://pkg.go.dev/github.com/oakmound/oak/v3/render)
    - `render.Renderable` interface
    - Sprite Sheet Batch Loading at startup
    - Manipulation
        - `render.Modifiable` interface
        - Built in Transformations and Filters
        - Some built-ins via [gift](https://github.com/disintegration/gift)
        - Extensible Modification syntax `func(image.Image) *image.RGBA`
    - Built in `Renderable` types covering common use cases
        - `Sprite`, `Sequence`, `Switch`, `Composite`
        - Primitive builders, `ColorBox`, `Line`, `Bezier`
        - History-tracking `Reverting`
    - Primarily 2D
1. [Particle System](https://pkg.go.dev/github.com/oakmound/oak/v3/render/particle)
1. [Mouse Handling](https://pkg.go.dev/github.com/oakmound/oak/v3/mouse)
1. [Joystick Support](https://pkg.go.dev/github.com/oakmound/oak/v3/joystick)
1. [Audio Support](https://pkg.go.dev/github.com/oakmound/oak/v3/audio)
    - Positional filters to pan and scale audio based on a listening position
1. [Collision](https://pkg.go.dev/github.com/oakmound/oak/v3/collision)
    - Collision R-Tree forked from [rtreego](https://github.com/dhconnelly/rtreego)
    - [2D Raycasting](https://pkg.go.dev/github.com/oakmound/oak/v3/collision/ray)
    - Collision Spaces
        - Attachable to Objects
        - Auto React to collisions through events
        - OnHit bindings `func(s1,s2 *collision.Space)`
        - Start/Stop collision with targeted objects
1. [2D Physics System](https://pkg.go.dev/github.com/oakmound/oak/v3/physics)
1. [Event Handler](https://pkg.go.dev/github.com/oakmound/oak/v3/event)
    - PubSub system: `event.CID` can `Bind(eventName,fn)` and `Trigger(eventName,payload)` events
1. [Shaping](https://pkg.go.dev/github.com/oakmound/oak/v3/shape)
    - Convert shapes into:
        - Containment checks
        - Outlines
        - 2D arrays
1. [Custom Console Commands](debugConsole.go)

## Support <a name="support"></a>

For discussions not significant enough to be an Issue or PR, feel free to ping us in the #oak channel on the [gophers slack](https://invite.slack.golangbridge.org/).

## Quick Start <a name="quick-start"></a>

This is an example of the most basic oak program:

```go
package main

import (
    "github.com/oakmound/oak/v3"
    "github.com/oakmound/oak/v3/scene"
)

func main() {
    oak.AddScene("firstScene", scene.Scene{
        Start: func(*scene.Context) {
            // ... draw entities, bind callbacks ... 
        }, 
    })
    oak.Init("firstScene")
}
```

See below or the [examples](examples) folder for longer demos, [godoc](https://pkg.go.dev/github.com/oakmound/oak/v3) for reference documentation, and the [wiki](https://github.com/oakmound/oak/wiki) for more guided feature sets, tutorials and walkthroughs.

## Implementation and Examples <a name="examples"></a>

| | | |
|:-------------------------:|:-------------------------:|:-------------------------:|
|<img width="1400"  src="examples/platformer-tutorial/6-complete/example.gif" a=examples/platformer-tutorial>  [Platformer](examples/platformer) |  <img width="1400"  src="examples/top-down-shooter-tutorial/6-performance/example.gif"> [Top down shooter](examples/top-down-shooter-tutorial)|<img width="1400"  src="examples/radar-demo/example.gif"> [Radar](examples/radar-demo) |
|<img width="1400"  src="examples/slide/example.gif"> [Slideshow](examples/slide) |  <img width="1400"  src="examples/bezier/example.PNG"> [Bezier Curves](examples/bezier) |<img width="1400"  src="examples/joystick-viz/example.gif"> [Joysticks](examples/joystick-viz)|
|<img width="1400"  src="examples/collision-demo/example.PNG"> [Collision Demo](examples/collision-demo)  |  <img width="1400"  src="examples/custom-cursor/example.PNG"> [Custom Mouse Cursor](examples/custom-cursor) |<img width="1400"  src="examples/fallback-font/example.PNG"> [Fallback Fonts](examples/fallback-font)| 
|<img width="1400"  src="examples/screenopts/example.PNG"> [Screen Options](examples/screenopts)  |  <img width="1400"  src="examples/multi-window/example.PNG"> [Multi Window](examples/multi-window) |<img width="1400"  src="examples/particle-demo/overviewExample.gif"> [Particle Demo](examples/particle-demo)| 

## Games using Oak <a name="finished-games"/>

| | |
|:-------------------------:|:-------------------------:|
|<img width="1400"  src="https://img.itch.zone/aW1hZ2UvMTk4MjIxLzkyNzUyOC5wbmc=/original/aRusLc.png" a=examples/platformer-tutorial>  [Agent Blue](https://oakmound.itch.io/agent-blue) |  <img width="1400"  src="https://img.itch.zone/aW1hZ2UvMTY4NDk1Lzc4MDk1Mi5wbmc=/original/hIjzFm.png"> [Fantastic Doctor](https://github.com/oakmound/lowrez17)
|<img width="1400"  src="https://img.itch.zone/aW1hZ2UvMzkwNjM5LzI2NzU0ODMucG5n/original/eaoFrd.png">  [Hiring Now: Looters](https://oakmound.itch.io/cheststacker) |  <img width="1400"  src="https://img.itch.zone/aW1hZ2UvMTYzNjgyLzc1NDkxOS5wbmc=/original/%2BwvZ7j.png"> [Jeremy The Clam](https://github.com/200sc/jeremy)
|<img width="1400"  src="https://img.itch.zone/aW1hZ2UvOTE0MjYzLzUxNjg3NDEucG5n/original/5btfEr.png">  [Diamond Deck Championship](https://oakmound.itch.io/diamond-deck-championship) |  
