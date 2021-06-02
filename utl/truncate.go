package utl

func Truncate(s string, l int) string {
	if l != 0 {
		var numRunes = 0
		for index, _ := range s {
			numRunes++
			if numRunes > l {
				return s[:index]
			}
		}
	}
	return s
}
