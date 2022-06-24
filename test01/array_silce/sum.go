package array_silce

//数组计算总和
func Sum(numbers []int) (sum int) {
	sum = 0

	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {

		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) (sum []int) {

	for _, number := range numbers {
		sum = append(sum, Sum(number))
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) (sum []int) {

	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {

			sum = append(sum, 0)
		} else {
			tail := numbers[1:]

			sum = append(sum, Sum(tail))
		}
	}

	return sum
}
