package features

import (
	"log"
	"strings"

	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

const japan = "ja"
const english = "en"

/*
	Is now not functional due to google policy,
	possible solution is to used another `python` module called "googletrans"
*/

func Translate(msg string) string {
	tokens := strings.Split(msg, " ")

	if len(tokens) < 3 {
		return ""
	}

	targetStr := ""
	var from, to language.Tag

	switch tokens[1] {
	case japan:
		targetIdx := strings.Index(msg, japan)
		targetStr = msg[targetIdx:]
	case english:
		targetIdx := strings.Index(msg, english)
		targetStr = msg[targetIdx:]
	default:
		targetIdx := strings.Index(msg, " ")
		targetStr = msg[targetIdx:]
	}

	result, err := gtranslate.Translate(targetStr, from, to)
	if err != nil {
		log.Print(err)
		return ""
	}

	return result
}
