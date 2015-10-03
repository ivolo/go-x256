# go-x256

Finds the nearest [xterm 256 color code](http://www.calmar.ws/vim/256-xterm-24bit-rgb-color-chart.html) for the rgb inputs you provide. Go port of [node-x256](https://github.com/substack/node-x256).

### Quickstart

```go
import "github.com/ivolo/go-x256"

code := x256.Code(220,40,150)
// 162
```