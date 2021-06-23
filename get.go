package ips

import (
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"os"
)

func Get(ctx *cli.Context) error {
	config, err := getConfig()
	if err !=  nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	if !ctx.Bool("noheader") {
		table.SetHeader([]string{"Name", "LoginUser", "MacAddress", "IpAddress"})
	}
	users, err := scan()
	if err != nil {
		return err
	}
	for _, user := range users {
		u := config.FindUser(user.MacAddress, user.IpAddress)
		if ctx.Args().Len() != 0 && Find(ctx.Args().Slice(), u.Name) == -1 {
			continue
		}
		table.Append(u.ToSlice())
	}
	table.Render()
	return nil
}
