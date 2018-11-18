TinyGo examples
=======
Some examples of [TinyGo](https://github.com/aykevl/tinygo).

## What is TinyGo?

[TinyGo](https://tinygo.org) is a project to bring Go to microcontrollers and small systems with a single processor core. It is similar to [emgo](https://github.com/ziutek/emgo) but a major difference is that the project aims to keep the Go memory model (which implies garbage collection of some sort). Another difference is that TinyGo uses LLVM internally instead of emitting C, which hopefully leads to smaller and more efficient code and certainly leads to more flexibility.

## Usage


#### BBC micro:bit
You should be able to use the docker image of tinygo to compile the examples:

```bash
docker run --rm -v $(pwd):/go/src/github.com/conejoninja/tinygoexamples tinygo/tinygo build -o /go/src/github.com/conejoninja/tinygoexamples/bin.hex -target=microbit github.com/conejoninja/tinygoexamples/microbit/scrolltext

```

Then copy the _bin.hex_ to your BBC **micro:bit**

## Documentation of TinyGo

Documentation is currently maintained on a [dedicated ReadTheDocs
page](https://tinygo.readthedocs.io/en/latest/).

## Getting help

If you're looking for a more interactive way to discuss TinyGo usage or
development, there's a [#TinyGo channel](https://gophers.slack.com/messages/CDJD3SUP6/)
on the [Gophers Slack](https://gophers.slack.com).

If you need an invitation for the Gophers Slack, you can generate one here which
should arrive fairly quickly (under 1 min): https://invite.slack.golangbridge.org

## License of these examples

The MIT License (MIT)

Copyright 2018 Daniel Esteban - conejo@conejo.me

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

