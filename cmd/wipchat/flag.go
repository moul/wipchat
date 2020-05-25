package main

import "fmt"

type stringSlice []string

func (s stringSlice) String() string {
	return fmt.Sprintf("%v", []string(s))
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}
