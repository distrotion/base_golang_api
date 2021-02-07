package main

import (
	"fmt"
	_"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
)


func main() {

	data := url.Values{
        "key":       {"Kuy_Parin"},
        "occupation": {"bigalow"},
    }

	
	resp, err := http.PostForm("https://us-central1-first-test-api-01.cloudfunctions.net/test/push", data)

    if err != nil {
        
    }

    var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    fmt.Println(res)


}
