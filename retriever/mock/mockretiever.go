package mock

import "fmt"

/**
 * @Project GoLearnProject
 * @File    mockretiever.go
 * @Author  Augus Lee
 * @Date    2022/11/4 10:35
 * @GoVersion go1.19.2
 * @Version 1.0
 */

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf(
		"Retriever: {Contents=%s}", r.Contents)
}

func (r *Retriever) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
