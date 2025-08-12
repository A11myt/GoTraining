package main

import (
	"fmt"
	"os"
)

func transposiert(matrix [5][5]int) [5][5]int {
	var result [5][5]int
	for i, row := range matrix {
		for j, val := range row {
			result[j][i] = val
		}
	}
	return result
}

func schreibeMatrixInDatei(dateiname string, m [5][5]int) error {
	file, err := os.OpenFile(dateiname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, row := range m {
		for j, val := range row {
			if j > 0 {
				fmt.Fprint(file, " ")
			}
			fmt.Fprint(file, val)
		}
		fmt.Fprintln(file)
	}
	return nil
}

func main() {
	for nr := 1; nr <= 4; nr++ {
		matrix := gibMatrix(nr)
		matrix = transposiert(matrix)
		fmt.Println("Transponierte Matrix:")
		for _, row := range matrix {
			fmt.Println(row)
		}
		schreibeMatrixInDatei("matrix.txt", matrix)
	}
}
