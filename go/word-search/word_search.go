package wordsearch

import "fmt"

type Direction struct {
    dx, dy int
}

func getAllDirections() []Direction {
    return []Direction{
        {1, 0},   // right
        {0, 1},   // down
        {1, 1},   // diagonal down-right
        {1, -1},  // diagonal up-right
        {-1, 0},  // left
        {0, -1},  // up
        {-1, -1}, // diagonal up-left
        {-1, 1},  // diagonal down-left
    }
}

func isValidPosition(x, y int, puzzle []string) bool {
    return y >= 0 && y < len(puzzle) && x >= 0 && x < len(puzzle[0])
}

// findWord searches for a word in the puzzle starting from position (x,y) in given direction
func findWord(word string, puzzle []string, x, y int, dir Direction) bool {
    if !isValidPosition(x, y, puzzle) {
        return false
    }
    
    for i := 0; i < len(word); i++ {
        newX, newY := x+i*dir.dx, y+i*dir.dy
        if !isValidPosition(newX, newY, puzzle) || puzzle[newY][newX] != word[i] {
            return false
        }
    }
    return true
}


func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
    if len(puzzle) == 0 || len(puzzle[0]) == 0 {
        return nil, fmt.Errorf("empty puzzle")
    }
    
    result := make(map[string][2][2]int)
    directions := getAllDirections()
    
    for _, word := range words {
        found := false
        for y := 0; y < len(puzzle); y++ {
            for x := 0; x < len(puzzle[0]); x++ {
            
                for _, dir := range directions {
                    if findWord(word, puzzle, x, y, dir) {
                        endX := x + (len(word)-1)*dir.dx
                        endY := y + (len(word)-1)*dir.dy
                        result[word] = [2][2]int{{x, y}, {endX, endY}}
                        found = true
                        break
                    }
                }
                if found {
                    break
                }
            }
            if found {
                break
            }
        }
        if !found {
            return nil, fmt.Errorf("word %s not found in puzzle", word)
        }
    }
    
    return result, nil
}