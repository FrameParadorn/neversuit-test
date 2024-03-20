package main

func CountSmilyFace(text []string) int {
	// TODO : start your code hero
	validSmily := map[string]struct{}{
		":)":  {},
		";)":  {},
		":D":  {},
		";D":  {},
		":-)": {},
		";-)": {},
		":~)": {},
		";~)": {},
		":-D": {},
		";-D": {},
		":~D": {},
		";~D": {},
	}

	count := 0
	for _, i := range text {
		if _, ok := validSmily[i]; ok {
			count++
		}
	}

	return count
}
