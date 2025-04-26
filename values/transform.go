package values

type TransCountOption struct {
	Symbol    string
	Breakline int
}

func TransCount(count int, symbol string, breakline int) string {
	var result string = ""

	for i := count; i > 0; i-- {
		if breakline > 0 && i%breakline == 0 {
			result += "\n"
		}

		result += symbol
	}

	return result
}
