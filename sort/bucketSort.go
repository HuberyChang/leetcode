package main

func bucketSort(arr []float32) []float32 {
	max := arr[0]
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	d := max - min

	bucketNum := len(arr)
	bucketList := make([]float32, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bu
	}

	for i := 0; i < len(arr); i++ {
		num := int((arr[i] - min) * (bucketNum - 1) / d)
		append(bucketList, arr[i])
	}

	for i := 0; i < len(bucketList); i++ {

	}

	return arr 
}

func main() {

}
