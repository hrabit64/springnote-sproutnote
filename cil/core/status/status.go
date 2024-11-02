package status

var runningStatus = NONE

func SetRunningStatus(status string) {
	runningStatus = status
}

func GetRunningStatus() string {
	return runningStatus
}
