package main

func GradeStudent(number int) int {
	if number < 40 {
		return number
	}

	multipedNumber := number

	for multipedNumber%5 != 0 {

		if multipedNumber%5 != 0 {
			multipedNumber += 1
		}
	}

	differenceBetweenNumberAndMultiped := multipedNumber - number

	if differenceBetweenNumberAndMultiped >= 3 {
		return number
	}

	return multipedNumber
}
