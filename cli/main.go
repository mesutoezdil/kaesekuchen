package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "make":
		cmdMake()
	case "audit":
		cmdAudit()
	case "classify":
		cmdClassify()
	case "help", "--help", "-h":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "kk: unknown command %q\n", os.Args[1])
		fmt.Fprintln(os.Stderr, "run 'kk help' for usage")
		os.Exit(1)
	}
}

func usage() {
	fmt.Print(`kk  Kaesekuchen CLI

commands:
    make        walk through the recipe, step by step
    audit       determine whether what you made is a Kaesekuchen
    classify    identify what kind of cheesecake you have

this tool cannot bake for you.
that part is yours.
`)
}
