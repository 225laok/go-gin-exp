package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gin-example/pkg/setting"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppS.PageSize
	}
	return result
}
