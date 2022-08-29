package handler

import (
	"net/http"
	"sign_in/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createObjectHandler(ctx *gin.Context) {
	var input model.Object

	if err := ctx.BindJSON(&input); err != nil {
		h.service.SendDataService(err)
		ctx.IndentedJSON(http.StatusBadRequest, "Ошибка при обработке запроса")
		return
	}

	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{
		"response": http.StatusOK,
	})

}
