package api

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/")
	{
		v1.POST("/enterprise", Register)
	}

	return r
}
