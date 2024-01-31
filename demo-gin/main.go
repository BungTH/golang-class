package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bungkapth/demo-gin/auth"
	"github.com/bungkapth/demo-gin/user"
)

// date "+%Y-%m-%dT%H:%M:%S%Z:00"
// go build \
// -ldflags "-X main.buildcommit=`git rev-parse --short HEAD` \
// -X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
// -o maim

var (
	buildcommit = "dev"
	buildtime = time.Now().String()
)

func main() {

	_, err := os.Create("/tmp/healthy")
	if err != nil {
		log.Fatal("Can not create healthy file")
	}
	defer os.Remove("/tmp/healthy")

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Can not connect to database: %v", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		// c.String(http.StatusOK, "pong")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"buildcommit": buildcommit,
			"buildtime": buildtime,
			"version": "1.0.0",
		})
	})

	// Usage: http://localhost:8080/greeting?name=John
	r.GET("/greeting", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	// Usage: http://localhost:8080/greeting/John
	r.GET("/greeting/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	db.AutoMigrate(
		&user.User{},
	)

	r.GET("/token", auth.GetToken(os.Getenv("WIFI_SECRET")))

	authRoute := r.Group("", auth.AuthMiddleware())
	userHandler := user.NewUserHandler(db)

	// Usage: curl -X POST -H "Content-Type: application/json" -d '{"name":"John"}' http://localhost:8080/users
	authRoute.POST("/users", userHandler.NewUser)

	// Usage: curl http://localhost:8080/users
	authRoute.GET("/users", userHandler.GetUser)

	// Usage: curl -X DELETE http://localhost:8080/users/id
	authRoute.DELETE("/users/:id", userHandler.DeleteUser)

	// Usage: curl -X PATCH -H "Content-Type: application/json" -d '{"name":"John"}' http://localhost:8080/users/id
	authRoute.PATCH("/users/:id", userHandler.UpdateUser)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
		Handler: r,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Go routine
	go func () {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err.Error())
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("Shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err.Error())
	}

	// Usage: curl -X DELETE http://localhost:8080/users/id
	// r.DELETE("/users/:id", func(c *gin.Context) {
	// 	id := c.Param("id")

	// 	idInt, err := strconv.Atoi(id)

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Can not convert ID to integer",
	// 		})
	// 		return
	// 	}

	// 	var tmp []User
	// 	for _, user := range users {
	// 		if user.ID != idInt {
	// 			tmp = append(tmp, user)
	// 		}
	// 	}

	// 	users = tmp

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "User deleted",
	// 	})
	// })

	// // Usage: curl -X PATCH -H "Content-Type: application/json" -d '{"name":"John"}' http://localhost:8080/users/id
	// r.PATCH("/users/:id", func(c *gin.Context) {
	// 	id := c.Param("id")

	// 	idInt, err := strconv.Atoi(id)

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Can not convert ID to integer",
	// 		})
	// 		return
	// 	}

	// 	var u User

	// 	err = c.ShouldBindJSON(&u)

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	for i, a := range users {
	// 		if a.ID == idInt {
	// 			users[i].Name = u.Name
	// 			break
	// 		}
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "User updated",
	// 	})
	// })

	r.Run()
}
