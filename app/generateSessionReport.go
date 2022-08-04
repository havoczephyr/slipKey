package app

import "fmt"

func generateSessionReport(gapArr []keypressGap) {
	for _, gap := range gapArr {
		for _, keypress := range gap.MissingKeypresses {
			fmt.Printf("TBegin: %f, TEnd: %f, Tier: %s, Content: %s", keypress.TBegin, keypress.TEnd, keypress.Tier, keypress.Content)
		}
		fmt.Printf("Trigger Time: %f, Duration: %f", gap.TriggerTime, gap.Duration)
	}
}
