package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseJSON(c *gin.Context, p interface{}, status int) {
	ubahkeByte, err := json.Marshal(p)

	if err != nil {
		// http.Error(c, "error om", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, "error")
	}

	c.JSON(200, ubahkeByte)
}
