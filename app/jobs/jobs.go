package jobs

import (
	"MirageC2_V1.0/app/webstart"
	"MirageC2_V1.0/public/common"
	"context"
	"fmt"
	"strconv"
)

func CreateJobs(ip string) {
	common.JobID += 1
	job := common.Job{
		ID:     common.JobID,
		IPort:  ip,
		Health: false,
		Status: true,
		Sleep:  common.MConfig.CheckInterval,
		Tasks:  make(map[int]string),
	}
	job.Ctx, job.Cancel = context.WithCancel(context.Background())
	common.Jobs = append(common.Jobs, job)
	fmt.Println("[+] Jobs Set: listen ip", common.Jobs[common.JobID-1].IPort)
	go webstart.MirageServer(ip, strconv.Itoa(common.JobID), job.Ctx)
}

func ShowJobs() {
	for _, job := range common.Jobs {
		if job.Status == true {
			fmt.Printf("[*] Job ID: \"%d\", IPort: \"%s\", Key: \"%s\", SleepTime: \"%d\"\n", job.ID, job.IPort, job.Key, job.Sleep)
		}
	}
}

func ShowAllJobs() {
	for _, job := range common.Jobs {
		fmt.Printf("[*] Job ID: \"%d\", IPort: \"%s\", Key: \"%s\", Status: \"%t\"\n", job.ID, job.IPort, job.Key, job.Status)
	}
}

/*
/*
func SleepTimeChange(intID, time int) {
	common.Jobs[intID-1].Sleep = time
}
*/
