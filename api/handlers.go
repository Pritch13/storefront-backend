package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func getItemsHandler(c *gin.Context) {
	var items = []Item{
		{
			"Cervelo",
			2213,
		},
		{
			"Canyon",
			3321,
		},
	}
	c.JSON(http.StatusOK, items)

}
