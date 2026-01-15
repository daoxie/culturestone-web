package main

import (
	"culturstone/config"
	"culturstone/controllers"
	_ "culturstone/docs" // Dokumentasi Swagger
	"culturstone/middlewares" // Import middleware autentikasi yang baru dibuat
	"culturstone/models"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Culturstone API
// @version 1.0
// @description This is the API documentation for the Culturstone web application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func main() {
	// 1. Koneksi Database
	config.ConnectDB()

	// 2. Inisialisasi Router
	r := gin.Default()

	// 3. Buat folder uploads jika belum ada
	os.MkdirAll("./uploads", os.ModePerm)

	// 4. Konfigurasi CORS (Penting untuk Frontend React)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true // Izinkan semua origin untuk Vercel
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	// 5. Middleware Global (Visitor Counter)
	r.Use(controllers.TrackVisitor)

	// 6. Serve Static Files (Gambar/Video)
	r.Static("/uploads", "./uploads")

	// 7. Swagger Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 8. Definisi Routes
	api := r.Group("/api")
	{
		// --- PUBLIC ROUTES (Tanpa Login) ---
		api.GET("/products", controllers.GetProducts)
		api.GET("/portfolios", controllers.GetPortfolios)
		api.GET("/testimoni", controllers.GetTestimoni)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/product-categories", controllers.GetProductCategories)
		api.POST("/contact", controllers.PostContact)
		
		// Auth Routes (Login & Register tetap Public)
		api.POST("/login", controllers.AdminLogin)
		api.POST("/register", controllers.AdminRegister)

		// --- ADMIN ROUTES (DILINDUNGI JWT) ---
		admin := api.Group("/admin")
		
		// Terapkan Middleware Autentikasi di sini
		admin.Use(middlewares.JwtAuthMiddleware()) 
		{
			// Dashboard Stats
			admin.GET("/stats", controllers.GetAdminStats)
			
			// Messages
			admin.GET("/messages", controllers.GetMessages)
			// @Summary Delete a contact message
			// @Router /admin/messages/{id} [delete]
			admin.DELETE("/messages/:id", controllers.DeleteEntity(&models.Message{}))

			// Visitors
			admin.GET("/visitors", controllers.GetVisitors)

			// CRUD Products
			admin.POST("/products", controllers.SaveProduct)
			admin.PUT("/products/:id", controllers.SaveProduct)
			// @Summary Delete a product
			// @Router /admin/products/{id} [delete]
			admin.DELETE("/products/:id", controllers.DeleteEntity(&models.Product{}))

			// CRUD Categories
			admin.POST("/categories", controllers.SaveCategory)
			admin.PUT("/categories/:id", controllers.SaveCategory)
			// @Summary Delete a category
			// @Router /admin/categories/{id} [delete]
			admin.DELETE("/categories/:id", controllers.DeleteEntity(&models.Category{}))

			// CRUD Portfolios
			admin.POST("/portfolios", controllers.SavePortfolio)
			admin.PUT("/portfolios/:id", controllers.SavePortfolio)
			// @Summary Delete a portfolio item
			// @Router /admin/portfolios/{id} [delete]
			admin.DELETE("/portfolios/:id", controllers.DeleteEntity(&models.Portfolio{}))

			// CRUD Testimoni (Testimonials)
			admin.GET("/testimoni", controllers.AdminGetTestimoni)
			admin.POST("/testimoni", controllers.SaveTestimoni)
			admin.PUT("/testimoni/:id", controllers.SaveTestimoni)
			// @Summary Delete a testimonial
			// @Router /admin/testimoni/{id} [delete]
			admin.DELETE("/testimoni/:id", controllers.DeleteEntity(&models.Testimoni{}))
		}
	}

	r.Run()
}

/* package main

import (
	"culturstone/config"
	"culturstone/controllers"
	_ "culturstone/docs"
	"culturstone/middlewares"
	"culturstone/models"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 1. Koneksi Database
	config.ConnectDB()

	// 2. Inisialisasi Router
	r := gin.Default()

	// 3. Buat folder uploads
	os.MkdirAll("./uploads", os.ModePerm)

	// 4. KONFIGURASI CORS (BAGIAN PENTING!)
	// Ganti string di bawah ini dengan URL Frontend (Port 5173) Anda dari VS Code
	frontendURL := "https://mw9qv5xz-5173.asse.devtunnels.ms" 

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173", // Untuk local
			frontendURL,             // Untuk akses HP/Public (DARI VS CODE PORT 5173)
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 5. Middleware & Static
	r.Use(controllers.TrackVisitor)
	r.Static("/uploads", "./uploads")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 6. Routes
	// Perhatikan: Semua route ada di dalam Group "/api"
	api := r.Group("/api")
	{
		// Public
		api.GET("/products", controllers.GetProducts)
		api.GET("/portfolios", controllers.GetPortfolios)
		api.GET("/testimoni", controllers.GetTestimoni)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/product-categories", controllers.GetProductCategories)
		api.POST("/contact", controllers.PostContact)
		
		// Auth
		api.POST("/login", controllers.AdminLogin)       // Endpoint jadi: /api/login
		api.POST("/register", controllers.AdminRegister) // Endpoint jadi: /api/register

		// Admin (Protected)
		admin := api.Group("/admin")
		admin.Use(middlewares.JwtAuthMiddleware()) 
		{
			admin.GET("/stats", controllers.GetAdminStats)
			admin.GET("/messages", controllers.GetMessages)
			admin.DELETE("/messages/:id", controllers.DeleteEntity(&models.Message{}))
			admin.GET("/visitors", controllers.GetVisitors)

			admin.POST("/products", controllers.SaveProduct)
			admin.PUT("/products/:id", controllers.SaveProduct)
			admin.DELETE("/products/:id", controllers.DeleteEntity(&models.Product{}))

			admin.POST("/categories", controllers.SaveCategory)
			admin.PUT("/categories/:id", controllers.SaveCategory)
			admin.DELETE("/categories/:id", controllers.DeleteEntity(&models.Category{}))

			admin.POST("/portfolios", controllers.SavePortfolio)
			admin.PUT("/portfolios/:id", controllers.SavePortfolio)
			admin.DELETE("/portfolios/:id", controllers.DeleteEntity(&models.Portfolio{}))

			admin.GET("/testimoni", controllers.AdminGetTestimoni)
			admin.POST("/testimoni", controllers.SaveTestimoni)
			admin.PUT("/testimoni/:id", controllers.SaveTestimoni)
			admin.DELETE("/testimoni/:id", controllers.DeleteEntity(&models.Testimoni{}))
		}
	}

	r.Run(":8080")
}*/