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

	dairy := ask(scanner, "Primary dairy ingredient (cheese, quark, or other):")
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

	baked := ask(scanner, "Was it baked?")
	if baked != "yes" {
		violations = append(violations, violation{
			code:    "NOT_BAKED",
			message: "a Kaesekuchen must be baked. gelatine is not heat.",
		})
	}

	base := ask(scanner, "Base type (none, Mürbeteig, Karottenkuchen, Graham cracker, biscuit, or other):")
	switch base {
	case "none", "mürbeteig", "karottenkuchen":
		// acceptable
	case "graham cracker", "graham", "biscuit", "digestive", "keks", "cracker":
		violations = append(violations, violation{
			code:    "BISCUIT_BASE",
			message: "processed crumb substrates belong to American cheesecake traditions. they should remain there.",
		})
	default:
		if base != "" {
			violations = append(violations, violation{
				code:    "UNRECOGNIZED_BASE",
				message: fmt.Sprintf("%q is not a recognized base for Kaesekuchen", base),
			})
		}
	}

	fruit := ask(scanner, "Does it contain fruit, or is there fruit on top?")
	if fruit == "yes" {
		violations = append(violations, violation{
			code:    "FRUIT_PRESENT",
			message: "fruit does not belong in or on a Kaesekuchen. if the cheese needed help, find better cheese.",
		})
	}

	temp := ask(scanner, "Serving temperature (cold, room temperature, or warm):")
	if temp == "warm" || temp == "room temperature" {
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
