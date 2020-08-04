package cron

import (
	"fmt"
	"os/exec"
)

func UploadToGit() {
	cmd := exec.Command("/bin/sh","pwd")
	fmt.Println(cmd.String())
	cmd.Run()
}
