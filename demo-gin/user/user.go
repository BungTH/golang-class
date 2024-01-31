package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (uh *UserHandler) NewUser(c *gin.Context) {
	var u User

	err := c.ShouldBindJSON(&u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Database insert
	r := uh.db.Create(&u)

	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, " + u.Name + "!",
		"id":      u.Model.ID,
	})
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	var users []User

	// Get context value
	ctxName := c.GetString("Name")
	fmt.Println(ctxName)

	r := uh.db.Find(&users)

	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not convert ID to integer",
		})
		return
	}

	r := uh.db.Delete(&User{}, idInt)

	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	IdInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can not convert ID to integer",
		})
		return
	}

	var u User
	err = c.ShouldBindJSON(&u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := uh.db.Model(&User{}).Where("id = ?", IdInt).Update("name", u.Name)

	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated",
	})
}
