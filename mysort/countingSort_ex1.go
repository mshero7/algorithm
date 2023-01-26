package mysort

import "fmt"

func main() {
	str := "afjsdjflasjdflkasf"

	var count [26]int
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}
	fmt.Println(count)
	maxCount := 0
	var maxCh byte

	for i := 0; i < len(count); i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxCh = byte('a' + i)
		}
	}

	fmt.Printf("max Character : %c count : %d", maxCh, maxCount)
}
