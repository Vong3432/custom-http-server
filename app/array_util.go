package main

// Generic function to get last element of a slice
func GetLast[T any](s []T) *T {
	if len(s) == 0 {
		return nil
	}
	return &s[len(s)-1]
}
