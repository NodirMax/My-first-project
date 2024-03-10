package pluschange

import "strings"

func Change(stroka string) string {
	z := strings.Replace(stroka, "+", " ", -1)
	return z
}
