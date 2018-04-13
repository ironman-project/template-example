package controller

import "net/http"

{{$controller := (print .Values.name "Controller")}}

type  {{$controller}} struct {

}

func (c *{{$controller}}) Handle(w http.ResponseWriter, r *http.Request) {
	//Handle request
}
