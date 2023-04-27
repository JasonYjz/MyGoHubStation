package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type AppInfo struct {
	PKey       string            `json:"pKey,omitempty"`
	PID        int               `json:"pID,omitempty"`
	PType      int               `json:"pType,omitempty"`
	PPort      int               `json:"pPort,omitempty"`
	Additional map[string]string `json:"additional,omitempty"`
}

func main() {
	fmt.Println("the server is started.")
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	app := gin.Default()

	app.Use(gin.Recovery())

	rootGrp := app.Group("/api/v1")
	rootGrp.POST("/app", func(c *gin.Context) {
		fmt.Println("receive a request to register new app")

		info := &AppInfo{}

		err := c.ShouldBindJSON(info)
		if err != nil {
			fmt.Printf("error occurs in bind request body. err:%s\n", err.Error())
			c.JSON(http.StatusBadRequest, "ERR")
			return
		}

		fmt.Printf("%v\n", info)

		c.JSON(http.StatusOK, "OK")

	})

	rootGrp.GET("/app", func(c *gin.Context) {
		fmt.Println("receive a request to query existed app")

		pValue, b := c.GetQuery("pKey")
		if b {
			fmt.Printf("pKey=%s\n", pValue)
		}

		pValue, b = c.GetQuery("pID")
		if b {
			fmt.Printf("pID=%s\n", pValue)
		}

		c.JSON(http.StatusOK, "OK")
	})

	rootGrp.DELETE("/app", func(c *gin.Context) {
		fmt.Println("receive a request to delete existed app")

		//info := &AppInfo{}
		//
		//err := c.ShouldBindJSON(info)
		//if err != nil {
		//	fmt.Printf("error occurs in bind request body. err:%s\n", err.Error())
		//	c.JSON(http.StatusBadRequest, "ERR")
		//	return
		//}
		//
		//fmt.Printf("pKey=%s, pID=%d\n", info.PKey, info.PID)

		pValue, b := c.GetQuery("pKey")
		if b {
			fmt.Printf("pKey=%s\n", pValue)
		}

		pValue, b = c.GetQuery("pID")
		if b {
			fmt.Printf("pID=%s\n", pValue)
		}

		c.JSON(http.StatusOK, "OK")
	})
	//app.POST("/update", func(c *gin.Context) {
	//fmt.Println("receive a request.")
	//c.JSON(http.StatusOK, "OK")
	//})

	srv := &http.Server{
		Addr:    ":5353",
		Handler: app,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Println("error occurs.")
		}
	}()

	<-sigc
}
