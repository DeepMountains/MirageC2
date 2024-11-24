package check

import "MirageC2_V1.0/public/common"

func JobsExistsCheck(jobID int) bool {
	if jobID < 0 || jobID > len(common.Jobs) {
		return false
	}
	return true
}

func SessionsExistsCheck(jobID int) bool {
	if jobID < 0 || jobID > len(common.Jobs) {
		return false
	}
	if common.Jobs[jobID-1].Health == false {
		return false
	}
	return true
}
