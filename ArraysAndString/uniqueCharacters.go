package ArraysAndString

/**
 *  Implement an algorithm to determine if a string has all unique characters.
 *  What if you can not use additional data structures?
 */

// func main() {
// 	str := "stringalpp"

// 	fmt.Println(isUnique(str))
// }

func isUnique(str string) bool {
	runes := []rune(str)
	hashMap := make(map[string]int)

	for _, char := range runes {
		if value, ok := hashMap[string(char)]; !ok {
			hashMap[string(char)]++
		} else {
			if value == 1 {
				return false
			}
		}
	}
	return true
}
