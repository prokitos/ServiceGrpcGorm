package services

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var secretKey string = "gpt45"

// отправка данных на второй сервер
func sendToSecond() (result []byte) {

	baseURL, _ := url.Parse("http://localhost:8112/getter")

	data := url.Values{
		"name":       {""},
		"surname":    {""},
		"patronymic": {""},
		"key":        {secretKey},
	}

	resp, _ := http.PostForm(baseURL.String(), data)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body

}

// структура, с помощью которой идёт обмен внутри сервисов, и потом в базу
type send_Owner struct {
	Owner_id   int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Secret     string `json:"key"`
}
