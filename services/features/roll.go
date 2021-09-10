package features

import (
	"strconv"
	"strings"
	"tiamat/m/v0/common"
)

const RollMagicWord = "!roll"

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

	var rand int
	if wrongArg {
		rand = common.Roll()
	} else {
		rand = common.RollWithRange(int(from), int(to))
	}
	return strconv.FormatInt(int64(rand), 10)
}
