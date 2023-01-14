package main

import "fmt"

func CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3 int) bool {
	if denominationCoin1 > 0 && costOfProduct%denominationCoin1 == 0 {
		return true
	}

	if denominationCoin2 > 0 && costOfProduct%denominationCoin2 == 0 {
		return true
	}

	if denominationCoin3 > 0 && costOfProduct%denominationCoin3 == 0 {
		return true
	}

	if denominationCoin1 > 0 && denominationCoin2 > 0 {
		remain := costOfProduct % denominationCoin1
		if remain%denominationCoin2 == 0 {
			return true
		}
	}

	if denominationCoin1 > 0 && denominationCoin2 > 0 {
		remain := costOfProduct % denominationCoin2
		if remain%denominationCoin1 == 0 {
			return true
		}
	}

	if denominationCoin2 > 0 && denominationCoin3 > 0 {
		remain := costOfProduct % denominationCoin2
		if remain%denominationCoin3 == 0 {
			return true
		}
	}

	if denominationCoin2 > 0 && denominationCoin3 > 0 {
		remain := costOfProduct % denominationCoin3
		if remain%denominationCoin2 == 0 {
			return true
		}
	}

	if denominationCoin1 > 0 && denominationCoin2 > 0 && denominationCoin3 > 0 {
		remain := costOfProduct % denominationCoin1
		remain2 := remain % denominationCoin2
		if remain2%denominationCoin3 == 0 {
			return true
		}
	}

	if denominationCoin1 > 0 && denominationCoin2 > 0 && denominationCoin3 > 0 {
		remain := costOfProduct % denominationCoin1
		remain2 := remain % denominationCoin3
		if remain2%denominationCoin2 == 0 {
			return true
		}
	}

	if denominationCoin1 > 0 && denominationCoin2 > 0 && denominationCoin3 > 0 {
		remain := costOfProduct % denominationCoin2
		remain2 := remain % denominationCoin1
		if remain2%denominationCoin3 == 0 {
			return true
		}
	}

	if denominationCoin1 > 0 && denominationCoin2 > 0 && denominationCoin3 > 0 {
		remain := costOfProduct % denominationCoin2
		remain2 := remain % denominationCoin3
		if remain2%denominationCoin1 == 0 {
			return true
		}
	}

	if denominationCoin1 > 0 && denominationCoin2 > 0 && denominationCoin3 > 0 {
		remain := costOfProduct % denominationCoin3
		remain2 := remain % denominationCoin1
		if remain2%denominationCoin2 == 0 {
			return true
		}
	}

	if denominationCoin1 > 0 && denominationCoin2 > 0 && denominationCoin3 > 0 {
		remain := costOfProduct % denominationCoin3
		remain2 := remain % denominationCoin2
		if remain2%denominationCoin1 == 0 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Введите стоимость товара:")
	var costOfProduct int
	fmt.Scan(&costOfProduct)

	fmt.Println("Введите номинал первой монеты:")
	var denominationCoin1 int
	fmt.Scan(&denominationCoin1)

	fmt.Println("Введите номинал второй монеты:")
	var denominationCoin2 int
	fmt.Scan(&denominationCoin2)

	fmt.Println("Введите номинал третьей монеты:")
	var denominationCoin3 int
	fmt.Scan(&denominationCoin3)

	if CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3) {
		fmt.Println("Можно заплатить без сдачи")

		return
	}

	fmt.Println("Нельзя заплатить без сдачи")
}
