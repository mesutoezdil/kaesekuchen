package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type violation struct {
	code    string
	message string
}

func ask(scanner *bufio.Scanner, question string) string {
	fmt.Print(question + " ")
	scanner.Scan()
	return strings.ToLower(strings.TrimSpace(scanner.Text()))
}

func cmdAudit() {
	scanner := bufio.NewScanner(os.Stdin)
	var violations []violation

	fmt.Println("Kaesekuchen audit.")
	fmt.Println("Answer honestly. The result reflects your cake, not your intentions.")
	fmt.Println()

	dairy := ask(scanner, "Primary dairy ingredient (cheese / quark / other):")
	switch {
	case dairy == "quark":
		violations = append(violations, violation{
			code:    "QUARK_AS_PRIMARY_DAIRY",
			message: "the name is Kaesekuchen. Käse means cheese. this is not a preference.",
		})
	case dairy == "cheese" || dairy == "käse":
		// correct
	default:
		violations = append(violations, violation{
			code:    "UNRECOGNIZED_DAIRY",
			message: fmt.Sprintf("%q is not a valid primary ingredient for Kaesekuchen", dairy),
		})
	}

	if !yesNo(scanner, "Was it baked?") {
		violations = append(violations, violation{
			code:    "NOT_BAKED",
			message: "a Kaesekuchen must be baked. gelatine is not heat.",
		})
	}

	if yesNo(scanner, "Is the base a Graham cracker, Digestive, or any other biscuit crust?") {
		violations = append(violations, violation{
			code:    "BISCUIT_BASE",
			message: "processed crumb substrates belong to American cheesecake traditions. they should remain there.",
		})
	}

	if yesNo(scanner, "Does it contain fruit, or is there fruit on top?") {
		violations = append(violations, violation{
			code:    "FRUIT_PRESENT",
			message: "fruit does not belong in or on a Kaesekuchen. if the cheese needed help, find better cheese.",
		})
	}

	if !yesNo(scanner, "Was it served cold?") {
		violations = append(violations, violation{
			code:    "WRONG_TEMPERATURE",
			message: "Kaesekuchen is served cold. the structure of the filling is designed for cold.",
		})
	}

	fmt.Println()

	if len(violations) == 0 {
		fmt.Println("audit passed.")
		fmt.Println()
		fmt.Println("this is a Kaesekuchen.")
		return
	}

	fmt.Printf("audit failed. %d violation(s):\n", len(violations))
	fmt.Println()
	for _, v := range violations {
		fmt.Printf("  [%s]\n", v.code)
		fmt.Printf("  %s\n", v.message)
		fmt.Println()
	}
	fmt.Println("this is not a Kaesekuchen.")
}
