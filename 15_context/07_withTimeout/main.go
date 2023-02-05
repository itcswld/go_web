package main

/*context make it possible to manage a chain of calls within the same call path by signaling context's Done channel.
context 用来解决 goroutine 之间退出通知、元数据传递的功能。
主要用来在 goroutine 之间传递上下文信息，包括：取消信号、超时时间、截止时间、k-v 等。
在一组 goroutine 之间传递共享的值、取消信号、deadline……协调多个 groutine 中的代码执行“取消”操作，并且可以存储键值对。最重要的是它是并发安全的。
与它协作的 API 都可以由外部控制执行“取消”操作，例如：取消一个 HTTP 请求的执行。
https://zhuanlan.zhihu.com/p/68792989#:~:text=Go%201.7%20%E6%A0%87%E5%87%86%E5%BA%93%E5%BC%95%E5%85%A5,%E3%80%81%E6%88%AA%E6%AD%A2%E6%97%B6%E9%97%B4%E3%80%81k%2Dv%20%E7%AD%89%E3%80%82
*/
import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	// log.Println(ctx)
	ctx = context.WithValue(ctx, "userID", 777)
	result, err := ctxAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	fmt.Fprintln(w, result)
}

func ctxAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second )
	defer cancel()

	ch := make(chan int)
	go func ()  {
		id := ctx.Value("userID").(int)
		time.Sleep(1 * time.Second)
		//check to make sure we're not running in vain
		//if ctx.Done() has
		if ctx.Err() != nil {
			return
		}
		ch <- id
	}()
	select {
	//when > 1 * time.Second
	case <- ctx.Done():
		return 0, ctx.Err()
	//when < = 1 * time.Second
	case i := <- ch:
		return i, nil
	}
}

