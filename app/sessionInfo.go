package app

//SessionInfo is the struct of every "session-"" folder filtered by processSessions().
type SessionInfo struct {
	SessionName              string
	SessionPath              string
	MergedTsvPath            string
	CuratedPostProcessedPath string
}
