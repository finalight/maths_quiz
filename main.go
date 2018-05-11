package main

import "fmt"
import "encoding/csv"
import "os"
import "bufio"

// func fn(m *map[int]int) {
// 	m = new(map[int]int)
// }

var totalQns, correct = 0, 0

func main() {
	readCSV("problems")
	fmt.Println("Number of corrects answer", correct)
}

func readCSV(filename string) {
	file, err := os.Open(filename + ".csv")

	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	data, err := r.ReadAll()
	totalQns = len(data)
	for value := range data {
		fmt.Print(data[value][0], " = ")
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			if scanner.Text() == data[value][1] {
				fmt.Println("correct!!")
				correct++
			}
			break
		}
		if scanner.Err() != nil {
			// handle error.
		}

	}

}
