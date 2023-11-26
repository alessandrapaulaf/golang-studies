package loops

func Repeat(character string, number int) (value string) {
	if number == 0 {
		number = 5
	}

	for i := 0; i < number; i++ {
		value += character
	}
	return
}
