package main

import (
	"bufio"
	"fmt"
	"os"
)

type category struct {
	name    string
	verdict string
}

var categories = map[string]category{
	"KAESEKUCHEN": {
		name:    "Kaesekuchen",
		verdict: "this is a Kaesekuchen. this output is rare. handle with care.",
	},
	"QUARKKUCHEN": {
		name:    "Quarkkuchen",
		verdict: "this is a Quarkkuchen. a decent product. not a Kaesekuchen.",
	},
	"NEW_YORK": {
		name:    "New York Cheesecake",
		verdict: "this is a New York Cheesecake. it succeeded on its own terms. those terms are not Kaesekuchen.",
	},
	"SAN_SEBASTIAN": {
		name:    "Basque Burnt Cheesecake",
		verdict: "this is a Basque Burnt Cheesecake. the burnt surface is the point. in Kaesekuchen it would be an error.",
	},
	"JAPANESE": {
		name:    "Japanese Cheesecake",
		verdict: "this is a Japanese Cheesecake. impressive texture. not Kaesekuchen.",
	},
	"UNDEFINED": {
		name:    "undefined",
		verdict: "this product cannot be classified. review your decisions.",
	},
}

func cmdClassify() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Kaesekuchen classifier.")
	fmt.Println("Answer yes or no.")
	fmt.Println()

	result := classify(scanner)
	cat := categories[result]

	fmt.Println()
	fmt.Printf("classification: %s\n", cat.name)
	fmt.Println()
	fmt.Println(cat.verdict)
}

func classify(scanner *bufio.Scanner) string {
	if !yesNo(scanner, "Was it baked?") {
		return "UNDEFINED"
	}

	if yesNo(scanner, "Is the top deliberately burnt and the center liquid at serving temperature?") {
		return "SAN_SEBASTIAN"
	}

	if yesNo(scanner, "Is the texture cloud-soft, and does the cake wobble noticeably when touched?") {
		return "JAPANESE"
	}

	if yesNo(scanner, "Is the base made from Graham crackers, Digestives, or any crushed biscuit?") {
		return "NEW_YORK"
	}

	if yesNo(scanner, "Is the primary dairy ingredient Quark?") {
		return "QUARKKUCHEN"
	}

	if yesNo(scanner, "Does the cake have fruit in the filling or on top?") {
		return "UNDEFINED"
	}

	if !yesNo(scanner, "Is the primary dairy ingredient actual cheese, not Quark, not cream cheese, not a spreadable substitute?") {
		return "UNDEFINED"
	}

	return "KAESEKUCHEN"
}

func yesNo(scanner *bufio.Scanner, question string) bool {
	for {
		fmt.Printf("  %s [yes/no] ", question)
		scanner.Scan()
		switch scanner.Text() {
		case "yes", "y":
			return true
		case "no", "n":
			return false
		default:
			fmt.Println("  answer yes or no.")
		}
	}
}
