package routes

import (
	"famima-transactional/controllers"
	"famima-transactional/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Rute yang dilindungi
	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware()) // Panggil middleware sebagai fungsi
	{
		protected.GET("/data", func(c *gin.Context) {
			c.String(200, "Data Karyawan Terlindungi")
		})

		// Rute untuk absensi
		protected.POST("/absensi", controllers.Absensi)
	}
}