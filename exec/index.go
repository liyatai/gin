package execute

import (
	"fmt"
	"os/exec"
)

func DoCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)

	err := cmd.Run()
	if err != nil {
		fmt.Println("命令执行出错", err)
	}
}
