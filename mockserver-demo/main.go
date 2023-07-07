package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type AidResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func responseBody(code int, msg string, data interface{}) *AidResponse {
	return &AidResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	app := gin.Default()
	rootGrp := app.Group("/api/v1")
	rootGrp.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	rootGrp.POST("/result/upload", uploadResult)

	srv := &http.Server{
		Addr:    ":58201",
		Handler: app,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Printf("listen: %s\n", err.Error())
			} else {
				log.Printf("start fail: %s\n", err.Error())
				os.Exit(1)
			}
		}
	}()
	//listener, err = net.Listen("tcp", )

	log.Println("listen at :58201")
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	log.Println("Shutting down server, notify to do clean work...")

	cancelFunc()
	// The context is used to inform the server it has 3 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err.Error())
	}

	<-ctx.Done()
}

type MyFiles struct {
	Files []*multipart.FileHeader `form:"xxx" binding:"required"`
}

func uploadResult(c *gin.Context) {
	//var myFiles MyFiles
	//
	//if err := c.ShouldBind(&myFiles); err != nil {
	//	c.JSON(500, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, responseBody(http.StatusBadRequest, err.Error(), nil))
		return
	}
	files := form.File["files"]
	//files, err := c.FormFile("xxx")
	if err != nil {
		c.JSON(http.StatusOK, responseBody(http.StatusBadRequest, err.Error(), nil))
		return
	}

	basePath := "./upload/"

	for _, file := range files {
		filename := basePath + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusOK, responseBody(http.StatusBadRequest, err.Error(), nil))

			return
		}
		log.Printf("the result file is successfully received. %s\n", filename)
	}

	c.JSON(http.StatusOK, responseBody(http.StatusOK, "ok", nil))
}
