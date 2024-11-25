package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Robot struct {
	name string
}

var (
	usedNames = make(map[string]bool)
	mu        sync.Mutex
)

const MaxNames = 26 * 26 * 1000 

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (r *Robot) Name() (string, error) {
	mu.Lock()
	defer mu.Unlock()

	if r.name != "" {
		return r.name, nil
	}

	if len(usedNames) >= MaxNames {
		return "", errors.New("all possible names have been allocated")
	}

	for {
		newName := generateRandomName()
		if !usedNames[newName] {
			r.name = newName
			usedNames[newName] = true
			break
		}
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	mu.Lock()
	defer mu.Unlock()


	if r.name != "" {
		delete(usedNames, r.name)
		r.name = ""
	}
}

func generateRandomName() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return fmt.Sprintf("%c%c%03d",
		letters[rand.Intn(26)],
		letters[rand.Intn(26)],
		rand.Intn(1000))
}
