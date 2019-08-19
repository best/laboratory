/*
 * @Time    : 2019/8/19 9:53
 * @Author  : XThundering
 * @File    : panjf2000-ants.go
 * @Software: GoLand
 */

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/panjf2000/ants"
)

type Request struct {
	Param  []byte
	Result chan []byte
}

func main() {
	go func() {
		for {
			log.Println(runtime.NumGoroutine())
			time.Sleep(time.Millisecond)
		}
	}()

	pool, _ := ants.NewPoolWithFunc(1000, func(payload interface{}) {
		request, ok := payload.(*Request)
		if !ok {
			return
		}
		reverseParam := func(s []byte) []byte {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}

			time.Sleep(time.Second)
			return s
		}(request.Param)

		request.Result <- reverseParam
	})
	defer pool.Release()

	http.HandleFunc("/reverse/ants", func(w http.ResponseWriter, r *http.Request) {
		param, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "request error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		request := &Request{Param: param, Result: make(chan []byte)}

		// Throttle the requests traffic with ants pool. This process is asynchronous and
		// you can receive a result from the channel defined outside.
		if err := pool.Invoke(request); err != nil {
			http.Error(w, "throttle limit error", http.StatusInternalServerError)
		}

		w.Write(<-request.Result)
	})

	http.ListenAndServe(":8081", nil)
}
