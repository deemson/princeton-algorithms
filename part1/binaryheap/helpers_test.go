package binaryheap_test

func sliceStringIntoCharacters(str string) []string {
	slice := make([]string, len(str))
	for index, char := range str {
		slice[index] = string(char)
	}
	return slice
}

func stringAsIs(str string) string { return str }
