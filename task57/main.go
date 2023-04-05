package main

func GradeStudents(array []int) []int {
	graded := []int{}

	for indexOfArray := 0; indexOfArray < len(array); indexOfArray++ {
		numberInArray := array[indexOfArray]

		if numberInArray < 38 {
			graded = append(graded, numberInArray)

			continue
		}

		for numberInArray%5 != 0 {
			if numberInArray%5 != 0 {
				numberInArray += 1
			}
		}

		differenceBetweenumbers := numberInArray - array[indexOfArray]

		if differenceBetweenumbers >= 3 {
			graded = append(graded, array[indexOfArray])
		} else if differenceBetweenumbers < 3 {
			graded = append(graded, numberInArray)
		}

		if numberInArray < 40 && differenceBetweenumbers < 3 {
			graded = append(graded, numberInArray)
		}

	}

	return graded
}
