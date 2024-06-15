package iteration

const limit = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < limit; i++ {
		repeated += character
	}
	return repeated
}
