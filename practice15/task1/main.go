package main

func CalculateEvenAndOddNumbers(a []int) (int, int) { // возвращаем (чет, нечет)
	countNumberOfEven := int(0)
	countNumberOfOdd := int(0)

	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			countNumberOfEven++
		}

		if a[i]%2 != 0 {
			countNumberOfOdd++
		}
	}

	return countNumberOfEven, countNumberOfOdd
}

func main() {

}
