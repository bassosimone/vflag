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

// This example shows how we can customize the usage for a curl-like command.
func ExampleFlagSet_curlHelpCustom() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("curl", vflag.ExitOnError)

	// Edit the default values
	usage := vflag.NewDefaultUsagePrinter()
	usage.AddDescription("curl is an utility to transfer URLs.")
	usage.AddExamples(
		"Fetch https://example.com/ and store the results at index.html:",
		"    curl -fsSL -o index.html https://example.com/",
		"Same as above but emit to stdout implicitly:",
		"    curl -fsSL https://example.com/",
		"Same as above but emit to stdout explicitly using `-`:",
		"    curl -fsSL -o- https://example.com/",
	)
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

	fset.BoolVar(
		&failFlag,
		'f',
		"fail",
		"Fail fast with no output at all on server errors.",
		"Default: `@DEFAULT_VALUE@`.",
	)

	fset.BoolVar(&locationFlag, 'L', "location", "Follow HTTP redirections.", "Default: `@DEFAULT_VALUE@`.")

	fset.AutoHelp('h', "help", "Show this help message and exit.")

	fset.StringVar(
		&outputFlag,
		'o',
		"output",
		"Write output to the given `FILE`.",
		"Use `-` to explicitly indicate the stdout.",
		"Default: `@DEFAULT_VALUE@`.",
	)

	fset.BoolVar(
		&showErrorFlag,
		'S',
		"show-error",
		"Show an error message, even when silent, on failure.",
		"Default: `@DEFAULT_VALUE@`.",
	)

	fset.BoolVar(&silentFlag, 's', "silent", "Silent or quiet mode.", "Default: `@DEFAULT_VALUE@`.")

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
	//     -f, --fail[=true|false]
	//
	//         Fail fast with no output at all on server errors.
	//
	//         Default: `false`.
	//
	//     -L, --location[=true|false]
	//
	//         Follow HTTP redirections.
	//
	//         Default: `false`.
	//
	//     -h, --help
	//
	//         Show this help message and exit.
	//
	//     -o FILE, --output FILE
	//
	//         Write output to the given `FILE`.
	//
	//         Use `-` to explicitly indicate the stdout.
	//
	//         Default: `-`.
	//
	//     -S, --show-error[=true|false]
	//
	//         Show an error message, even when silent, on failure.
	//
	//         Default: `false`.
	//
	//     -s, --silent[=true|false]
	//
	//         Silent or quiet mode.
	//
	//         Default: `false`.
	//
	// Examples
	//
	//     Fetch https://example.com/ and store the results at index.html:
	//
	//         curl -fsSL -o index.html https://example.com/
	//
	//     Same as above but emit to stdout implicitly:
	//
	//         curl -fsSL https://example.com/
	//
	//     Same as above but emit to stdout explicitly using `-`:
	//
	//         curl -fsSL -o- https://example.com/
}

// This example shows how we print the default usage for a curl-like command.
func ExampleFlagSet_curlHelpDefault() {
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
	fset.StringVar(&outputFlag, 'o', "output", "Write output to the given `FILE`.")
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
	//     curl [flags] arg [arg ...]
	//
	// Flags
	//
	//     -f, --fail[=true|false]
	//
	//         Fail fast with no output at all on server errors.
	//
	//     -L, --location[=true|false]
	//
	//         Follow HTTP redirections.
	//
	//     -h, --help
	//
	//         Show this help message and exit.
	//
	//     -o FILE, --output FILE
	//
	//         Write output to the given `FILE`.
	//
	//     -S, --show-error[=true|false]
	//
	//         Show an error message, even when silent, on failure.
	//
	//     -s, --silent[=true|false]
	//
	//         Silent or quiet mode.
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
	// curl: try `curl --help' for more help.
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
	fset.StringVar(&outputFlag, 'o', "output", "Write output to the given `FILE`.")
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

// This example shows how we can customize the usage for a dig-like tool.
func ExampleFlagSet_digHelpCustom() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("dig", vflag.ExitOnError)

	// Edit the default values
	usage := vflag.NewDefaultUsagePrinter()
	usage.AddDescription("dig is an utility to query the domain name system.")
	usage.AddExamples("dig +short IN A -46 example.com")
	usage.PositionalArgumentsUsage = "[@server] name [type] [class]"
	fset.SetMinMaxPositionalArgs(1, 4)
	fset.UsagePrinter = usage

	// Add the supported flags
	var (
		httpsFlag = "/dns-query"
		ipv4Flag  = false
		ipv6Flag  = false
		shortFlag = false
	)

	// Short-only flags for -4, -6, -h using AddShortFlag
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&ipv4Flag), '4', "Enable using IPv4."))
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&ipv6Flag), '6', "Enable using IPv6."))
	fset.AddShortFlag(vflag.NewShortFlagAutoHelp(vflag.ValueAutoHelp{}, 'h', "Show this help message and exit."))

	// Long-only flag for +https with optional value using AddLongFlagDig
	//
	// Note: the backtick syntax (e.g., `URL_PATH`) in the first description paragraph
	// overrides the default ArgumentName in help output. We also use @DEFAULT_VALUE@
	// to show the default value in the help text.
	httpsValue := vflag.NewValueString(&httpsFlag)
	fset.AddLongFlagDig(&vflag.LongFlag{
		Description:  []string{"Enable using DNS-over-HTTPS with optional `URL_PATH`.", "Default: @DEFAULT_VALUE@."},
		ArgumentName: "[=STRING]",
		DefaultValue: httpsValue.String(),
		Name:         "https",
		MakeOption:   vflag.LongFlagMakeOptionWithOptionalValue,
		Prefix:       "--",
		Value:        httpsValue,
	})

	// Long-only flag for +short using AddLongFlagDig
	fset.AddLongFlagDig(vflag.NewLongFlagBool(
		vflag.NewValueBool(&shortFlag), "short",
		"Write terse output.",
		"Default: @DEFAULT_VALUE@.",
	))

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
	//     dig [flags] [@server] name [type] [class]
	//
	// Description
	//
	//     dig is an utility to query the domain name system.
	//
	// Flags
	//
	//     -4
	//
	//         Enable using IPv4.
	//
	//     -6
	//
	//         Enable using IPv6.
	//
	//     -h
	//
	//         Show this help message and exit.
	//
	//     +https[=URL_PATH]
	//
	//         Enable using DNS-over-HTTPS with optional `URL_PATH`.
	//
	//         Default: /dns-query.
	//
	//     +short[=true|false]
	//
	//         Write terse output.
	//
	//         Default: false.
	//
	// Examples
	//
	//     dig +short IN A -46 example.com
}

// This example shows how we print the default usage for a dig-like tool.
func ExampleFlagSet_digHelpDefault() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("dig", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(1, 4)

	// Add the supported flags
	var (
		httpsFlag = "/dns-query"
		ipv4Flag  = false
		ipv6Flag  = false
		shortFlag = false
	)

	// Short-only flags for -4, -6, -h using AddShortFlag
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&ipv4Flag), '4', "Enable using IPv4."))
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&ipv6Flag), '6', "Enable using IPv6."))
	fset.AddShortFlag(vflag.NewShortFlagAutoHelp(vflag.ValueAutoHelp{}, 'h', "Show this help message and exit."))

	// Long-only flag for +https with optional value using AddLongFlagDig
	httpsValue := vflag.NewValueString(&httpsFlag)
	fset.AddLongFlagDig(&vflag.LongFlag{
		Description:  []string{"Enable using DNS-over-HTTPS with optional `URL_PATH`.", "Default: @DEFAULT_VALUE@."},
		ArgumentName: "[=STRING]",
		DefaultValue: httpsValue.String(),
		Name:         "https",
		MakeOption:   vflag.LongFlagMakeOptionWithOptionalValue,
		Prefix:       "--",
		Value:        httpsValue,
	})

	// Long-only flag for +short using AddLongFlagDig
	fset.AddLongFlagDig(vflag.NewLongFlagBool(
		vflag.NewValueBool(&shortFlag), "short", "Write terse output.",
	))

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
	//     dig [flags] arg [arg ...]
	//
	// Flags
	//
	//     -4
	//
	//         Enable using IPv4.
	//
	//     -6
	//
	//         Enable using IPv6.
	//
	//     -h
	//
	//         Show this help message and exit.
	//
	//     +https[=URL_PATH]
	//
	//         Enable using DNS-over-HTTPS with optional `URL_PATH`.
	//
	//         Default: /dns-query.
	//
	//     +short[=true|false]
	//
	//         Write terse output.
}

// This example shows how we print errors caused by invalid flags.
func ExampleFlagSet_digInvalidFlag() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("dig", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(0, 4)

	// Add the supported flags
	var (
		ipv4Flag  = false
		ipv6Flag  = false
		shortFlag = false
	)

	// Short-only flags for -4, -6, -h using AddShortFlag
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&ipv4Flag), '4', "Enable using IPv4."))
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&ipv6Flag), '6', "Enable using IPv6."))
	fset.AddShortFlag(vflag.NewShortFlagAutoHelp(vflag.ValueAutoHelp{}, 'h', "Show this help message and exit."))

	// Long-only flag for +short using AddLongFlagDig
	fset.AddLongFlagDig(vflag.NewLongFlagBool(
		vflag.NewValueBool(&shortFlag), "short", "Write terse output.",
	))

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
	// dig: try `dig -h' for more help.
}

// This example shows a successful invocation of a dig-like tool.
func ExampleFlagSet_digSuccess() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("dig", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(0, 4)

	// Add the supported flags
	var (
		httpsFlag = "/dns-query"
		ipv4Flag  = false
		ipv6Flag  = false
		shortFlag = false
	)

	// Short-only flags for -4, -6, -h using AddShortFlag
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&ipv4Flag), '4', "Enable using IPv4."))
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&ipv6Flag), '6', "Enable using IPv6."))
	fset.AddShortFlag(vflag.NewShortFlagAutoHelp(vflag.ValueAutoHelp{}, 'h', "Show this help message and exit."))

	// Long-only flag for +https with optional value using AddLongFlagDig
	httpsValue := vflag.NewValueString(&httpsFlag)
	fset.AddLongFlagDig(&vflag.LongFlag{
		Description:  []string{"Enable using DNS-over-HTTPS with optional `URL_PATH`.", "Default: @DEFAULT_VALUE@."},
		ArgumentName: "[=STRING]",
		DefaultValue: httpsValue.String(),
		Name:         "https",
		MakeOption:   vflag.LongFlagMakeOptionWithOptionalValue,
		Prefix:       "--",
		Value:        httpsValue,
	})

	// Long-only flag for +short using AddLongFlagDig
	fset.AddLongFlagDig(vflag.NewLongFlagBool(
		vflag.NewValueBool(&shortFlag), "short", "Write terse output.",
	))

	// Invoke with command line arguments
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

// This example shows how we can customize the usage for a tar-like tool.
func ExampleFlagSet_tarHelpCustom() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("tar", vflag.ExitOnError)

	// Edit the default values
	usage := vflag.NewDefaultUsagePrinter()
	usage.AddDescription("tar is an utility to manage possibly-compressed archives.")
	usage.AddExamples("tar -cvzf archive.tar.gz file1.txt file2.txt file3.txt")
	usage.PositionalArgumentsUsage = "FILE ..."
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)
	fset.UsagePrinter = usage

	// Add the supported flags
	var (
		createFlag  = false
		fileFlag    = "-"
		gzipFlag    = false
		verboseFlag = false
	)

	// Short-only flags (tar style) using AddShortFlag
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&createFlag), 'c', "Create a new archive."))
	fset.AddShortFlag(vflag.NewShortFlagString(vflag.NewValueString(&fileFlag), 'f', "Write to `FILE`.", "Default: @DEFAULT_VALUE@."))
	fset.AutoHelp('h', "help", "Show this help message and exit.")
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&verboseFlag), 'v', "Print files added to the archive to the stdout."))
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&gzipFlag), 'z', "Compress using gzip."))

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
	//     tar [flags] FILE ...
	//
	// Description
	//
	//     tar is an utility to manage possibly-compressed archives.
	//
	// Flags
	//
	//     -c
	//
	//         Create a new archive.
	//
	//     -f FILE
	//
	//         Write to `FILE`.
	//
	//         Default: -.
	//
	//     -h, --help
	//
	//         Show this help message and exit.
	//
	//     -v
	//
	//         Print files added to the archive to the stdout.
	//
	//     -z
	//
	//         Compress using gzip.
	//
	// Examples
	//
	//     tar -cvzf archive.tar.gz file1.txt file2.txt file3.txt
}

// This example shows how we print the default usage for a tar-like tool.
func ExampleFlagSet_tarHelpDefault() {
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

	// Short-only flags (tar style) using AddShortFlag
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&createFlag), 'c', "Create a new archive."))
	fset.AddShortFlag(vflag.NewShortFlagString(vflag.NewValueString(&fileFlag), 'f', "Write to `FILE`."))
	fset.AutoHelp('h', "help", "Show this help message and exit.")
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&verboseFlag), 'v', "Print files added to the archive to the stdout."))
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&gzipFlag), 'z', "Compress using gzip."))

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
	//     tar [flags] arg [arg ...]
	//
	// Flags
	//
	//     -c
	//
	//         Create a new archive.
	//
	//     -f FILE
	//
	//         Write to `FILE`.
	//
	//     -h, --help
	//
	//         Show this help message and exit.
	//
	//     -v
	//
	//         Print files added to the archive to the stdout.
	//
	//     -z
	//
	//         Compress using gzip.
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

	// Short-only flags (tar style) using AddShortFlag
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&createFlag), 'c', "Create a new archive."))
	fset.AddShortFlag(vflag.NewShortFlagString(vflag.NewValueString(&fileFlag), 'f', "Write to `FILE`."))
	fset.AutoHelp('h', "help", "Show this help message and exit.")
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&verboseFlag), 'v', "Print files added to the archive to the stdout."))
	fset.AddShortFlag(vflag.NewShortFlagBool(vflag.NewValueBool(&gzipFlag), 'z', "Compress using gzip."))

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
	// tar: try `tar --help' for more help.
}

// This example shows how we can customize the usage for a go-like tool.
func ExampleFlagSet_goHelpCustom() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("go test", vflag.ExitOnError)

	// Edit the default values
	usage := vflag.NewDefaultUsagePrinter()
	usage.AddDescription("go test runs package tests.")
	usage.AddExamples("go test -race -count=1 -v ./...")
	usage.PositionalArgumentsUsage = "package ..."
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)
	fset.UsagePrinter = usage

	// Add the supported flags
	//
	// Go-style uses `-` prefix for long flags and no short flags
	var (
		countFlag = int64(0)
		raceFlag  = false
		vFlag     = false
	)

	// Long-only flags with `-` prefix (Go style) using AddLongFlag
	countLongFlag := vflag.NewLongFlagInt64(
		vflag.NewValueInt64(&countFlag), "count", "Run tests `N` times.", "Default: @DEFAULT_VALUE@.",
	)
	countLongFlag.Prefix = "-"
	fset.AddLongFlag(countLongFlag)

	hLongFlag := vflag.NewLongFlagAutoHelp(vflag.ValueAutoHelp{}, "h", "Show this help message and exit.")
	hLongFlag.Prefix = "-"
	fset.AddLongFlag(hLongFlag)

	helpLongFlag := vflag.NewLongFlagAutoHelp(vflag.ValueAutoHelp{}, "help", "Alias for -h.")
	helpLongFlag.Prefix = "-"
	fset.AddLongFlag(helpLongFlag)

	raceLongFlag := vflag.NewLongFlagBool(
		vflag.NewValueBool(&raceFlag), "race", "Run tests using the race detector.", "Default: @DEFAULT_VALUE@.",
	)
	raceLongFlag.Prefix = "-"
	fset.AddLongFlag(raceLongFlag)

	vLongFlag := vflag.NewLongFlagBool(
		vflag.NewValueBool(&vFlag), "v", "Verbose output.", "Default: @DEFAULT_VALUE@.",
	)
	vLongFlag.Prefix = "-"
	fset.AddLongFlag(vLongFlag)

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
	//     go test [flags] package ...
	//
	// Description
	//
	//     go test runs package tests.
	//
	// Flags
	//
	//     -count N
	//
	//         Run tests `N` times.
	//
	//         Default: 0.
	//
	//     -h
	//
	//         Show this help message and exit.
	//
	//     -help
	//
	//         Alias for -h.
	//
	//     -race[=true|false]
	//
	//         Run tests using the race detector.
	//
	//         Default: false.
	//
	//     -v[=true|false]
	//
	//         Verbose output.
	//
	//         Default: false.
	//
	// Examples
	//
	//     go test -race -count=1 -v ./...
}

// This example shows how we print the default usage for a go-like tool.
func ExampleFlagSet_goHelpDefault() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("go test", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)

	// Add the supported flags
	var (
		countFlag = int64(0)
		raceFlag  = false
		vFlag     = false
	)

	// Long-only flags with `-` prefix (Go style) using AddLongFlag
	countLongFlag := vflag.NewLongFlagInt64(
		vflag.NewValueInt64(&countFlag), "count", "Run tests `N` times.",
	)
	countLongFlag.Prefix = "-"
	fset.AddLongFlag(countLongFlag)

	hLongFlag := vflag.NewLongFlagAutoHelp(vflag.ValueAutoHelp{}, "h", "Show this help message and exit.")
	hLongFlag.Prefix = "-"
	fset.AddLongFlag(hLongFlag)

	helpLongFlag := vflag.NewLongFlagAutoHelp(vflag.ValueAutoHelp{}, "help", "Alias for -h.")
	helpLongFlag.Prefix = "-"
	fset.AddLongFlag(helpLongFlag)

	raceLongFlag := vflag.NewLongFlagBool(
		vflag.NewValueBool(&raceFlag), "race", "Run tests using the race detector.",
	)
	raceLongFlag.Prefix = "-"
	fset.AddLongFlag(raceLongFlag)

	vLongFlag := vflag.NewLongFlagBool(
		vflag.NewValueBool(&vFlag), "v", "Verbose output.",
	)
	vLongFlag.Prefix = "-"
	fset.AddLongFlag(vLongFlag)

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
	//     go test [flags] arg [arg ...]
	//
	// Flags
	//
	//     -count N
	//
	//         Run tests `N` times.
	//
	//     -h
	//
	//         Show this help message and exit.
	//
	//     -help
	//
	//         Alias for -h.
	//
	//     -race[=true|false]
	//
	//         Run tests using the race detector.
	//
	//     -v[=true|false]
	//
	//         Verbose output.
}

// This example shows a successful invocation of a go-like tool.
func ExampleFlagSet_goSuccess() {
	// Create an empty flag set
	fset := vflag.NewFlagSet("go test", vflag.ExitOnError)

	// Edit the default values
	fset.SetMinMaxPositionalArgs(1, math.MaxInt)

	// Add the supported flags
	var (
		countFlag = int64(0)
		raceFlag  = false
		vFlag     = false
	)

	// Long-only flags with `-` prefix (Go style) using AddLongFlag
	countLongFlag := vflag.NewLongFlagInt64(
		vflag.NewValueInt64(&countFlag), "count", "Run tests `N` times.",
	)
	countLongFlag.Prefix = "-"
	fset.AddLongFlag(countLongFlag)

	hLongFlag := vflag.NewLongFlagAutoHelp(vflag.ValueAutoHelp{}, "h", "Show this help message and exit.")
	hLongFlag.Prefix = "-"
	fset.AddLongFlag(hLongFlag)

	helpLongFlag := vflag.NewLongFlagAutoHelp(vflag.ValueAutoHelp{}, "help", "Alias for -h.")
	helpLongFlag.Prefix = "-"
	fset.AddLongFlag(helpLongFlag)

	raceLongFlag := vflag.NewLongFlagBool(
		vflag.NewValueBool(&raceFlag), "race", "Run tests using the race detector.",
	)
	raceLongFlag.Prefix = "-"
	fset.AddLongFlag(raceLongFlag)

	vLongFlag := vflag.NewLongFlagBool(
		vflag.NewValueBool(&vFlag), "v", "Verbose output.",
	)
	vLongFlag.Prefix = "-"
	fset.AddLongFlag(vLongFlag)

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
