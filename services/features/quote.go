package features

import (
	"strings"
	"tiamat/m/v0/externals/yahoo"
)

func Quote(msg string) string {
	tokens := strings.Split(msg, " ")

	if len(tokens) < 2 {
		return quoteHelp()
	}

	if len(tokens) > 2 {
		return "不要搞啦!!"
	}
	return yahoo.GetSymbolQuote(tokens[1])
}

func quoteHelp() string {
	var helpMsg string
	helpMsg += "台股 !quote [股票代號].TW\n"
	helpMsg += "美股 !quote [股票代號]"
	return helpMsg
}
