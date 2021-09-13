package yahoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tiamat/m/v0/models"
)

const getSymbolQuote = "%s/v6/finance/quote?region=%s&lang=%s&symbols=%s"
const region = "TW"
const lang = "zh"

func GetSymbolQuote(symbol string) string {
	client := http.Client{}
	url := fmt.Sprintf(getSymbolQuote, baseUrl, region, lang, symbol)
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	defer resp.Body.Close()
	symbolJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	data := models.YahooQuotes{}
	err = json.Unmarshal(symbolJson, &data)
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	if data.QuoteResponse.Results == nil {
		log.Printf("wrong symbol type : %s", symbol)
		return ""
	}

	return fmt.Sprintf("%s : %s\n%s : %s\n%s : %.2f", "股票名稱", data.QuoteResponse.Results[0].LongName, "股票代碼", symbol, "股票價格", data.QuoteResponse.Results[0].RegularMarketPrice)
}
