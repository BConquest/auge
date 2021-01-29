package lib

import (
	"net/http"
)

func CheckLinkResponse(url string) bool {
	//Assume if error happens that link is not valid
	resp, _ := http.Get(url)

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true
	}
	return false
}
