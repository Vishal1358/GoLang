package main

import "fmt"

var pl = fmt.Println

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"ohio":         11614373,
	}

	statePopulations["Georgia"] = 10310371
	// pl("\n", statePopulations["Texas"])
	pl(statePopulations)
	delete(statePopulations, "Georgia")
	pl(statePopulations)

}
