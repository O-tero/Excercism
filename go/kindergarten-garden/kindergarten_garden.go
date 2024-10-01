package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

type Garden map[string][]string

func parsePlant(b byte) (string, error) {
	switch b {
	case 'V':
		return "violets", nil
	case 'R':
		return "radishes", nil
	case 'C':
		return "clover", nil
	case 'G':
		return "grass", nil
	default:
		return "", fmt.Errorf("invalid plant '%c'", b)
	}
}

func parsePlants(bytes ...byte) ([]string, error) {
	plants := make([]string, 0, len(bytes))
	for _, b := range bytes {
		p, err := parsePlant(b)
		if err != nil {
			return nil, err
		}
		plants = append(plants, p)
	}
	return plants, nil
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(diagram) == 0 || diagram[0] != '\n' {
		return nil, fmt.Errorf("diagram must start with a newline")
	}
	rows := strings.Split(diagram[1:], "\n")
	if len(rows) != 2 {
		return nil, fmt.Errorf("diagram has %d rows, want 2", len(rows))
	}
	if len(rows[0]) != len(rows[1]) || len(rows[0]) != len(children)*2 {
		return nil, fmt.Errorf("invalid diagram length")
	}
	g := Garden{}
	// Make a copy to avoid modifying input slice
	sortedChildren := sort.StringSlice(append([]string{}, children...))
	sortedChildren.Sort()
	for i, child := range sortedChildren {
		if _, ok := g[child]; ok {
			return nil, fmt.Errorf("duplicate child \"%s\"", child)
		}
		p, err := parsePlants(rows[0][i*2], rows[0][i*2+1], rows[1][i*2], rows[1][i*2+1])
		if err != nil {
			return nil, err
		}
		g[child] = p
	}
	return &g, nil
}

func (g Garden) Plants(child string) ([]string, bool) {
	plants, ok := g[child]
	return plants, ok
}