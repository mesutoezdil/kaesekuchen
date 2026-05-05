package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type step struct {
	instruction string
	warning     string
}

var recipe = []step{
	{
		instruction: "Take your eggs out of the refrigerator. Wait until they reach room temperature.",
		warning:     "cold eggs will affect the texture. this step is not optional.",
	},
	{
		instruction: "Prepare your Springform. Butter the interior thoroughly. Set aside.",
		warning:     "",
	},
	{
		instruction: "Grate the Zitronenschale from an unwaxed lemon. Set aside.",
		warning:     "use an unwaxed lemon. the zest is functional, not decorative.",
	},
	{
		instruction: "Combine the echter Käse, egg yolks, Zucker, Vanille, Zitronenschale, and Speisestärke. Mix until just combined.",
		warning:     "do not overmix. overmixing changes the texture.",
	},
	{
		instruction: "If separating egg whites: beat to soft peaks and fold into the mixture gently.",
		warning:     "",
	},
	{
		instruction: "Pour the filling into the Springform. Smooth the surface.",
		warning:     "",
	},
	{
		instruction: "Place in oven at 165°C fan-assisted. Do not open the oven door during baking.",
		warning:     "opening the oven door causes the filling to sink. do not open it.",
	},
	{
		instruction: "Bake until the center has a slight movement when the form is gently shaken, similar to a set custard, not a liquid.",
		warning:     "",
	},
	{
		instruction: "Turn off the oven. Open the door slightly. Leave the Kaesekuchen inside for 30 to 40 minutes.",
		warning:     "removing it immediately causes surface cracking.",
	},
	{
		instruction: "Remove from oven. Cool completely at room temperature.",
		warning:     "",
	},
	{
		instruction: "Refrigerate for at least 4 hours. Overnight is better.",
		warning:     "the flavor develops during this time. patience is a required parameter.",
	},
	{
		instruction: "Serve cold. No toppings. No fruit. Nothing on top.",
		warning:     "",
	},
}

func cmdMake() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Kaesekuchen. Step by step.")
	fmt.Println("Press enter when each step is done.")
	fmt.Println()

	for i, s := range recipe {
		fmt.Printf("[%d/%d] %s\n", i+1, len(recipe), s.instruction)
		if s.warning != "" {
			fmt.Printf("      note: %s\n", s.warning)
		}
		fmt.Print("      [enter to continue] ")
		scanner.Scan()

		line := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if strings.Contains(line, "fruit") || strings.Contains(line, "strawberr") || strings.Contains(line, "quark") {
			fmt.Println()
			fmt.Println("kk: violation detected in your input.")
			fmt.Println("    this session is over.")
			os.Exit(1)
		}

		fmt.Println()
	}

	fmt.Println("if you followed the steps, you have a Kaesekuchen.")
	fmt.Println("if you added fruit at any point, you do not.")
}
