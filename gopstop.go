package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	processList, err := process.Processes()
	if err != nil {
		fmt.Println("process.Processes() failed")
		return
	}

	n := rand.Int() % len(processList)
	target := processList[n]
	err = target.Kill()
	if err != nil {
		fmt.Println("process.Processes() Failed, you are lucky!")
	}
}
