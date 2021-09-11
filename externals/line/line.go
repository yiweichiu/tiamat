package line

import (
	"io"
	"net/http"
)

const domain = "https://api.line.me/v2/bot"
const channelSecret = "8273b1607e7aad58ea027dda7dbcc57c"
const channelAccessToken = "uzNr4czd17pPxHKZ2aJ3erVkBU0XK7NcAYItFwqSEMMgR85BawivpWKMo4cFkSJTSjRERQv5XEMTXfzeKB5T1GTEGJwju80ZDQqJQzCBQTUTIr4860hOAyCeJFb1597sRb58kxD6HbcS+Vw1Y39AhwdB04t89/1O/w1cDnyilFU="

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+channelAccessToken)
	return req, nil
}

func GetChannelAccessToken() string {
	return channelAccessToken
}

func GetChannelSecret() string {
	return channelSecret
}
