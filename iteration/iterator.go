package iteration

func repeat(character string, repeatCount int) string {
	if repeatCount == 0 {
		return character
	}

	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
