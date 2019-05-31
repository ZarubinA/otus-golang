package tools

// CheckAndPanicError ...
func CheckAndPanicError(e error) {
	if e != nil {
		panic(e)
	}
}
