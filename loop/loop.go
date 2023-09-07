package loop

func LoopA(c string) string {
	var final string

	for i := 0; i < 6; i++ {
		final = final + c
	}

	return final
}
