package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type IgMeAccountStruct struct {
	Data []struct {
		AccessToken  string `json:"access_token"`
		Category     string `json:"category"`
		CategoryList []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"category_list"`
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
	} `json:"paging"`
}

func InstagramActivate() {
	accessToken := os.Getenv("INSTAGRAM_ACCESS_TOKEN")
	resp, geterr := http.Get(fmt.Sprintf("https://graph.facebook.com/v15.0/me/accounts?access_token=%s", accessToken))
	if geterr != nil {
		log.Panic(geterr)
	}
	// how to print out body example
	// dump, dumpyerr := httputil.DumpResponse(resp, true)
	// if dumpyerr != nil {
	// 	log.Panic(dumpyerr)
	// }
	// log.Printf("Response Body: %s", dump)

	body, bodyreaderr := io.ReadAll(resp.Body)
	if bodyreaderr != nil {
		log.Panic(bodyreaderr)
	}
	// log.Printf("respones body: %s", string(body))

	//unpack json object
	var igMeAccountInfo IgMeAccountStruct
	marshallerr := json.Unmarshal(body, &igMeAccountInfo)
	if marshallerr != nil {
		log.Panic(marshallerr)
	}
	log.Printf("creator id struct %s", igMeAccountInfo)
	log.Printf("creator id pulled %s", igMeAccountInfo.Data[0].ID)
	resp.Body.Close()

}
