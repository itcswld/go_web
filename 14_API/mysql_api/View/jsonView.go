package jsonView

type Response struct{
	Code int 			`json:"Code"`
	Body interface{} 	`json:"Body"` //"interface{}" = anyType
}

type PostRequest struct{
	Name string	`json:"Name"`
	Task string `json:"Task"`
}
