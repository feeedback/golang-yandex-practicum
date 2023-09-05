package main // Тема 1/4: Композитные типы → Урок 4/5
// https://practicum.yandex.ru/trainer/go-basics/lesson/b1ee4a40-18d6-44f2-ba62-b2918538f811/

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`

	Data []struct {
		Type string `json:"type"`
		ID   int    `json:"id"`

		Attributes struct {
			Email      string `json:"email"`
			ArticleIDs []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data"`
}

func ReadResponse(rawResp string) (Response, error) {
	var resp Response

	err := json.Unmarshal([]byte(rawResp), &resp)
	if err != nil {
		return Response{}, err
	}

	return resp, nil
}

func main() {
	rawJSON := `{
        "header": {
            "code": 0,
            "message": ""
        },
        "data": [{
            "type": "user",
            "id": 100,
            "attributes": {
                "email": "bob@yandex.ru",
                "article_ids": [10, 11, 12]
            }
        }]
    }`

	parsedResponse, err := ReadResponse(rawJSON)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", parsedResponse)
}
