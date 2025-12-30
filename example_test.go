//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/pkg/nflag/example_test.go
//

package vflag_test

import (
	"fmt"
	"math"
	"os"

	"github.com/bassosimone/flagparser"
	"github.com/bassosimone/vflag"
)

// This example shows the behavior when no flags are defined.
func ExampleFlagSet_noFlags() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("curl", vflag.ExitOnError)

	// Note: no flags have been configured

	// Override Exit to transform it into a panic
	fset.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Override Stderr to be the Stdout otherwise the testable example fails
	fset.Stderr = os.Stdout

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Invoke with `--verbose` which yields an error because `verbose` is not defined
	fset.Parse([]string{"--verbose"})

	// Output:
	// curl: unknown option: --verbose
}

// This example shows how we print the usage for a curl-like command line.
func ExampleFlagSet_curlHelp() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("curl", vflag.ExitOnError)

	// Edit the default values
	fset.AddDescription("curl is an utility to transfer URLs.")
	fset.AddExamples("curl -fsSL -o index.html https://example.com/")
	fset.PositionalArgumentsUsage = "URL ..."
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)

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

	// Override Exit to transform it into a panic
	fset.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Invoke with `--help`
	fset.Parse([]string{"--help"})

	// Output:
	// Usage
	//
	//     curl [flags] URL ...
	//
	// Description
	//
	//     curl is an utility to transfer URLs.
	//
	// Flags
	//
	//     -f, --fail[=BOOL] (default: `false`)
	//
	//         Fail fast with no output at all on server errors.
	//
	//     -L, --location[=BOOL] (default: `false`)
	//
	//         Follow HTTP redirections.
	//
	//     -h, --help
	//
	//         Show this help message and exit.
	//
	//     -o STRING, --output STRING (default: `-`)
	//
	//         Write output to the file indicated by VALUE.
	//
	//     -S, --show-error[=BOOL] (default: `false`)
	//
	//         Show an error message, even when silent, on failure.
	//
	//     -s, --silent[=BOOL] (default: `false`)
	//
	//         Silent or quiet mode.
	//
	// Examples
	//
	//     curl -fsSL -o index.html https://example.com/
}

// This example shows how we print errors when there are too few arguments.
func ExampleFlagSet_curlTooFewArguments() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("curl", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)

	// Add the supported flags
	fset.AutoHelp('h', "help", "Show this help message and exit.")

	// Override Exit to transform it into a panic
	fset.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Override Stderr to be the Stdout otherwise the testable example fails
	fset.Stderr = os.Stdout

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Invoked with not arguments at all
	fset.Parse([]string{})

	// Output:
	// curl: too few positional arguments: expected at least 1, got 0
	// hint: try `curl --help' for more help.
}

// This example shows a successful invocation of a curl-like tool.
func ExampleFlagSet_curlSuccess() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("curl", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)

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
	fset.Parse([]string{"-fsSLo", "index.html", "https://example.com/"})

	// Print the parsed flags
	fmt.Println("---")
	fmt.Printf("fail: %v\n", failFlag)
	fmt.Printf("location: %v\n", locationFlag)
	fmt.Printf("output: %s\n", outputFlag)
	fmt.Printf("show-error: %v\n", showErrorFlag)
	fmt.Printf("silent: %v\n", silentFlag)

	// Print the positional arguments
	fmt.Println("---")
	fmt.Printf("positional arguments: %v\n", fset.Args())

	// Output:
	// ---
	// fail: true
	// location: true
	// output: index.html
	// show-error: true
	// silent: true
	// ---
	// positional arguments: [https://example.com/]
}

// This example shows how we print the usage for a dig-like tool.
func ExampleFlagSet_digHelp() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("dig", vflag.ExitOnError)

	// Edit the default values
	fset.AddDescription("dig is an utility to query the domain name system.")
	fset.AddExamples("dig +short IN A -46 example.com")
	fset.PositionalArgumentsUsage = "[@server] [name] [type] [class]"
	fset.SetMinMaxPositionalArgs(0, 4)

	// Modify the long prefix to use dig conventions
	_fixLongOpt := func(fx *vflag.Flag) {
		fx.LongPrefix = "+"
	}

	// Add the supported flags
	//
	// Note: to support dig flags we need:
	//
	// 1. to fix `+short` to be a `+` introduced long option
	//
	// 2. a custom [*vflag.Flag] for `+https`
	var (
		httpsFlag = ""
		ipv4Flag  = false
		ipv6Flag  = false
		shortFlag = false
	)
	fset.BoolVar(&ipv4Flag, '4', "", "Enable using IPv4.")
	fset.BoolVar(&ipv6Flag, '6', "", "Enable using IPv6.")
	fset.AutoHelp('h', "", "Show this help message and exit.")
	fset.AddFlag(&vflag.Flag{
		Description: []string{
			"Enable using DNS-over-HTTPS with optional URL path.",
			"We use `/dns-query` if URL_PATH is omitted.",
		},
		LongArgumentName: "[=URL_PATH]",
		LongName:         "https",
		LongPrefix:       "+",
		MakeOptions: func(fx *vflag.Flag) []*flagparser.Option {
			return []*flagparser.Option{{
				DefaultValue: "/dns-query",
				Prefix:       fx.LongPrefix,
				Name:         fx.LongName,
				Type:         flagparser.OptionTypeStandaloneArgumentOptional,
			}}
		},
		Value: vflag.NewValueString(&httpsFlag),
	})
	_fixLongOpt(fset.BoolVar(&shortFlag, 0, "short", "Write terse output."))

	// Override Exit to transform it into a panic
	fset.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Invoke with `-h`
	fset.Parse([]string{"-h"})

	// Output:
	// Usage
	//
	//     dig [flags] [@server] [name] [type] [class]
	//
	// Description
	//
	//     dig is an utility to query the domain name system.
	//
	// Flags
	//
	//     -4 (default: `false`)
	//
	//         Enable using IPv4.
	//
	//     -6 (default: `false`)
	//
	//         Enable using IPv6.
	//
	//     -h
	//
	//         Show this help message and exit.
	//
	//     +https[=URL_PATH] (default: ``)
	//
	//         Enable using DNS-over-HTTPS with optional URL path.
	//
	//         We use `/dns-query` if URL_PATH is omitted.
	//
	//     +short[=BOOL] (default: `false`)
	//
	//         Write terse output.
	//
	// Examples
	//
	//     dig +short IN A -46 example.com
}

// This example shows how we print errors caused by invalid flags.
func ExampleFlagSet_digInvalidFlag() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("dig", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(0, 4)

	// Modify the long prefix to use dig conventions
	_fixLongOpt := func(fx *vflag.Flag) {
		fx.LongPrefix = "+"
	}

	// Add the supported flags
	var (
		ipv4Flag  = false
		ipv6Flag  = false
		shortFlag = false
	)
	fset.BoolVar(&ipv4Flag, '4', "", "Enable using IPv4.")
	fset.BoolVar(&ipv6Flag, '6', "", "Enable using IPv6.")
	fset.AutoHelp('h', "", "Show this help message and exit.")
	_fixLongOpt(fset.BoolVar(&shortFlag, 0, "short", "Write terse output."))

	// Override Exit to transform it into a panic
	fset.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Override Stderr to be the Stdout otherwise the testable example fails
	fset.Stderr = os.Stdout

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Invoke with a flag that has not been defined
	fset.Parse([]string{"+tls"})

	// Output:
	// dig: unknown option: +tls
	// hint: try `dig -h' for more help.
}

// This example shows a successful invocation of a dig-like tool.
func ExampleFlagSet_digSuccess() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("dig", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(0, 4)

	// Modify the long prefix to use dig conventions
	_fixLongOpt := func(fx *vflag.Flag) {
		fx.LongPrefix = "+"
	}

	// Add the supported flags
	//
	// Note: to support dig flags we need:
	//
	// 1. to fix `+short` to be a `+` introduced long option
	//
	// 2. a custom [*vflag.Flag] for `+https`
	var (
		httpsFlag = ""
		ipv4Flag  = false
		ipv6Flag  = false
		shortFlag = false
	)
	fset.BoolVar(&ipv4Flag, '4', "", "Enable using IPv4.")
	fset.BoolVar(&ipv6Flag, '6', "", "Enable using IPv6.")
	fset.AutoHelp('h', "", "Show this help message and exit.")
	fset.AddFlag(&vflag.Flag{
		Description: []string{
			"Enable using DNS-over-HTTPS with optional URL path.",
			"We use `/dns-query` if URL_PATH is omitted.",
		},
		LongArgumentName: "[=URL_PATH]",
		LongName:         "https",
		LongPrefix:       "+",
		MakeOptions: func(fx *vflag.Flag) []*flagparser.Option {
			return []*flagparser.Option{{
				DefaultValue: "/dns-query",
				Prefix:       fx.LongPrefix,
				Name:         fx.LongName,
				Type:         flagparser.OptionTypeStandaloneArgumentOptional,
			}}
		},
		Value: vflag.NewValueString(&httpsFlag),
	})
	_fixLongOpt(fset.BoolVar(&shortFlag, 0, "short", "Write terse output."))

	// Invoke with `-h`
	fset.Parse([]string{"IN", "A", "@8.8.8.8", "+https", "www.example.com", "+short", "-4"})

	// Print the parsed flags
	fmt.Println("---")
	fmt.Printf("httpsFlag: %v\n", httpsFlag)
	fmt.Printf("ipv4Flag: %v\n", ipv4Flag)
	fmt.Printf("ipv6Flag: %v\n", ipv6Flag)
	fmt.Printf("shortFlag: %v\n", shortFlag)

	// Print the positional arguments
	fmt.Println("---")
	fmt.Printf("positional arguments: %v\n", fset.Args())

	// Output:
	// ---
	// httpsFlag: /dns-query
	// ipv4Flag: true
	// ipv6Flag: false
	// shortFlag: true
	// ---
	// positional arguments: [IN A @8.8.8.8 www.example.com]
}

// This example shows how we print the usage for a tar-like tool.
func ExampleFlagSet_tarHelp() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("tar", vflag.ExitOnError)

	// Edit the default values
	fset.AddDescription("tar is an utility to manage possibly-compressed archives.")
	fset.AddExamples("tar -cvzf archive.tar.gz file1.txt file2.txt file3.txt")
	fset.PositionalArgumentsUsage = "[FILE ...]"
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)

	// Add the supported flags
	var (
		createFlag  = false
		fileFlag    = "-"
		gzipFlag    = false
		verboseFlag = false
	)
	fset.BoolVar(&createFlag, 'c', "", "Create a new archive.")
	fset.StringVar(&fileFlag, 'f', "", "Specify the output file path.")
	fset.AutoHelp('h', "help", "Show this help message and exit.")
	fset.BoolVar(&verboseFlag, 'v', "", "Print files added to the archive to the stdout.")
	fset.BoolVar(&gzipFlag, 'z', "", "Compress using gzip.")

	// Override Exit to transform it into a panic
	fset.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Invoke with `--help`
	fset.Parse([]string{"--help"})

	// Output:
	// Usage
	//
	//     tar [flags] [FILE ...]
	//
	// Description
	//
	//     tar is an utility to manage possibly-compressed archives.
	//
	// Flags
	//
	//     -c (default: `false`)
	//
	//         Create a new archive.
	//
	//     -f STRING (default: `-`)
	//
	//         Specify the output file path.
	//
	//     -h, --help
	//
	//         Show this help message and exit.
	//
	//     -v (default: `false`)
	//
	//         Print files added to the archive to the stdout.
	//
	//     -z (default: `false`)
	//
	//         Compress using gzip.
	//
	// Examples
	//
	//     tar -cvzf archive.tar.gz file1.txt file2.txt file3.txt
}

// This example shows how we print errors caused by a missing mandatory argument.
func ExampleFlagSet_tarMissingOptionArgument() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("tar", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)

	// Add the supported flags
	var (
		createFlag  = false
		fileFlag    = "-"
		gzipFlag    = false
		verboseFlag = false
	)
	fset.BoolVar(&createFlag, 'c', "", "Create a new archive.")
	fset.StringVar(&fileFlag, 'f', "", "Specify the output file path.")
	fset.AutoHelp('h', "help", "Show this help message and exit.")
	fset.BoolVar(&verboseFlag, 'v', "", "Print files added to the archive to the stdout.")
	fset.BoolVar(&gzipFlag, 'z', "", "Compress using gzip.")

	// Override Exit to transform it into a panic
	fset.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Override Stderr to be the Stdout otherwise the testable example fails
	fset.Stderr = os.Stdout

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Invoke missing argument for the `-f` option
	fset.Parse([]string{"-cvf"})

	// Output:
	// tar: option requires an argument: -f
	// hint: try `tar --help' for more help.
}

// This example shows how we print the usage for a go-like tool.
func ExampleFlagSet_goHelp() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("go test", vflag.ExitOnError)

	// Edit the default values
	fset.AddDescription("go test runs package tests.")
	fset.AddExamples("go test -race -count=1 -v ./...")
	fset.PositionalArgumentsUsage = "[package ...]"
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)

	// Modify the prefixes to use go conventions
	_fixOpt := func(fx *vflag.Flag) {
		fx.LongPrefix = "-"
		fx.ShortPrefix = ""
	}

	// Add the supported flags
	var (
		countFlag = int64(0)
		raceFlag  = false
		vFlag     = false
	)
	_fixOpt(fset.Int64Var(&countFlag, 0, "count", "Set to 1 to avoid using the test cache."))
	_fixOpt(fset.AutoHelp(0, "h", "Show this help message and exit."))
	_fixOpt(fset.AutoHelp(0, "help", "Alias for -h."))
	_fixOpt(fset.BoolVar(&raceFlag, 0, "race", "Run tests using the race detector."))
	_fixOpt(fset.BoolVar(&vFlag, 0, "v", "Print details about the tests progress and results."))

	// Override Exit to transform it into a panic
	fset.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Invoke with `-help`
	fset.Parse([]string{"-help"})

	// Output:
	// Usage
	//
	//     go test [flags] [package ...]
	//
	// Description
	//
	//     go test runs package tests.
	//
	// Flags
	//
	//     -count INT64 (default: `0`)
	//
	//         Set to 1 to avoid using the test cache.
	//
	//     -h
	//
	//         Show this help message and exit.
	//
	//     -help
	//
	//         Alias for -h.
	//
	//     -race[=BOOL] (default: `false`)
	//
	//         Run tests using the race detector.
	//
	//     -v[=BOOL] (default: `false`)
	//
	//         Print details about the tests progress and results.
	//
	// Examples
	//
	//     go test -race -count=1 -v ./...
}

// This example shows a successful invocation of a go-like tool.
func ExampleFlagSet_goSuccess() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("go test", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)

	// Modify the prefixes to use go conventions
	_fixOpt := func(fx *vflag.Flag) {
		fx.LongPrefix = "-"
		fx.ShortPrefix = ""
	}

	// Add the supported flags
	var (
		countFlag = int64(0)
		raceFlag  = false
		vFlag     = false
	)
	_fixOpt(fset.Int64Var(&countFlag, 0, "count", "Set to 1 to avoid using the test cache."))
	_fixOpt(fset.AutoHelp(0, "h", "Show this help message and exit."))
	_fixOpt(fset.AutoHelp(0, "help", "Alias for -h."))
	_fixOpt(fset.BoolVar(&raceFlag, 0, "race", "Run tests using the race detector."))
	_fixOpt(fset.BoolVar(&vFlag, 0, "v", "Print details about the tests progress and results."))

	// Invoke with command line arguments.
	//
	// Note that `-count=1` is equivalent to [`-count`, `1`].
	fset.Parse([]string{"-race", "-count", "1", "-v", "./..."})

	// Print the parsed flags
	fmt.Println("---")
	fmt.Printf("countFlag: %v\n", countFlag)
	fmt.Printf("raceFlag: %v\n", raceFlag)
	fmt.Printf("vFlag: %v\n", vFlag)

	// Print the positional arguments
	fmt.Println("---")
	fmt.Printf("positional arguments: %v\n", fset.Args())

	// Output:
	// ---
	// countFlag: 1
	// raceFlag: true
	// vFlag: true
	// ---
	// positional arguments: [./...]
}
