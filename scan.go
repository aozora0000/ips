package ips

import (
	"bufio"
	"os/exec"
	"strings"
)

func scan() ([]User, error) {
	var config Config
	var users []User
	cmd := exec.Command("sh", "-c", "sudo arp -en | sed '1d' | awk '{print $3,$1}'")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return users, err
	}
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for i := 0; scanner.Scan(); i++ {
		item := strings.Split(scanner.Text(), " ")
		users = append(users, config.FindUser(item[0], item[1]))
	}
	return users, cmd.Wait()
}
