package app

import "fmt"

func generateSessionReport(gapArr []keypressGap) {
	for _, gap := range gapArr {
		fmt.Printf("GAP BEGIN \n")
		for _, keypress := range gap.MissingKeypresses {
			fmt.Printf("TBegin: %f ,\nTEnd: %f,\nTier: %s ,\nContent: %s \n", keypress.TBegin, keypress.TEnd, keypress.Tier, keypress.Content)
			fmt.Printf("\n")
		}
		fmt.Printf("Trigger Time: %f,\nDuration: %f \n", gap.TriggerTime, gap.Duration)
		fmt.Printf("GAP END \n")
	}
}
