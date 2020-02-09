package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"reflect"
	"bytes"
)


func main (){
	// Get 请求
	resp,err := http.Get("http://test-restful-api.gloryholiday.com/nightking/exchangeRate?providerName=%22%22&cid=iwoflyCOM&originalCode=USD&targetCode=QAR")
	if err != nil {
		fmt.Println("error")
	}	
	fmt.Println("resp:",resp)
	fmt.Println("resp.Body:",resp.Body)
	fmt.Println("resp type:",reflect.TypeOf(resp))
	defer resp.Body.Close()
	Body,err := ioutil.ReadAll(resp.Body)
	fmt.Println("readall Body:",string(Body))
	fmt.Println("Body type:",reflect.TypeOf(Body))

	// Post 请求
	body := `{
		    "originalCode":"USD",
			"targetCode":"CNY",
			"publish_timestamp": "2019-11-28T00:00:00Z"}`
	fmt.Println("body type:",reflect.TypeOf(body))

	urlPost := "http://test-restful-api.gloryholiday.com/currencyservice/getCurrency"
	res,err := http.Post(urlPost,"application/json;charset=utf-8",bytes.NewBuffer([]byte(body)))
	fmt.Println("res.Body:",res.Body)
	if err != nil {
		fmt.Println("ERR:",err.Error())
	}
	defer res.Body.Close()
	content,err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err2 ",err.Error())
	}
	fmt.Println(string(content))
	fmt.Println("content type:",reflect.TypeOf(content))
}	