package arr

var value = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// RomanToInt
// https://redirect.liyou-chen.com/kadapt
func RomanToInt(s string) int {
	result := value[string(s[len(s)-1])]
	for i := len(s) - 2; i >= 0; i-- {
		if value[string(s[i])] < value[string(s[i+1])] {
			result -= value[string(s[i])]
			continue
		}
		result += value[string(s[i])]
	}
	return result
}
