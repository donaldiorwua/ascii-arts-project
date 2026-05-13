package main

func MergeBanner(standard, shadow map[rune][]string) map[rune][]string {
	result := make(map[rune][]string)

	for k, v := range standard {
		if _, exists := result[k]; !exists {
			result[k] = v
		}
	}
	for k, v := range shadow{
		if _, exists := result[k]; !exists{
			result[k] = v
		}
	}
	return result
}
