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

	for _, gap := range gapArr {
		missingKeyTotal += len(gap.MissingKeypresses)
		interruptTotal += float32(gap.TriggerTime)
		durationTotal += float32(gap.Duration)
		if smallestInterruptThreshold > float32(gap.TriggerTime) {
			smallestInterruptThreshold = float32(gap.TriggerTime)
		}

	}
	if boolean {
		tsvGenerated = "YES"
	}
	keypressGapAvg = float32(missingKeyTotal) / float32(len(gapArr))
	averageInterruptDelta = interruptTotal / float32(len(gapArr))
	averageGapDurationDelta = durationTotal / float32(len(gapArr))

	slipKeyHeader := make([]string, 0)
	// slipKeyHeader = append(slipKeyHeader,
	// 	"# SlipKey Report\n",
	// 	fmt.Sprintf("## Session Name - %s\n", sessName),
	// 	fmt.Sprintf("### Entry Date - %d %d %d\n", now.Month(), now.Day(), now.Year()),
	// 	"\n",
	// 	fmt.Sprintf("postprocessed-fixed.tsv generated?: %s\n", tsvGenerated),
	// 	"\n",
	// 	fmt.Sprintf("- Amount of Missing Keypresses: %d\n", missingKeyTotal),
	// 	fmt.Sprintf("- Average Keypress Gap Size: %.2f\n", keypressGapAvg),
	// 	fmt.Sprintf("- Average Interrupt Time Δ: %.3f\n", averageInterruptDelta),
	// 	fmt.Sprintf("- Smallest Interrupt Threshold: %.3f\n", smallestInterruptThreshold),
	// 	fmt.Sprintf("- Average Gap Duration: %.3f\n", averageGapDurationDelta),
	// )
	slipKeyHeader = append(slipKeyHeader,
		"<!DOCTYPE html>\n",
		fmt.Sprintf(`<html lang="en">`+"\n"),
		"<head>\n",
		fmt.Sprintf(`<meta charset="UTF-8">`+"\n"),
		fmt.Sprintf(`<meta http-equiv="X-UA-Compatible" content="IE=edge">`+"\n"),
		fmt.Sprintf(`<meta name="viewport" content="width=device-width, initial-scale=1.0">`+"\n"),
		"<title>slipKey Report</title>\n",
		fmt.Sprintf(`<link rel="stylesheet" href="style.css">`+"\n"),
		"</head>\n",
		"<body>\n",
		fmt.Sprintf(`<img src="TG-horizontal_logo.svg">`+"\n"),
		"<h1>slipKey Report</h1>\n",
		"<hr>\n",
		fmt.Sprintf("<h2>Session Name - %s</h2>\n", sessName),
		fmt.Sprintf("<h3>Date %d %d %d</h3>\n", now.Month(), now.Day(), now.Year()),
		fmt.Sprintf("<p><b>postprocessed-fixed.tsv generated?</b>: %s</p>\n", tsvGenerated),
		"<ul>\n",
		fmt.Sprintf("<li>Amount of Missing Keypresses: %d</li>\n", missingKeyTotal),
		fmt.Sprintf("<li>Average Keypress Gap size: %.2f</li>\n", keypressGapAvg),
		fmt.Sprintf("<li>Average Interrupt Time Δ: %.3f</li>\n", averageInterruptDelta),
		fmt.Sprintf("<li>Smallest Interrupt Threshold: %.3f</li>\n", smallestInterruptThreshold),
		fmt.Sprintf("<li>Average Gap Duration: %.3f</li>\n", averageGapDurationDelta),
		"</ul>\n",
		"<hr>\n",
		"</body>\n",
		"</html>\n",
	)
	report, err := os.Create(fmt.Sprintf("%s/SlipKey-Report-%s.html", folderStr, sessName))
	if err != nil {
		fmt.Printf("failed to create Report -%v", err)
	}
	defer report.Close()

	for _, str := range slipKeyHeader {
		report.WriteString(str)
	}

	for index, gap := range gapArr {
		report.WriteString(fmt.Sprintf("\n<hr>\n<b>Gap Set %d</b>\n", index+1))
		report.WriteString("<ul>\n")
		for _, item := range gap.MissingKeypresses {
			report.WriteString(fmt.Sprintf("<li>%.3f\t%.3f\t%s</li>\n", item.TBegin, item.TEnd, item.Content))
		}
		report.WriteString("</ul>\n")
	}
}
