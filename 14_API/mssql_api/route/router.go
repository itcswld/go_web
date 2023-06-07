package route

import (
	"github.com/gorilla/mux"
	httpHandler "wms.api/Controller/httpHandler"
)
func Init() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/api/stk",httpHandler.CRUD("SML","stk"))
	return r
}
