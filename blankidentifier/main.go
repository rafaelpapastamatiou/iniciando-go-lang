package blankidentifier

import (
	"fmt"
	"net/http"
)

func Describe() {
	fmt.Println("\n\nBlank identifier")
	res, _ := http.Get("http://google.com")
	//body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	//fmt.Printf("%s", body)
}
