package features

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const rateUrl = "https://rate.bot.com.tw/xrt?Lang=zh-TW"

func Rate(msg string) string {
	tokens := strings.Split(msg, " ")
	c := colly.NewCollector()

	if len(tokens) == 1 {
		return rateHelp()
	}

	var transfer int64
	if len(tokens) > 2 {
		tr, err := strconv.ParseInt(tokens[2], 10, 0)
		if err != nil {
			log.Print("Wrong price transfer from user :" + tokens[2])
			return "wrong information"
		}
		transfer = tr
	}

	targetCurrency := tokens[1]
	spotBuy := ""
	spotSell := ""
	cashBuy := ""
	cashSell := ""
	isPrice := false

	c.OnHTML("table[title='牌告匯率']", func(e *colly.HTMLElement) {
		tableBody := e.DOM.Find("tbody")
		tableBody.Each(func(cnt int, body *goquery.Selection) {
			tableRaw := body.Find("tr")
			tableRaw.Each(func(cnt int, raw *goquery.Selection) {
				typeColumn := raw.Find("td[data-table='幣別']")
				typeText := typeColumn.Find("div[class='visible-phone print_hide']")
				if strings.Contains(typeText.Text(), strings.ToLower(targetCurrency)) || strings.Contains(typeText.Text(), strings.ToUpper(targetCurrency)) {
					isPrice = true
					cashBuy = strings.TrimSpace(raw.Find("td[data-table='本行現金買入']").First().Text())
					cashSell = strings.TrimSpace(raw.Find("td[data-table='本行現金賣出']").First().Text())
					spotBuy = strings.TrimSpace(raw.Find("td[data-table='本行即期買入']").First().Text())
					spotSell = strings.TrimSpace(raw.Find("td[data-table='本行即期賣出']").First().Text())
				}
			})
		})
	})
	c.Visit(rateUrl)

	if isPrice {
		resp := ""
		resp += strings.ToUpper(targetCurrency)
		resp += "\n"
		resp += "即期買入 : " + spotBuy + "\n"
		resp += "即期賣出 : " + spotSell + "\n"
		resp += "現金買入 : " + cashBuy + "\n"
		resp += "現金賣出 : " + cashSell + "\n"
		if transfer != 0 {
			var price float64
			var err error
			if spotBuy == "-" {
				price, err = strconv.ParseFloat(spotBuy, 0)
				if err != nil {
					log.Print("Wrong price transfer from crawler :" + spotBuy)
					return "wrong information"
				}
			} else {
				price, err = strconv.ParseFloat(cashBuy, 0)
				if err != nil {
					log.Print("Wrong price transfer from crawler :" + cashBuy)
					return "wrong information"
				}
			}
			result := float64(transfer) * price
			resp += "換算台幣 : " + strconv.FormatFloat(result, 'f', 0, 64)
		}
		return resp
	}
	return rateHelp()
}

func rateHelp() string {
	var helpMsg string
	helpMsg += "usd 美元\n"
	helpMsg += "hkd 港幣\n"
	helpMsg += "gbp 英鎊\n"
	helpMsg += "aud 澳大利亞元\n"
	helpMsg += "cad 加拿大幣\n"
	helpMsg += "sgd 新加坡幣\n"
	helpMsg += "chf 瑞士法郎\n"
	helpMsg += "jpy 日幣\n"
	helpMsg += "zar 南非鍰\n"
	helpMsg += "sek 瑞典克朗\n"
	helpMsg += "nzd 紐西蘭元\n"
	helpMsg += "thb 泰銖\n"
	helpMsg += "php 菲律賓披索\n"
	helpMsg += "idr 印尼盾\n"
	helpMsg += "eur 歐元\n"
	helpMsg += "krw 韓圓\n"
	helpMsg += "vnd 越南盾\n"
	helpMsg += "myr 令吉\n"
	helpMsg += "cny 人民幣\n"
	return helpMsg
}
