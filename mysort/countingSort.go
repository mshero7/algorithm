package mysort

import "fmt"

func main() {
	arr := []int{2, 6, 1, 2, 4, 5, 2, 3, 5, 6, 1, 2, 6, 8, 9, 7, 6, 5}

	// arr에 할당된 값의 범위 +1 만큼 배열을 할당.
	var count [11]int

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	sorted := make([]int, 0, len(arr))
	// for문을 2번도는 약점
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}

	fmt.Println(sorted)
}
