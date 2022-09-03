package utils

import (
	"errors"
	"regexp"
)

func IsValidContent(content, name string) error {
	if content == "" {
		return errors.New(name + " can't be empty")
	}
	return nil
}

type UUID_Detail struct {
	UUID string
	Name string
}

var UUID_Data = []UUID_Detail{}

func IsValidUUID(data []UUID_Detail) error {
	if data[0].UUID == "" {
		return errors.New("invalid UUID " + data[0].Name)
	}
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	for i := 0; i < len(data); i++ {
		status := r.MatchString(data[i].UUID)
		if !status {
			return errors.New("invalid UUID " + data[i].Name)
		}
	}
	return nil
}
