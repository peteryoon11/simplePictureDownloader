package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	//router.POST("/getMyBook", getWebSourcTest)
	router.GET("/getMyBook", getWebSourcTest)
	//router.POST("/getUserInfo", getUser)
	//router.POST("/getUserInfo/:test", getUser)

	log.Fatal(http.ListenAndServe(":8090", router))
}
func getWebSourcTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//fmt.Println(ps.ByName("test"))
	//mem := Member{"Alex", 10, true}

	//test, err := ioutil.ReadAll(r.Body)

	//	var testmem Member
	//var testAuth structModule.ValAPIKey

	//	err = json.Unmarshal(test, &testAuth)

	/* 	var respondUser structModule.Response

	   	var tempBookArray []structModule.EBookInfo

	   	fmt.Println(testAuth)

	   	if validationModule.CheckAPIToken(testAuth) {
	   		tempBookArray, err = dbConnectModule.GetMyOwnBook(testAuth)
	   		respondUser = structModule.Response{200, "Respond Success", tempBookArray}
	   	} else {

	   		respondUser = structModule.Response{403, "Invalid Auth key or Expire key", tempBookArray}
	   	} */
	respondHTML := "<!DOCTYPE html>"
	respondHTML += "<html>"
	respondHTML += "<head> <meta charset='utf-8'> <meta name='viewport' content='width=device-width'> <title>JS Bin</title>  </head>"
	respondHTML += "<body>"
	respondHTML += "<img src='https://t1.daumcdn.net/cfile/tistory/99C4233C5BE7E1FE05' style='cursor: pointer;max-width:100%;height:auto'  width='820' height='1100'  filemime='image/jpeg' />"
	respondHTML += "</body>"
	respondHTML += "</html>"
	//temp, err := json.Marshal(respondHTML)
	/* 	if err != nil {
		fmt.Println(err)
	} */
	//fmt.Println()
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write([]byte(respondHTML))

}
