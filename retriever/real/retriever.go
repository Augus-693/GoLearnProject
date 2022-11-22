package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

/**
 * @Project GoLearnProject
 * @File    retriever.go
 * @Author  Augus Lee
 * @Date    2022/11/4 10:49
 * @GoVersion go1.19.2
 * @Version 1.0
 */

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(
		resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
