# godown

A talking countdown written in golang.

### Installation

Unfortunately this application has dependencies. Currently (2015-04-11) there is no os independent possiblity to play audio files in golang :(

Install [CSFML](http://www.sfml-dev.org/download/csfml/) (tested with `v2.3`).

#### Linux

Install [SFML](http://www.sfml-dev.org/download/sfml/2.3/). Then just download the [godown](https://github.com/MMore/godown/releases/download/v1.0.1/godown) binary and use it.

#### Mac OS X

Clone this repository to your machine and execute `make install_deps_osx`, which installs the dependencies for you. Then just just download the [godown](https://github.com/MMore/godown/releases/download/v1.0.1/godown) binary and use it.

If you already have a working CSFML and SFML, you can install it from source with `go get github.com/MMore/godown`.

### Usage
`godown [<flags>] <duration>`

Flags:
* --female   male or female voice
* --lang=en  english (en) or german (de) language

Args:
* <duration>  duration of the countdown (i.e. 1m30s)

### Tested Environments

* Mac OS X Yosemite and go1.4.1
* Mac OS X High Sierra and go1.8.1

### License
[LICENSE](LICENSE)

## Copyright

Made by Mathias Nestler (c) 2015. Feel free to contribute!

Sounds Copyright Epic games (Unreal Tournament 2004)
