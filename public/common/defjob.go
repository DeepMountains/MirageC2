package common

import (
	"context"
	"sync"
	"time"
)

type Job struct {
	ID          int
	Sleep       int
	Health      bool // 判断Sessions的健康状况
	Status      bool // 判断Jobs的健康状况
	IPort       string
	Key         string
	Platform    string
	Processname string
	Processid   string
	User        string
	LastRequest time.Time
	Timer       *time.Timer
	ClientIP    string
	Mu          sync.Mutex
	Tasks       map[int]string // 命令执行任务池
	Ctx         context.Context
	Cancel      context.CancelFunc
}

var Jobs []Job

var JobID = 0
