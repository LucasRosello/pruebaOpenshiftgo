package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//SetupRouter - Definicion de las rutas de las aplicaciones
func SetupRouter(r *gin.Engine) {

	v2 := r.Group("/api/v2")
	{
		envios := v2.Group("/example")
		{
			envios.GET(":id", Example)
		}
	}
	r.GET("/api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // ../api/doc/index.html
}
