package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows     int
	Cols     int
	Elements [][]int
}

func CreateMatrix(r, c int) Matrix {
	// everything 0
	elements := make([][]int, r)
	for i := range elements {
		elements[i] = make([]int, c)
	}
	return Matrix{
		Rows:     r,
		Cols:     c,
		Elements: elements,
	}
}

func (m *Matrix) GetRows() int {
	return m.Rows
}
func (m *Matrix) GetCols() int {
	return m.Cols
}
func (m *Matrix) SetField(i, j, val int) {
	m.Elements[i][j] = val
}

func (m Matrix) AddMatrix(other Matrix) Matrix {
	result := CreateMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}

	return result
}

func (m Matrix) MatrixToJSON() (string, error) {
	// Marshal function to convert here // unmarshal to do the reverse
	json, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func main() {
	matrix1 := CreateMatrix(3, 4)
	matrix2 := CreateMatrix(3, 4)
	matrix1.SetField(0, 0, 3)
	matrix1.SetField(1, 1, 6)
	matrix1.SetField(2, 2, 9)

	matrix2.SetField(0, 0, 3)
	matrix2.SetField(0, 1, 2)
	matrix2.SetField(0, 2, 5)
	matrix2.SetField(1, 0, 2)
	matrix2.SetField(1, 1, 3)
	matrix2.SetField(1, 2, 5)
	matrix2.SetField(2, 0, 9)
	matrix2.SetField(2, 1, 3)
	matrix2.SetField(2, 2, 8)
	//get rows
	fmt.Printf("matrix1 has %d rows \n", matrix1.GetRows())
	//get cols
	fmt.Printf("matrix1 has %d cols \n", matrix1.GetCols())
	//set field implemented above
	//add matrix
	ans := matrix1.AddMatrix(matrix2)
	fmt.Println(ans.Elements)
	//matrix2json
	json, err := matrix1.MatrixToJSON()
	if err =git = nil {
		fmt.Println("json is ", json)
	} else {
		fmt.Println("error is", err)
	}
}
