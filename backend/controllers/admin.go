package controllers

import (
	"culturstone/config"
	"culturstone/models"
	"culturstone/utils" // Pastikan import ini ada untuk akses ke GenerateToken
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AdminRegister godoc
// @Summary Register a new admin
// @Description Creates a new administrator account.
// @Tags Public
// @Accept  json
// @Produce  json
// @Param   admin  body   AdminInput  true  "Admin Credentials"
// @Success 201 {object} models.Admin
// @Failure 400 {object} gin.H
// @Router /register [post]
func AdminRegister(c *gin.Context) {
	var input AdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	admin := models.Admin{Username: input.Username, Password: string(hashed)}
	if err := config.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Admin registered", "admin": admin})
}

// AdminLogin godoc
// @Summary Login as an admin
// @Description Authenticates an administrator and returns a token.
// @Tags Public
// @Accept  json
// @Produce  json
// @Param   admin  body   AdminInput  true  "Admin Credentials"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /login [post]
func AdminLogin(c *gin.Context) {
	var input AdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var admin models.Admin
	// Cari admin berdasarkan username
	if err := config.DB.Where("username = ?", input.Username).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// --- GENERATE JWT TOKEN ---
	token, err := utils.GenerateToken(admin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Kembalikan token ke frontend
	c.JSON(http.StatusOK, gin.H{
		"message":  "Login successful",
		"token":    token,
		"username": admin.Username,
	})
}

// GetAdminStats godoc
// @Summary Get dashboard statistics
// @Description Retrieves a summary of website data for the admin dashboard.
// @Tags Admin
// @Produce  json
// @Success 200 {object} gin.H
// @Router /admin/stats [get]
func GetAdminStats(c *gin.Context) {
	var tProd, tMsg, tPort, tTestimoni, tCat, tPCat int64
	var visitors []models.Visitor
	config.DB.Model(&models.Product{}).Count(&tProd)
	config.DB.Model(&models.Message{}).Count(&tMsg)
	config.DB.Model(&models.Portfolio{}).Count(&tPort)
	config.DB.Model(&models.Testimoni{}).Count(&tTestimoni)
	config.DB.Model(&models.Category{}).Count(&tCat)
	config.DB.Model(&models.ProductCategory{}).Count(&tPCat)
	config.DB.Order("date desc").Limit(7).Find(&visitors)
	c.JSON(http.StatusOK, gin.H{
		"products":           tProd,
		"messages":           tMsg,
		"portfolios":         tPort,
		"testimoni":          tTestimoni,
		"categories":         tCat,
		"product_categories": tPCat,
		"visitors":           visitors,
	})
}
// GetMessages godoc
// @Summary Get all contact messages
// @Description Retrieves a list of all messages submitted through the contact form.
// @Tags Admin
// @Produces  json
// @Success 200 {array} models.Message
// @Router /admin/messages [get]
func GetMessages(c *gin.Context) {
	var msgs []models.Message
	config.DB.Order("created_at desc").Find(&msgs)
	c.JSON(http.StatusOK, msgs)
}

// SaveProduct godoc
// @Summary Create or update a product
// @Description Creates a new product or updates an existing one if an ID is provided.
// @Tags Admin
// @Consumes  multipart/form-data
// @Produces  json
// @Param id path int false "Product ID"
// @Param name formData string true "Product Name"
// @Param description formData string true "Product Description"
// @Param category formData string true "Product Category"
// @Param quality formData string false "Product Quality"
// @Param size formData string false "Product Size"
// @Param finishing formData string false "Product Finishing"
// @Param image formData file false "Product Image"
// @Success 200 {object} models.Product
// @Router /admin/products [post]
// @Router /admin/products/{id} [put]
func SaveProduct(c *gin.Context) {
	var p models.Product
	id := c.Param("id")
	if id != "" {
		config.DB.First(&p, id)
	}

	p.Name = c.PostForm("name")
	p.Description = c.PostForm("description")
	p.Category = c.PostForm("category")
	p.Quality = c.PostForm("quality")
	p.Size = c.PostForm("size")
	p.Finishing = c.PostForm("finishing")

	file, _ := c.FormFile("image")
	if file != nil {
		filename := fmt.Sprintf("prod_%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		c.SaveUploadedFile(file, "./uploads/"+filename)
		p.ImageURL = "http://localhost:8080/uploads/" + filename
	}
	if id != "" {
		config.DB.Save(&p)
	} else {
		config.DB.Create(&p)
	}
	c.JSON(http.StatusOK, p)
}

// SaveCategory godoc
// @Summary Create or update a category
// @Description Creates a new category or updates an existing one if an ID is provided.
// @Tags Admin
// @Consumes  multipart/form-data
// @Produces  json
// @Param id path int false "Category ID"
// @Param name formData string true "Category Name"
// @Param description formData string false "Category Description"
// @Param image formData file false "Category Image"
// @Success 200 {object} models.Category
// @Router /admin/categories [post]
// @Router /admin/categories/{id} [put]
func SaveCategory(c *gin.Context) {
	var cat models.Category
	id := c.Param("id")
	if id != "" {
		config.DB.First(&cat, id)
	}

	cat.Name = c.PostForm("name")
	cat.Description = c.PostForm("description")

	file, _ := c.FormFile("image")
	if file != nil {
		filename := fmt.Sprintf("cat_%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		c.SaveUploadedFile(file, "./uploads/"+filename)
		cat.ImageURL = "http://localhost:8080/uploads/" + filename
	}
	if id != "" {
		config.DB.Save(&cat)
	} else {
		config.DB.Create(&cat)
	}
	c.JSON(http.StatusOK, cat)
}

// SavePortfolio godoc
// @Summary Create or update a portfolio item
// @Description Creates a new portfolio item or updates an existing one if an ID is provided.
// @Tags Admin
// @Consumes  multipart/form-data
// @Produces  json
// @Param id path int false "Portfolio ID"
// @Param title formData string true "Portfolio Title"
// @Param description formData string false "Portfolio Description"
// @Param image formData file false "Portfolio Image"
// @Success 200 {object} models.Portfolio
// @Router /admin/portfolios [post]
// @Router /admin/portfolios/{id} [put]
func SavePortfolio(c *gin.Context) {
	var pf models.Portfolio
	id := c.Param("id")
	if id != "" {
		config.DB.First(&pf, id)
	}

	pf.Title = c.PostForm("title")
	pf.Description = c.PostForm("description")

	file, _ := c.FormFile("image")
	if file != nil {
		filename := fmt.Sprintf("port_%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		c.SaveUploadedFile(file, "./uploads/"+filename)
		pf.ImageURL = "http://localhost:8080/uploads/" + filename
	}
	if id != "" {
		config.DB.Save(&pf)
	} else {
		config.DB.Create(&pf)
	}
	c.JSON(http.StatusOK, pf)
}


// SaveTestimoni godoc
// @Summary Create or update a testimonial
// @Description Creates a new testimonial or updates an existing one if an ID is provided.
// @Tags Admin
// @Consumes  multipart/form-data
// @Produces  json
// @Param id path int false "Testimoni ID"
// @Param client_name formData string true "Client Name"
// @Param testimonial_text formData string true "Testimonial Text"
// @Param portfolio_id formData int false "Associated Portfolio ID"
// @Param image formData file false "Client Image"
// @Param video formData file false "Testimonial Video"
// @Success 200 {object} models.Testimoni
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /admin/testimoni [post]
// @Router /admin/testimoni/{id} [put]
func SaveTestimoni(c *gin.Context) {
	var t models.Testimoni
	id := c.Param("id")
	if id != "" {
		if err := config.DB.First(&t, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Testimoni not found"})
			return
		}
	}

	t.ClientName = c.PostForm("client_name")
	t.TestimonialText = c.PostForm("testimonial_text")

	portfolioIDStr := c.PostForm("portfolio_id")
	if portfolioIDStr != "" {
		portfolioID, err := strconv.ParseUint(portfolioIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid portfolio_id format"})
			return
		}
		t.PortfolioID = uint(portfolioID)
	} else {
		// Opsional: set 0 atau biarkan value lama jika tidak diupdate, 
		// tapi logika ini akan menset 0 jika string kosong dikirim.
		if id == "" {
			t.PortfolioID = 0
		}
	}

	file, err := c.FormFile("image")
	if err == nil {
		filename := fmt.Sprintf("testimoni_img_%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		uploadPath := "./uploads/" + filename
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image: " + err.Error()})
			return
		}
		t.ImagePath = "http://localhost:8080/uploads/" + filename
	}

	videoFile, err := c.FormFile("video")
	if err == nil {
		videoFilename := fmt.Sprintf("testimoni_video_%d_%s", time.Now().Unix(), filepath.Base(videoFile.Filename))
		videoUploadPath := "./uploads/" + videoFilename
		if err := c.SaveUploadedFile(videoFile, videoUploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video: " + err.Error()})
			return
		}
		t.VideoPath = "http://localhost:8080/uploads/" + videoFilename
	}

	if id != "" {
		if err := config.DB.Save(&t).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update testimoni: " + err.Error()})
			return
		}
	} else {
		if err := config.DB.Create(&t).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create testimoni: " + err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, t)
}


// AdminGetTestimoni godoc
// @Summary Get all testimonials for admin
// @Description Retrieves a list of all testimonials.
// @Tags Admin
// @Produces json
// @Success 200 {array} models.Testimoni
// @Router /admin/testimoni [get]
func AdminGetTestimoni(c *gin.Context) {
	var testimoni []models.Testimoni
	config.DB.Order("id desc").Preload("Portfolio").Find(&testimoni)
	c.JSON(http.StatusOK, testimoni)
}

// Generic Delete
func DeleteEntity(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.DB.Delete(model, c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	}
}

// GetVisitors godoc
// @Summary Get all visitor records
// @Description Retrieves a list of daily visitor counts.
// @Tags Admin
// @Produces json
// @Success 200 {array} models.Visitor
// @Router /admin/visitors [get]
func GetVisitors(c *gin.Context) {
	var visitors []models.Visitor
	config.DB.Order("date desc").Find(&visitors)
	c.JSON(http.StatusOK, visitors)
}