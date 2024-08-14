package iteration

// Repeat takes a str string and concat itself 'times' times
func Repeat(str string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += str
	}
	return repeated
}
