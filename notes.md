# Notes

(Scheduled functions) [https://docs.netlify.com/functions/scheduled-functions/]
- 


- followed this documentation to create go structure https://docs.netlify.com/functions/build/?fn-language=go


- A Synchronous function is a function that does not return until the work is completed or has failed. 

# Meta API

- can publish up to 25 posts per day
- JPEG only 


must create a media container first before publishing 
https://developers.facebook.com/docs/instagram-api/reference/ig-user/media#create-container
https://stackoverflow.com/questions/65886724/instagram-graph-api-post-an-image

instagram authorization flow 
https://developers.facebook.com/docs/instagram-basic-display-api/guides/getting-access-tokens-and-permissions/


redirect uri -  a dedicated URI that can capture redirect query string parameters

testing on egirl__world instagram account -> pink stream profile 


how to set up instagram business account 
https://superface.ai/blog/instagram-setup



https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go


generate access tokens using https://developers.facebook.com/tools/accesstoken
- if youre getting this error 
- https://twitter.com/endingwithali/status/1614502025383014400




https://codepen.io/vcurd/details/RwaQPrb



getting user id from instagram api

https://developers.facebook.com/docs/instagram-api/getting-started/
- {"instagram_business_account":{"id":"17841435321631283"},"id":"100285556302931"}%
^^ using this tutorail in combination with your command line, and the facebook api graph explorer to generate your access tokens, you can create your first api call to activate your meta api and ask for permissions to post to your instagram account


instagram account 

https://stackoverflow.com/questions/15049903/how-to-use-custom-packages


unminify.com 


response for 
	resp, err := http.Get(fmt.Sprintf("https://graph.facebook.com/v15.0/me/accounts?access_token=%s", accessToken))

{
    "data": [
        {
            "access_token": "",
            "category": "Digital creator",
            "category_list": [
                { "id": "", "name": "Digital creator" }
            ],
            "name": "Egirlworld",
            "id": ""
        }
    ],
    "paging": { "cursors": { "before": "", "after": "" } }
}

body[data][0][id]


accessing json with golang https://dev.to/billylkc/parse-json-api-response-in-go-10ng
https://go.dev/blog/json