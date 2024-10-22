package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)


type team struct {
	name   string 
	played int    
	won    int    
	lost   int  
	drawn  int    
}


type teams map[string]*team

func (t *team) win() {
	t.played++
	t.won++
}

func (t *team) lose() {
	t.played++
	t.lost++
}

func (t *team) draw() {
	t.played++
	t.drawn++
}

func (t *team) points() int {
	return 3*t.won + 1*t.drawn
}

func (t teams) getOrCreateTeam(n string) *team {
	if t[n] == nil {
		t[n] = &team{name: n}
	}
	return t[n]
}


func (t teams) sort() []*team {
	var all []*team
	for _, team := range t {
		all = append(all, team)
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i].points() == all[j].points() {
			return all[i].name < all[j].name
		}
		return all[i].points() > all[j].points()
	})
	return all
}


func (t teams) fromResults(r io.Reader) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		if l == "" || l[0] == '#' { 
			continue
		}

		chunks := strings.Split(l, ";")
		if len(chunks) != 3 {
			return fmt.Errorf("wrong field count for line: %s (got %d fields)", l, len(chunks))
		}

		t1, t2, outcome := chunks[0], chunks[1], chunks[2]

		if !isValidTeamName(t1) || !isValidTeamName(t2) {
			return fmt.Errorf("invalid team name: %s or %s", t1, t2)
		}

		switch outcome {
		case "win":
			t.getOrCreateTeam(t1).win()
			t.getOrCreateTeam(t2).lose()
		case "loss":
			t.getOrCreateTeam(t1).lose()
			t.getOrCreateTeam(t2).win()
		case "draw":
			t.getOrCreateTeam(t1).draw()
			t.getOrCreateTeam(t2).draw()
		default:
			return fmt.Errorf("invalid outcome %s", outcome)
		}
	}
	return nil
}

func Tally(r io.Reader, w io.Writer) error {
	t := make(teams)
	err := t.fromResults(r)
	if err != nil {
		return err
	}


	_, err = fmt.Fprintf(w, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}

	for _, team := range t.sort() {
		_, err := fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n", team.name, team.played, team.won, team.drawn, team.lost, team.points())
		if err != nil {
			return err
		}
	}

	return nil
}

func isValidTeamName(name string) bool {
	for _, r := range name {
		if !(r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r == ' ' || r == '-') {
			return false
		}
	}
	return true
}
