package decrypt

func Decrypted(str string) string {
	final := ""
	for _, i := range str {
		ascii := int(i)
		character := string(ascii - 3)
		final += character
	}
	return final

}
