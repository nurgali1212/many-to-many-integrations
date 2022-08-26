package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sign_in/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendData(c *gin.Context) {
	var input model.Object
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postBody, _ := json.Marshal(input)
	responseBody := bytes.NewBuffer(postBody)
	log.Printf(string(postBody))
	resp, err := http.Post("http://localhost:8081/api/console", "application/json", responseBody)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
	c.JSON(http.StatusOK, sb)
}
