package reverse

//Reverse returns the reversed version of the string passed as input
func Reverse(word string) string {
	var revString []rune
	chars := []rune(word)
	for i := 0; i < len(chars); i++ {
		revString = append(revString, chars[len(chars)-i-1])
	}
	return string(revString)
}
