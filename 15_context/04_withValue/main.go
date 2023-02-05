package main

//per request variables
//good candidate for putting into context
import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Eve")
	log.Println(ctx)
	result := dbAccess(ctx)
	log.Println("userID:",result)
	fmt.Fprintln(w, result)
}
func dbAccess(ctx context.Context) int {
	id := ctx.Value("userID").(int)
	return id
}
