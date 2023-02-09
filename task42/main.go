package main

import "fmt"

func main() {
	langByPerson := map[string]string{
		"ondrys":    "go",
		"yurkevich": "php",
		"kry":       "go",
		"roma":      "c#",
	}

	people := []string{"ondrys", "yurkevich", "kry", "roma"}
	langs := []string{"go", "php", "go", "c#"}

	key := 2
	valuePeople := people[key]
	valueLang := langs[key]

	fmt.Println(key, valuePeople, valueLang)

	person := "kry"
	lang := langByPerson[person]

	fmt.Println(lang)
}
