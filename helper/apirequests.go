package helpers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	jsonBody := []byte(`{"prompt": "a white siamese cat", "n": 1,"size": "1024x1024"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bodyReader)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPEN_AI_CODE")))
	log.Print(req)
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	log.Print(resp)

}

// curl https://api.openai.com/v1/images/generations \
//   -H 'Content-Type: application/json' \
//   -H "Authorization: Bearer $OPENAI_API_KEY" \
//   -d '{
//     "prompt": "a white siamese cat",
//     "n": 1,
//     "size": "1024x1024"
//   }'
