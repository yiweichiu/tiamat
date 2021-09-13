package features

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const defaultRangeFrom = 0
const defaultRangeTo = 100

func Roll(msg string) string {
	tokens := strings.Split(msg, " ")
	wrongArg := false

	var from, to int64
	var err error

	if len(tokens) < 3 {
		wrongArg = true
	} else {
		from, err = strconv.ParseInt(tokens[1], 10, 0)
		if err != nil {
			wrongArg = true
		}

		to, err = strconv.ParseInt(tokens[2], 10, 0)
		if err != nil {
			wrongArg = true
		}
	}

	if from < 0 || to < 0 {
		return "不要搞啦!!"
	}

	if from == to {
		wrongArg = true
	}

	var rand int
	if wrongArg {
		rand = roll(defaultRangeFrom, defaultRangeTo)
	} else {
		rand = roll(int(from), int(to))
	}
	return strconv.FormatInt(int64(rand), 10)
}

var (
	isRolled = false
	seed     *rand.Rand
)

func roll(from, to int) int {
	if from > to {
		from = defaultRangeFrom
		to = defaultRangeTo
	}

	if !isRolled {
		src := rand.NewSource(time.Now().UnixNano())
		isRolled = true
		seed = rand.New(src)
	}

	return seed.Intn(to-from) + from
}
