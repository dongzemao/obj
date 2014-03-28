package common

func Substr(str string, star int, num int) string {
	rs := []rune(str)
	l := len(str)
	if star < 0 {
		star = 0
	}
	end := star + num
	if end > l {
		end = l
	}

	return string(rs[star:end])
}
