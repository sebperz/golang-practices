package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	valorDeSalida := map[string]user{}
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid len")
	}
	for i, name := range names {
		valorDeSalida[name] = user{
			phoneNumber: phoneNumbers[i],
			name:        name}
	}
	return valorDeSalida, nil
}

type user struct {
	name        string
	phoneNumber int
}
