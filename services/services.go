package services

import "github.com/gin-gonic/gin"

func Run() {
	r := gin.Default()
	r.GET("/linebot", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	// client := &http.Client{}
	// bot, err := linebot.New(linebotSecret, linebotToken, linebot.WithHTTPClient(client))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// req := http.Request{}
	// for {
	// 	events, err := bot.ParseRequest(req)
	// }
}
