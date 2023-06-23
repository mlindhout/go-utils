package mystrings

func Reverse(msg string) (tmp string) {
	tmp = ""
	for i := len(msg) - 1; i >= 0; i-- {
		tmp += msg[i : i+1]
	}
	return
}
