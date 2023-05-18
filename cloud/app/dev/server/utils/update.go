package utils

import "os/exec"

func UpdateHermes() {
	var cmd *exec.Cmd = exec.Command("go", "get", "-u", "github.com/realTristan/Hermes")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
