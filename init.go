package ips

import (
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"os"
)

func Initialize(ctx *cli.Context) error {
	p, err := configPath()
	if err != nil {
		return err
	}
	if !Exists(p) {
		f, err := os.Create(p)
		if err != nil {
			return err
		}
		defer f.Close()
		users, err := scan()
		if err != nil {
			return err
		}
		config := &Config{Users: users}
		data, _ := yaml.Marshal(&config)
        // 正常に生成されたファイルに書き込み
        _, err = f.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}
