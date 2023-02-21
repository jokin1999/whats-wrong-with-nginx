package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/jokin1999/whats-wrong-with-nginx/tank"
)

func main() {
	// root user
	err := syscall.Setuid(0)
	if err != nil {
		log.Fatal("Failed to run as root")
	}

	tank.PROC = "nginx"

	// service or systemd
	com := "service"
	service_r := exec.Command("service", "-h")
	service_ret, err := service_r.CombinedOutput()

	if !strings.HasPrefix(string(service_ret), "Usage") || err != nil {
		systemd_r := exec.Command("systemd")
		systemd_ret, err := systemd_r.CombinedOutput()
		if !strings.HasPrefix(string(systemd_ret), "Trying") || err != nil {
			log.Fatal("Failed")
		}
		com = "systemd"
	}
	rand.Seed(time.Now().UnixNano())
	start := func() {
		tr := rand.Intn(int(tank.START_MAX-tank.START_MIN)) + int(tank.START_MIN)
		fmt.Println(tr)
		time.Sleep(time.Duration(tr) * time.Second)
	}
	stop := func() {
		tr := rand.Intn(int(tank.STOP_MAX-tank.STOP_MIN)) + int(tank.STOP_MIN)
		fmt.Println(tr)
		time.Sleep(time.Duration(tr) * time.Second)
	}
	for {
		run(com, "stop")
		start()
		run(com, "start")
		stop()
	}
}

func run(com string, act string) {
	cmd := []string{}
	if com == "systemd" {
		cmd = []string{
			act,
			tank.PROC,
		}
		c := exec.Command(com, cmd...)
		c.Run()
	} else {
		cmd = []string{
			tank.PROC,
			act,
		}
		c := exec.Command(com, cmd...)
		c.Run()
	}
	fmt.Println(com, cmd)
}
