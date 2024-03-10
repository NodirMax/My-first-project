package shifrsezary

func Shifrovka(s string) string {
	b := []rune(s)
	f := make([]int32, 0)
	for i := 0; i < len(b); i++ {
		f = append(f, (b[i] + 3))
	}
	return string(f)
}

func RasShifrovka(s string) string {
	b := []rune(s)
	f := make([]int32, 0)
	for i := 0; i < len(b); i++ {
		f = append(f, (b[i] - 3))
	}
	return string(f)
}
