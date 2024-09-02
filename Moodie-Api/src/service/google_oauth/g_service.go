package google_oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserInfo(tkn string) (map[string]interface{}, error) {
	userInfoEndPoint := "https://www.googleapis.com/oauth2/v2/userinfo"

	res, err := http.Get(fmt.Sprintf("%s?access_token=%s", userInfoEndPoint, tkn)) // https://www.googleapis.com/oauth2/v2/userinfo?access_token=your_access_token
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var userInfo map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	fmt.Println(userInfo["email"], userInfo["id"])
	return userInfo, nil

}
