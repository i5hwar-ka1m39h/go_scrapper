package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization : bingo {actual api key}
// check for the above heades check for the actual api key and validate it

func GetApiKey(headers http.Header)(string,error){
	value := headers.Get("Authorization")
	if value == ""{
		return  "", errors.New("no header value is provide")
	}

	vals := strings.Split(value, " ")
	if len(vals) != 2{
		return "", errors.New("you fucked up space")
	}

	if vals[0] != "bingo"{
		return "", errors.New("you fucked up, bingo")
	}

	return vals[1], nil
}