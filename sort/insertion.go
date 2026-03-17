package sort

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j-1 >= 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
