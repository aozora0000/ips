package ips

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

// +gen * slice:"Map"
type User struct {
	Name       string `yaml:"name"`
	LoginUser  string `yaml:"login_user"`
	MacAddress string `yaml:"mac_address"`
	IpAddress  string `yaml:"ip_address,omitempty"`
}

func (u User) ToSlice() []string {
	return []string{
		u.Name,
		u.LoginUser,
		u.MacAddress,
		u.IpAddress,
	}
}

type Config struct {
	Users []User `yaml:"users"`
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return home, err
	}
	return path.Join(home, ".ips.yaml"), nil
}

func getConfig() (Config, error) {
	var config Config
	p, err := configPath()
	if err != nil {
		return config, err
	}
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		return config, err
	}
	if yaml.Unmarshal(buf, &config) != nil {
		return config, err
	}
	return config, nil
}

func (c Config) FindUser(mac_address string, ip_address string) User {
	for _, user := range c.Users {
		if user.MacAddress == mac_address {
			user.IpAddress = ip_address
			if user.LoginUser == "" {
				user.LoginUser = user.Name
			}
			return user
		}
	}
	return User{Name: "Unknown", LoginUser: "Unknown", MacAddress: mac_address, IpAddress: ip_address}
}
