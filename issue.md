```
Compare the Keypress rows in the merged.tsv and curated_postprocessed.tsv files. Observe the fact that some keypresses are present in the merged.tsv (i.e., the input to the ELAN curation process), but missing from curated_postprocessed.tsv (i.e., the output of the curation). This issue includes the following steps:

1. Write a Python script to analyze the percentage of such missing keypresses. Itereate over all the session folders that contain both TSV files as mentioned above. Generate a report. Prove or disprove the hypothesis that the keypresses are missing due to close spacing in time (e.g., <1 ms), which causes ELAN to handle the timestamps incorrectly.

2. Write a Python script to copy over the missing keypresses from merged.tsv into the corresponding curated_postprocessed.tsv, preserving the chronological ordering of the event rows.

3. Write a Python script or a new button in data manager to re-upload the fixed curated_postprocessed.tsv files to GCS.

4. changing stuff to force a push
```

## Report and Fix
---

I will write out a data struct for keypresses. which will contain the data values.

```go
type Keypress struct {
    TBegin float64
    TEnd   float64
    Content string
}
```

and a struct for all qualified session folders in the working directory.

```go
type SessionInfo struct {
	SessionName              string
	MergedTsvPath            string
	CuratedPostProcessedPath string
}
```
CLI invocation can be ran without arguments which **only** generate a report of missing keypresses of all session folders inside of the working directory, and posit the reports into a timestamped folder `slipKey_reports_MMDDYYHHMMSS`. If you put in the CLI argument `-f` (example: `slipKey -f`), you will also generate `postprocessed-fixed.tsv, in each session folder. Reports will also reflect this.

the function `ProcessSessions()` will iterate through every folder in the working directory with the prefix of `session-` that also contains a `merged.tsv` and a `curated-postprocessed.tsv`. it will return a struct array of `sessionInfos` of type `SessionInfo{}`

`main()` will then proceed to iterate through `sessionInfos` and perform the following actions. 

1. run the function `BuildMaster()`, which will open the `merged.tsv` file in the seelected folder and store & return values of tier `Keypress` into an array of type `Item{}`.

<!-- 2. run the function `ComparePostProcessed()` will take the returned struct-array and compare its values line-by-line with the stored values of tier `Keypress` inside of `curated_postprocessed.tsv`. if a value is missing, the value is added in the struct Array `KeypressMissing[]` of type `Item{}`. the return will be `KeypressMissing[]`. additionally all values from `curated-postprocessed.tsv` with any missing values from `Keypresses` will be generated in a new file: `postprocessed-fixed.tsv`. -->

2. run the function `ComparePostProcessed()` which will take the post keypress array and the merged keypress array and iterate through them with the goal of collecting:
   - the keypresses that were missed
   - the **Interrupt size** which is the delta(Δ) from when the last successful keypress went through, and the first keypress in a cluster failed.
   - the **Gap size** which describes the number of keypresses that failed to succeed.
   - and the **Gap length** which describes the delta(Δ) of time from when the Interrupt starts and the next successful keypress begins.
 - All of which will be used to generate the metrics used in the exported report. We can use the following struct to reflect this information.
```go
type KeypressGap struct {
	// Start             *Item
	// End               *Item
	MissingKeypresses []Item
	TriggerTime       float64
	Duration          float64
}

func (g *KeypressGap) Size() int {
	return len(g.MissingKeypresses)
}
```
 one KeypressGap would contain a plural of missing Keypresses and from there, using the TriggerTime + the TEnd of the last missing

3. Once `ComparePostProcessed()` is completed, next to run `GenerateSessionReport()` which will export a report `keypress_report.md` containing all missing values for that session, the delta between the missing keypress and the last existing keypress and the percentage of missing keypresses.

this will repeat until `sessionInfos` is exhausted.



