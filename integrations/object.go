package integrations

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sign_in/model"
)

func (c *client) SendData(err error) error {
	errorMessage := &model.ErrorMessage{Message: err.Error()}
	postBody, _ := json.Marshal(errorMessage)
	requestBody := bytes.NewBuffer(postBody)
	_, err = http.Post("http://localhost:8081/api/console", "application/json", requestBody)

	return err

}
