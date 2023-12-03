package auth

import (
	"errors"
	"net/http"
	"strings"
)

//extract API key from http header
/*
Authorization:{
	ApiKey: {.........}
}

*/
func GetAPIKey(headers http.Header)(string , error){
	k:=headers.Get("Authorization")
	if k == ""{
		return "",errors.New("authentication fail : API key is missing")
	}
	vals:= strings.Split(k, " ")
	if len(vals)!=2 {
		return "",errors.New("authentication fail :API key is malformed")
	}
	if vals[0]!="ApiKey" {
		return "",errors.New("authentication fail :Malformed key")
	}
    return vals[1],nil

	
	
}