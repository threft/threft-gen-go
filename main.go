package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

var options struct {
	Debugging bool     `short:"d" long:"debug" description:"Enable logging of debug messages to StdOut"`
	Protocols []string `short:"p" long:"protocol" description:"Protocol this genererator should include in generated code"`
}

func main() {
	args, err := flags.Parse(&options)
	if err != nil {
		flagError := err.(*flags.Error)
		if flagError.Type == flags.ErrHelp {
			fmt.Println("Example: threft-gen-go -p json -p binary")
			fmt.Println()
			return
		}
		if flagError.Type == flags.ErrUnknownFlag {
			fmt.Println("Use --help to view all available options.")
			return
		}
		fmt.Printf("Error parsing flags: %s\n", err)
		return
	}
	if len(args) > 0 {
		fmt.Printf("Unknown argument '%s'.", args[0])
		return
	}

	fmt.Println(options.Protocols)
}
