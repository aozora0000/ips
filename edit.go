package ips

import (
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

func Edit(context *cli.Context) error {
	p, err := configPath()
	if err != nil {
		return err
	}
	cmd := exec.Command(os.Getenv("EDITOR"), p)
					cmd.Stdin = os.Stdin
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					return cmd.Run()
}