package helper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func OpenAIGenerate() {
	client := &http.Client{}
	jsonBody := []byte(`{"prompt": "a white siamese cat", "n": 1,"size": "1024x1024"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bodyReader)
	if err != nil {
		log.Panic(err)
	}

	//Setting Up HTTP Headers
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPEN_AI_CODE")))
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("client: got response!\n")
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	log.Print(body)
	log.Printf("%q", dump)

	log.Printf("client: status code: %d\n", resp.StatusCode)
	resp.Body.Close()

}

// curl https://api.openai.com/v1/images/generations \
//   -H 'Content-Type: application/json' \
//   -H "Authorization: Bearer $OPENAI_API_KEY" \
//   -d '{
//     "prompt": "a white siamese cat",
//     "n": 1,
//     "size": "1024x1024"
//   }'
