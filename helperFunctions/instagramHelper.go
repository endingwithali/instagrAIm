package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type FBBusinessAccountInfo struct {
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

type IGBusinessAccountInfo struct {
	InstagramBusinessAccount struct {
		ID string `json:"id"`
	} `json:"instagram_business_account"`
	ID string `json:"id"`
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
	var fbBusinessAccountInfo FBBusinessAccountInfo
	marshallerr := json.Unmarshal(body, &fbBusinessAccountInfo)
	if marshallerr != nil {
		log.Panic(marshallerr)
	}
	businessId := fbBusinessAccountInfo.Data[0].ID
	log.Printf("creator id struct %s", fbBusinessAccountInfo)
	log.Printf("creator id pulled %s", businessId)
	resp.Body.Close()

	resp2, get2err := http.Get(fmt.Sprintf("https://graph.facebook.com/v15.0/%s?fields=instagram_business_account&access_token=%s", businessId, accessToken))

	if get2err != nil {
		log.Panic(get2err)
	}
	// // how to print out body example
	// dump, dumpyerr := httputil.DumpResponse(resp2, true)
	// if dumpyerr != nil {
	// 	log.Panic(dumpyerr)
	// }
	// log.Printf("Response Body: %s", dump)

	body2, body2readerr := io.ReadAll(resp2.Body)
	if body2readerr != nil {
		log.Panic(body2readerr)
	}
	log.Printf("respones body: %s", string(body2))

	//unpack json object
	var igBusnessAccountInfo IGBusinessAccountInfo
	igError := json.Unmarshal(body2, &igBusnessAccountInfo)
	if igError != nil {
		log.Panic(igError)
	}
	log.Printf("igInfo %v", igBusnessAccountInfo)
}
