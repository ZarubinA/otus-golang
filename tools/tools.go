package tools

// CheckAndPanicError check the given error and panic if occurred some one
func CheckAndPanicError(e error) {
	if e != nil {
		panic(e)
	}
}
