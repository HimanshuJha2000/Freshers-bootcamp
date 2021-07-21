package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Matrix struct {
	Rows, Columns int
	Elements [][]int
}

func getNumberOfRows(m Matrix) int {
	return m.Rows
}

func getNumberOfColumns(m Matrix) int {
	return m.Columns
}

func setElement(m *Matrix, i, j, value int) {
	m.Elements[i][j] = value
}

func addMatrices (m1, m2 Matrix) (Matrix, error) {
	if (m1.Rows != m2.Rows) && (m2.Columns != m1.Columns) {
		return Matrix{Rows: 0, Columns: 0, Elements: nil}, errors.New("Addition can't be performed, different number of rows and columns")
	}

	r := getNumberOfRows(m1)
	c := getNumberOfColumns(m1)

	m := make([][]int, r)
	for i := range m {
		m[i] = make([]int, c)
	}
	for i:=0; i<r; i++ {
		for j:=0; j<c; j++ {
			m[i][j] = m1.Elements[i][j] + m2.Elements[i][j];
		}
	}
	return Matrix{Rows: r, Columns: c, Elements: m}, nil
}

func printAsJSON (matrix Matrix) {
	data, err := json.Marshal(matrix)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

func main() {
	matrix1 := [][]int {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix2 := [][]int {{3, 5, 7}, {4, 2, 1}, {3, 6, 5}}

	m1 := Matrix{Rows: 3, Columns: 3, Elements: matrix1}
	m2 := Matrix{Rows: 3, Columns: 3, Elements: matrix2}

	fmt.Println(getNumberOfRows(m1))
	fmt.Println(getNumberOfColumns(m2))

	mpointer := &m1
	setElement(mpointer, 1, 1, 22)

	if result, err := addMatrices(m1, m2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	printAsJSON(m1)
}