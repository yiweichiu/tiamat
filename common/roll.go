package common

import (
	"math/rand"
	"time"
)

const defaultRangeFrom = 0
const defaultRangeTo = 100

var (
	isRolled = false
	seed     *rand.Rand
)

func Roll() int {
	return RollWithRange(defaultRangeFrom, defaultRangeTo)
}

func RollWithRange(from, to int) int {
	if from > to {
		from = defaultRangeFrom
		to = defaultRangeTo
	}

	if !isRolled {
		src := rand.NewSource(time.Now().UnixNano())
		isRolled = true
		seed = rand.New(src)
	}

	return seed.Intn(to) + from
}
