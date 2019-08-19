/*
 * @Time    : 2019/8/19 10:21
 * @Author  : XThundering
 * @File    : gin.go
 * @Software: GoLand
 */

package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		for {
			log.Println(runtime.NumGoroutine())
			time.Sleep(time.Millisecond)
		}
	}()

	router := gin.New()
	router.GET("/reverse/gin", func(context *gin.Context) {
		param, err := context.GetRawData()
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
		}

		reverseParam := func(s []byte) []byte {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}

			time.Sleep(time.Second)
			return s
		}(param)

		context.JSON(http.StatusOK, reverseParam)
	})
	_ = router.Run()
}
