package httpHandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wms.api/model"

	sqlsvr "wms.api/Controller/db"
)


func CRUD(dbName string, tb string) http.HandlerFunc{
	var db = sqlsvr.ConnRemoteMssql(dbName)
	switch tb {
		case "stk" :
			return func(w http.ResponseWriter,r *http.Request){
				c := make(chan []model.STK)
				//Get
				if r.Method == http.MethodGet {
					//read-all
					pn := r.URL.Query().Get("pn")
					if pn == "" {
						tsql := fmt.Sprintf(`SELECT pro_no,made_dt,Exp_dt FROM %s;`,tb)
						// defer db.Close()
						sqlsvr.Read(db,tsql)
						rows := <- sqlsvr.C
						defer rows.Close()

						var d []model.STK
						go func(){
							for rows.Next(){
								rd := model.STK{}
								rows.Scan(&rd.Pro_no, &rd.Made_dt, &rd.Exp_dt)
								d = append(d,rd)
							}
							c <- d
						}()
						w.WriteHeader(http.StatusOK)
						model.IfErr(json.NewEncoder(w).Encode(<-c),"httpHandler:")
					}else{
						tsql := fmt.Sprintf(`SELECT pro_no,made_dt,Exp_dt FROM %s where pro_no='%s';`,tb,pn)
						fmt.Println(tsql)
						// defer db.Close()
						sqlsvr.Read(db,tsql)
						rows := <- sqlsvr.C
						defer rows.Close()

						var d []model.STK
						go func(){
							for rows.Next(){
								rd := model.STK{}
								rows.Scan(&rd.Pro_no, &rd.Made_dt, &rd.Exp_dt)
								d = append(d,rd)
							}
							c <- d
						}()
						w.WriteHeader(http.StatusOK)
						model.IfErr(json.NewEncoder(w).Encode(<-c),"httpHandler:")
					}
				//update
				}else if r.Method == http.MethodPost{
					stk := model.STK{}
					json.NewDecoder(r.Body).Decode(&stk)
					tsql := fmt.Sprintf(`update stk set made_dt='%s' where pro_no='%s';`,stk.Made_dt,stk.Pro_no)
					sqlsvr.Update(db,tsql)
					w.WriteHeader(http.StatusCreated)
					json.NewEncoder(w).Encode(stk)
				}
			}
		default:
			return func(w http.ResponseWriter,r *http.Request){
					w.WriteHeader(http.StatusOK)
			}
	}
}
