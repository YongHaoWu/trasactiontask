package main

import (
	"fmt"
	"trasactiontask/task"
)

type RedisCluster struct {
	port string
	msg  string
}

func (a *RedisCluster) handleMsg(action string) bool {
	switch action {
	case "acquirePort":
		return acquirePort(a.port)
	case "releasePort":
		return releasePort()
	case "CreateRedisCluster":
		return CreateRedisCluster()
	case "cleanRedisCluster":
		return cleanRedisCluster()
	}
	return false
}

func acquirePort(port string) bool {
	fmt.Println("acquirePort", port)
	return true
}
func CreateRedisCluster() bool {
	fmt.Println("CreateRedisCluster")
	return false
}
func releasePort() bool {
	fmt.Println("releasePort")
	return true
}
func cleanRedisCluster() bool {
	fmt.Println("cleanRedisCluster")
	return true
}

func main() {
	r := RedisCluster{port: "8542"}
	t := task.Task{ActionCallback: r.handleMsg}
	acquirePortProcess := task.Process{
		Goon:     "acquirePort",
		Rollback: "releasePort",
		Action:   "continue",
	}
	createRedisClusterProcess := task.Process{
		Goon:     "CreateRedisCluster",
		Rollback: "cleanRedisCluster",
		Action:   "continue",
	}
	t.Processes = []task.Process{acquirePortProcess, createRedisClusterProcess}
	t.Exec()
}
