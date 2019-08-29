// 这个例子程序实现了网页计数的一个小例子。
// 服务启动之后，通过浏览器http://localhost:8080/count链接访问可以得到访问数。

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
