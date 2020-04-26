# gameoflife
A toy implementation of John Conway's Game of Life in Go.

The primary purpose of this project is to create something I can collaborate on with my son. I am a novice in both golang and
graphics programming, so this was interesting to me. I hope that the rules.UpdateBoard interface is simple enough to be accessible to a
teenager, whilst also being interesting. We'll see...

Running the program will give you an empty 100x100 grid, and the game will be paused. The controls are very minimal:

* Left mouse button: populate pieces
* Right mouse button: delete pieces
* Space: toggle running/paused
* Right arrow: increase speed
* Left arrow: decrease speed
* Escape: exit

The master branch contains rules which simply populate a new piece at random on each tick. There is also a rulesimpl branch which contains
a simple implementation of the actual game of life rules.

PRs gratefully received which make it more idiomatic, or simpler. I'd consider more efficient, but given its primary purpose not at the
expense of complexity.

## Screenshot

![Screenshot showing running game of life](/docs/screenshot.png)
