package ips

import "os"

func Find(slice []string, str string) int {
	for index, data := range slice {
		if data == str {
			return index
		}
	}
	return -1
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}