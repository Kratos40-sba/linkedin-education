package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func killContainer(fileName string) error  {
	data , err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			return
		}
	}(fileName)
	cid := strings.TrimSpace(string(data))
	if !validCID(cid) {
		return fmt.Errorf("bad container id:%v",cid)
	}
	log.Printf("Stopping container %v",cid)
	if err = exec.Command("docker", "rm", "-f", cid).Run(); err != nil {
		return fmt.Errorf("failed to to stop %v : %w",cid,err)
	}
	return nil
}

func validCID( cid string) bool {
	return len(cid) == 12 || len(cid) == 64
}