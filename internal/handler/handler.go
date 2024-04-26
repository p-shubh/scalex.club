package handler

import (
	"net/http"
	"scalex/internal/auth"
	"scalex/internal/storage"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// type User struct {
// 	Username string
// 	Password string
// 	UserType string
// }

// var users = []User{
// 	{Username: "admin", Password: "admin123", UserType: "admin"},
// 	{Username: "user", Password: "user123", UserType: "regular"},
// }

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userType, err := storage.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	token, err := auth.GenerateToken(req.Username, userType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Home(c *gin.Context) {
	userType := c.MustGet("userType").(string)
	books, err := storage.GetBooks(userType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve books",
			"info":  err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

type addBookRequest struct {
	BookName        string `json:"book_name"`
	Author          string `json:"author"`
	PublicationYear string `json:"publication_year"`
}

type deleteBookRequest struct {
	BookName string `json:"book_name"`
}

func AddBook(c *gin.Context) {
	var req addBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userType := c.MustGet("userType").(string)
	if userType != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only admin can add books"})
		return
	}

	err := storage.AddBook(userType, req.BookName, req.Author, req.PublicationYear)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book added successfully"})
}

func DeleteBook(c *gin.Context) {
	var req deleteBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userType := c.MustGet("userType").(string)
	if userType != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only admin can delete books"})
		return
	}

	err := storage.DeleteBook(userType, req.BookName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}
