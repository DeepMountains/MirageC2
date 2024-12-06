package jobs

import (
	"MirageC2_V1.0/public/common"
	"fmt"
)

func ShowSessions() {
	for _, session := range common.Jobs {
		if session.Health == true {
			fmt.Printf("[*] Job ID: \"%d\", Client IP: \"%s\", Shell User: \"%s\", Platform: \"%s\"\n", session.ID, session.ClientIP, session.User, session.Platform)
		}
	}
}

// 当Sessions中，有任务被持续添加时，TaskPool中的任务就会增加。
var taskListID = 0

func TaskPoolAdd(intID int, command string) {
	job := &common.Jobs[intID-1]
	job.Tasks[taskListID] = command
	taskListID = taskListID + 1
}

func ShowTasksPool(intID int) {
	tasks := common.Jobs[intID-1].Tasks
	if len(tasks) == 0 {
		return
	}
	for id, task := range tasks {
		fmt.Printf("[*] Task ID: %d, Task Command: %s\n", id, task)
	}
}
