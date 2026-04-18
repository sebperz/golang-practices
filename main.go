package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	output := map[string]user{}
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	for i, name := range names {
		output[name] = user{
			phoneNumber: phoneNumbers[i],
			name:        name}
	}
	return output, nil
}

type user struct {
	name        string
	phoneNumber int
}
