package mysort

import "fmt"

func main() {
	arr := []int{2, 6, 1, 2, 4, 5, 2, 3, 5, 6, 1, 2, 6, 8, 9, 7, 6, 5}

	// arr에 할당된 값의 범위 +1 만큼 배열을 할당.
	var count [11]int

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	fmt.Println("count1:", count)

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	fmt.Println("count2:", count)

	sorted := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	fmt.Println(sorted)
}
