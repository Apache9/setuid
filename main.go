package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func main() {
	uid := os.Getuid()
	u, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		log.Fatalf("failed to get uid: %s", err.Error())
	}
	fmt.Printf("user: %s\n", u.Username)

	euid := os.Geteuid()

	eu, err := user.LookupId(strconv.Itoa(euid))
	if err != nil {
		log.Fatalf("failed to get euid: %s", err.Error())
	}
	fmt.Printf("effective user: %s\n", eu.Username)
	if err := syscall.Setuid(1000); err != nil {
		log.Fatalf("failed to setuid: %s", err.Error())
	}

	if err := syscall.Exec("/bin/bash", []string{"test.sh"}, make([]string, 0)); err != nil {
		log.Fatalf("failed to exec: %s", err.Error())
	}
}
