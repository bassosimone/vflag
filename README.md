# Golang Versatile CLI Flags Parser

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/vflag)](https://pkg.go.dev/github.com/bassosimone/vflag) [![Build Status](https://github.com/bassosimone/vflag/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/vflag/actions) [![codecov](https://codecov.io/gh/bassosimone/vflag/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/vflag)

The `vflag` Go package contains a versatile flags parser
similar to the stdlib `flag` package.

For example:

```Go
import (
	"os"

	"github.com/bassosimone/runtimex"
	"github.com/bassosimone/vflag"
)

// Create an empty flag set
fset := vflag.NewFlagSet("curl", vflag.ExitOnError)

// Edit the default values
usage := vflag.NewDefaultUsagePrinter()
usage.AddDescription("curl is an utility to transfer URLs.")
usage.AddExamples("curl -fsSL -o index.html https://example.com/")
usage.PositionalArgumentsUsage = "URL ..."
fset.SetMinMaxPositionalArgs(1, math.MaxInt)
fset.UsagePrinter = usage

// Add the supported flags
var (
	failFlag      = false
	locationFlag  = false
	outputFlag    = "-"
	showErrorFlag = false
	silentFlag    = false
)
fset.BoolVar(&failFlag, 'f', "fail", "Fail fast with no output at all on server errors.")
fset.BoolVar(&locationFlag, 'L', "location", "Follow HTTP redirections.")
fset.AutoHelp('h', "help", "Show this help message and exit.")
fset.StringVar(&outputFlag, 'o', "output", "Write output to the file indicated by VALUE.")
fset.BoolVar(&showErrorFlag, 'S', "show-error", "Show an error message, even when silent, on failure.")
fset.BoolVar(&silentFlag, 's', "silent", "Silent or quiet mode.")

// Invoke with command line arguments
runtimex.PanicOnError0(fset.Parse(os.Args[1:]))
```

The above example configures GNU style options but we support a
wide variety of command-line-flags styles including Go, dig, Windows,
and traditional Unix. See [example_test.go](example_test.go).

## Installation

To add this package as a dependency to your module:

```sh
go get github.com/bassosimone/vflag
```

## Development

To run the tests:
```sh
go test -v .
```

To measure test coverage:
```sh
go test -v -cover .
```

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```

## History

Adapted from [bassosimone/clip](https://github.com/bassosimone/clip/tree/v0.8.0).
