package matrix

import (
    "strings"
    "strconv"
    "fmt"
)
// Define the Matrix and Pair types here.
type Matrix struct {
    data [][]int
    rows int
    cols int
}

type Pair struct {
    Row int
    Col int
}

func New(s string) (*Matrix, error) {
    s = strings.TrimSpace(s)
    if s == "" {
        return &Matrix{[][]int{}, 0, 0}, nil
    }

    lines := strings.Split(s, "\n")
    data := [][]int{}
    
    for _, line := range lines {
        fields := strings.Fields(line)
        if len(fields) == 0 {
            continue
        }
        row := []int{}
        for _, f := range fields {
            n, err := strconv.Atoi(f)
            if err != nil {
                return nil, err
            }
            row = append(row, n)
        }
        data = append(data, row)
    }

    rows := len(data)
    cols := len(data[0])
    for _, row := range data {
        if len(row) != cols {
            return nil, fmt.Errorf("inconsistent row lengths")
        }
    }

    return &Matrix{data, rows, cols}, nil
}

func (m *Matrix) Saddle() []Pair {
    if m.rows == 0 || m.cols == 0 {
        return []Pair{}
    }

    var result []Pair

    // Step 1: Find max in each row
    for i := 0; i < m.rows; i++ {
        row := m.data[i]
        maxVal := row[0]
        maxIndices := []int{0}

        for j := 1; j < m.cols; j++ {
            if row[j] > maxVal {
                maxVal = row[j]
                maxIndices = []int{j}
            } else if row[j] == maxVal {
                maxIndices = append(maxIndices, j)
            }
        }

        // Step 2: For each max in row, check if it's min in its column
        for _, j := range maxIndices {
            isMin := true
            for k := 0; k < m.rows; k++ {
                if m.data[k][j] < maxVal {
                    isMin = false
                    break
                }
            }
            if isMin {
                result = append(result, Pair{Row: i + 1, Col: j + 1})
            }
        }
    }

    return result
}
