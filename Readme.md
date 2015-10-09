# go-x256

Finds the nearest [xterm 256 color code](http://www.calmar.ws/vim/256-xterm-24bit-rgb-color-chart.html) for the rgb inputs you provide. Go port of [node-x256](https://github.com/substack/node-x256).

### Quickstart

Get the ansi color code closest to your image color:

```go
import "github.com/ivolo/go-x256"

// get the cloest ansi code to RGB point {220,40,150}
code := x256.ClosestCode(220,40,150)
// 162
```

and you can print a colored string:

```go
import "github.com/ivolo/go-x256"
import "github.com/mgutz/ansi"

/// get the ansi color code
code := x256.ClosestCode(220,40,150)

// write the colored string
buffer.WriteString(ansi.ColorCode(strconv.Itoa(code)))
buffer.WriteString("woo!")
buffer.WriteString(ansi.ColorCode("reset"))
```

Then [make something funny](https://github.com/ivolo/giffy): 

![](https://cloud.githubusercontent.com/assets/658544/10387602/9147f6be-6e17-11e5-8521-0f3d71fa3878.gif)