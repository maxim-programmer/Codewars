// Write a function that accepts a non-negative integer n and a string s as parameters, and returns a string of s repeated exactly n times.
package kata

func RepeatStr(repetitions int, value string) string {
	str := ""
	for i := 0; i < repetitions; i++ {
		str += value
	}
	return str
}
