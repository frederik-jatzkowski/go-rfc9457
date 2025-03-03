package luck

import (
	"errors"
	"math/rand"
)

var (
	ErrOutOfLuck = errors.New("out of luck")
)

type LuckRepo struct {
}

func (r LuckRepo) FindLuck() (bool, error) {
	if rand.Intn(2) == 1 {
		return true, nil
	}

	return false, ErrOutOfLuck
}
