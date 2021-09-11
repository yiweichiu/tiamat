package line

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tiamat/m/v0/models"
)

func GetUserName(userId string) string {
	client := http.Client{}
	url := fmt.Sprintf("%s/profile/%s", domain, userId)
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
	userJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	user := models.LineProfile{}
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		log.Print(err.Error())
		return ""
	}
	return user.Name()
}
