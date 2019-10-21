package main

import (
	"go-crawler/frontend/controller"
	"net/http"
)

func main() {
	//文章展示的功能，只要不是search，就拿目录文件
	http.Handle("/", http.FileServer(http.Dir("/Users/xue.zeng/XueZeng/github/cumulative/demo/Go/src/go-crawler/frontend/view")))

	http.Handle("/search", controller.CreateSearchResultHandler("/Users/xue.zeng/XueZeng/github/cumulative/demo/Go/src/go-crawler/frontend/view/template.html"))

	//http.Handle("/search",controller.SearchResultHandler{})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
