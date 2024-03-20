package main

import (
	"strings"
)

func main() {
	Manipulate([]string{"a", "a", "b", "b"})
}

func Manipulate(text []string) []string {
	// TODO : start your code here
	var (
		results = map[string]struct{}{}
		length  = len(text)
	)

	if length == 1 {
		return text
	}

	for i := 0; i < length; i++ {

		if i > 0 {
			text = moveElementSliceToPrepend(text, i)
		}

		var temp = make([]string, length)
		copy(temp, text)

		for j := 0; j < length-1; j++ {
			if j == 0 {
				results[strings.Join(text, "")] = struct{}{}
				continue
			}

			temp[j], temp[j+1] = temp[j+1], temp[j]
			results[strings.Join(temp, "")] = struct{}{}
		}
	}

	return mapToSlice(results)
}

func moveElementSliceToPrepend(slice []string, elementPosition int) []string {
	newText := []string{slice[elementPosition]}
	newText = append(newText, slice[:elementPosition]...)
	newText = append(newText, slice[elementPosition+1:]...)
	return newText
}

func mapToSlice(maps map[string]struct{}) []string {
	var results []string
	for k := range maps {
		results = append(results, k)
	}
	return results
}
