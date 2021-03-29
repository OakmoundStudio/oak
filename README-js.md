# JS development notes

3/29/2021

shiny branch - `feature/js-try-3`

There's two ways of doing JS development:

- GopherJS
- syscall/js

GopherJS has worked in the past (see pong), but it is stuck at Go 1.12, and we want to move
towards smart `embed` support instead of `go-bindata`, so we can't go back to Go 1.12 in any
capacity-- its especially a problem with the concept of file access as JS absolutely needs file
embedding (it can't open OS files easily).

syscall/js is supposed to be the same thing as GopherJS but it panics constantly. An
implementation using it and a helper `dom` package looks like it should work to me but can't
advance past one frame.