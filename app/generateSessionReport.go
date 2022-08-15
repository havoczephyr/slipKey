package app

import (
	"fmt"
	"os"
	"time"
)

func generateSessionReport(gapArr []keypressGap, sessName, sessPath string, boolean bool, folderStr string) {
	// for _, gap := range gapArr {
	// 	fmt.Printf("GAP BEGIN \n")
	// 	for _, keypress := range gap.MissingKeypresses {
	// 		fmt.Printf("TBegin: %f ,\nTEnd: %f,\nTier: %s ,\nContent: %s \n", keypress.TBegin, keypress.TEnd, keypress.Tier, keypress.Content)
	// 		fmt.Printf("\n")
	// 	}
	// 	fmt.Printf("Trigger Time: %f,\nDuration: %f \n", gap.TriggerTime, gap.Duration)
	// 	fmt.Printf("GAP END \n")
	// }
	now := time.Now()
	var missingKeyTotal int
	var interruptTotal float32
	var durationTotal float32
	var tsvGenerated string = "NO"
	var keypressGapAvg float32
	var averageInterruptDelta float32
	var averageGapDurationDelta float32
	var smallestInterruptThreshold float32 = 999.999

	var interrupt interruptData
	interrupt.folderPath = folderStr
	interrupt.sessionName = sessName

	for _, gap := range gapArr {
		missingKeyTotal += len(gap.MissingKeypresses)

		interruptTotal += float32(gap.TriggerTime)
		interrupt.triggerTimes = append(interrupt.triggerTimes, gap.TriggerTime)

		durationTotal += float32(gap.Duration)

		if smallestInterruptThreshold > float32(gap.TriggerTime) {
			smallestInterruptThreshold = float32(gap.TriggerTime)
			interrupt.smallestInterrupt = smallestInterruptThreshold
		}

	}
	if boolean {
		tsvGenerated = "YES"
	}
	keypressGapAvg = float32(missingKeyTotal) / float32(len(gapArr))
	averageInterruptDelta = interruptTotal / float32(len(gapArr))
	averageGapDurationDelta = durationTotal / float32(len(gapArr))

	slipKeyHeader := make([]string, 0)
	slipKeyHeader = append(slipKeyHeader,
		"# SlipKey Report\n",
		fmt.Sprintf("## Session Name - %s\n", sessName),
		fmt.Sprintf("### Entry Date - %d %d %d\n", now.Month(), now.Day(), now.Year()),
		"\n",
		fmt.Sprintf("postprocessed-fixed.tsv generated?: %s\n", tsvGenerated),
		"\n",
		fmt.Sprintf("- Amount of Missing Keypresses: %d\n", missingKeyTotal),
		fmt.Sprintf("- Average Keypress Gap Size: %.2f\n", keypressGapAvg),
		fmt.Sprintf("- Average Interrupt Time Δ: %.3f\n", averageInterruptDelta),
		fmt.Sprintf("- Smallest Interrupt Threshold: %.3f\n", smallestInterruptThreshold),
		fmt.Sprintf("- Average Gap Duration: %.3f\n", averageGapDurationDelta),
	)

	report, err := os.Create(fmt.Sprintf("%s/SlipKey-Report-%s.md", folderStr, sessName))
	if err != nil {
		fmt.Printf("failed to create Report -%v", err)
	}
	defer report.Close()

	for _, str := range slipKeyHeader {
		report.WriteString(str)
	}
	for index, gap := range gapArr {
		report.WriteString(fmt.Sprintf("\n---\n**Gap Set #%d**\n", index+1))
		for _, item := range gap.MissingKeypresses {
			report.WriteString(fmt.Sprintf("- %.3f\t%.3f\t%s\n", item.TBegin, item.TEnd, item.Content))
		}
	}
	freqPath, err := generateReportHist(interrupt)
	if err != nil {
		fmt.Printf("failed to generate frequency graph -%v", err)
	}
	report.WriteString(fmt.Sprintf("\n<img src=%s>", freqPath))
}
