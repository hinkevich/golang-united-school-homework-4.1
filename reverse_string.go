package reverse_string

func ReverseString(input string) (output string) {
	var inputArrayRune = []rune(input)
	var outputArrayRune []rune
	for i := len(inputArrayRune) - 1; i >= 0; i-- {
		outputArrayRune = append(outputArrayRune, inputArrayRune[i])
	}
	return string(outputArrayRune)
}
