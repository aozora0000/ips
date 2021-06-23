package ips

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type User struct {
	Name       string `yaml:"name"`
	LoginUser  string `yaml:"login_user"`
	MacAddress string `yaml:"mac_address"`
	IpAddress  string `yaml:"ip_address"`
}

func (u User) ToSlice() []string {
	return []string{
		u.Name,
		u.LoginUser,
		u.MacAddress,
		u.IpAddress,
	}
}

func (e *User) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {

	var aux interface{}
	if err = unmarshal(&aux); err != nil {
		return
	}

	switch raw := aux.(type) {
	case string:
		*e = []string{raw}

	case []interface{}:
		list := make([]string, len(raw))
		for i, r := range raw {
			v, ok := r.(string)
			if !ok {
				return fmt.Errorf("An item in evn cannot be converted to a string: %v", aux)
			}
			list[i] = v
		}
		*e = list

	}
	return
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
	return User{Name: "Unknown", LoginUser: "Unknown" ,MacAddress: mac_address, IpAddress: ip_address}
}
