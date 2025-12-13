package main

import "strings"

func parseList(input string) []string {
	parts := strings.Split(input, ";")
	var result []string
	for _, part := range parts {
		v := strings.TrimSpace(part)
		if v != "" {
			result = append(result, v)
		}

	}
	return result
}
