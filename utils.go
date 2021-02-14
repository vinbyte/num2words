package num2words

func splitByThousand(number int64) []int64 {
	triplets := []int64{}

	for number > 0 {
		triplets = append([]int64{number % 1000}, triplets...)
		number = number / 1000
	}

	return triplets
}
