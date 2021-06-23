package ips

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Path(context *cli.Context) error {
	p, err := configPath()
	if err != nil {
		return err
	}
	fmt.Println(p)
	return nil
}
