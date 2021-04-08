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
	if len(os.Args) < 3 {
		fmt.Println("./su-starter <target username> <command> ...")
		os.Exit(1)
	}
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
	targetUser, err := user.Lookup(os.Args[1])
	if err != nil {
		log.Fatalf("failed to find target user %s: %s", os.Args[1], err.Error())
	}
	targetUid, err := strconv.Atoi(targetUser.Uid)
	if err != nil {
		log.Fatalf("failed to parse target user id %s: %s", targetUser.Uid, err.Error())
	}
	fmt.Printf("The uid for target user %s is %d\n", os.Args[1], targetUid)
	if err := syscall.Setuid(targetUid); err != nil {
		log.Fatalf("failed to setuid: %s", err.Error())
	}
	execArgs := append([]string{"-c"}, os.Args[2:]...)
	fmt.Printf("going to execute: %v\n", execArgs)
	if err := syscall.Exec("/bin/bash", execArgs, make([]string, 0)); err != nil {
		log.Fatalf("failed to exec: %s", err.Error())
	}
}
