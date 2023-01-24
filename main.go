package main

import (
	"fmt"
	"log"

	helper "github.com/endingwithali/instagrAIm/helperFunctions"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	fmt.Printf("in main \n")
	// helper.OpenAIGenerate()
	helper.InstagramActivate()
}
