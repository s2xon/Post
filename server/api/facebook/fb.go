package fb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"root/server/api"
)

type FacebookAccessToken struct {
	Access_Token	string	`json:"access_token"`
	Token_Type	string	`json:"token_type"`
	Expiration	int	`json:"expires_in"`
}

type FacebookMetaData struct {
	Access_Token	string	`json:"access_token"`
	Token_Type	string	`json:"token_type"`
	Expiration	int	`json:"expires_in"`
}

// Returns the User Access token for Facebook
func ACC_TKN(w http.ResponseWriter, r *http.Request) string{
	ur, err := url.Parse(r.URL.RequestURI())
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(ur.RawQuery)
	code := (m["code"][0])

	u, err := url.Parse("https://graph.facebook.com/v19.0/oauth/access_token?client_id={app-id}&redirect_uri={redirect-uri}&client_secret={app-secret}&code={code-parameter}")
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("client_id", keys.FB_ID)
	q.Set("client_secret", keys.FB_SECRET)
	q.Set("redirect_uri", keys.FB_URI)
	q.Set("code", code)

	u.RawQuery = q.Encode()
	url := u.String()

	res, err := http.Get(url)

	tkn := &FacebookAccessToken{}
	dc := json.NewDecoder(res.Body).Decode(tkn)
	if dc != nil {
		panic(tkn)
	}

	return tkn.Access_Token
}

func APP_TKN() string{
	u, err := url.Parse("https://graph.facebook.com/oauth/access_token?client_id={your-app-id}&client_secret={your-app-secret}&grant_type=client_credentials")
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("client_id", keys.FB_ID)
	q.Set("client_secret", keys.FB_SECRET)
	u.RawQuery = q.Encode()

	url := u.String()

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	tkn := &FacebookAccessToken{}
	dc := json.NewDecoder(res.Body).Decode(tkn)
	if dc != nil {
		panic(tkn)
	}

	return tkn.Access_Token
}


type Response struct {
	DATA struct {
		APP_ID		string		`json:"app_id"`
		TYPE		string		`json:"type"`
		APPLICATION string   `json:"application"`
		EXPIRES_AT	int64    `json:"expires_at"`
		IS_VALID	bool     `json:"is_valid"`
		ISSUED		int64    `json:"issued_at"`
		METADATA	struct {
			SSO	string `json:"sso"`
			} `json:"metadata"`
	
		SCOPES      []string `json:"scopes"`
		USER_ID      string   `json:"user_id"`
		} `json:"data"`
	}


func USER(acc_tkn string, app_tkn string) Response {
	
	u, err := url.Parse("https://graph.facebook.com/debug_token?input_token={token-to-inspect}&access_token={app-token-or-admin-token}")
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("input_token", acc_tkn)
	q.Set("access_token", app_tkn)
	u.RawQuery = q.Encode()

	url := u.String()
	// log.Println(url)
	
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	UserData := &Response{}
	dc := json.NewDecoder(res.Body).Decode(UserData)
	if dc != nil {
		panic(err)
	}

	log.Println(res.Body)

	return *UserData
}

type PageResponse struct {
    Data []struct {
        AccessToken  string `json:"access_token"`
        Category     string `json:"category"`
        CategoryList []struct {
            ID   string `json:"id"`
            Name string `json:"name"`
        } `json:"category_list"`
        Name  string   `json:"name"`
        ID    string   `json:"id"`
        Tasks []string `json:"tasks"`
    } `json:"data"`
}

	// {
	// 	"data": [
	// 	  {
	// 		"access_token": "page_access_token",
	// 		"category": "Internet Company",
	// 		"category_list": [
	// 		  {
	// 			"id": "2256",
	// 			"name": "Internet Company"
	// 		  }
	// 		],
	// 		"name": "Name of this Page",
	// 		"id": "page_id",
	// 		"tasks": [
	// 		  "ANALYZE",
	// 		  "ADVERTISE",
	// 		  "MODERATE",
	// 		  "CREATE_CONTENT"
	// 		]
	// 	  },

func PAGE_TKN(user_id string, user_tkn string) PageResponse {
	u, err := url.Parse("https://graph.facebook.com/v19.0/user_id/accounts?access_token=user_access_token")
	if err != nil {
		panic(err)
	}
	u.Path = fmt.Sprintf("v19.0/%s/accounts", user_id)
	q := u.Query()
	q.Set("access_token", user_tkn)
	u.RawQuery = q.Encode()

	url := u.String()

	log.Println(url)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	PageData := &PageResponse{}
	dc := json.NewDecoder(res.Body).Decode(PageData)
	if dc != nil {
		panic(err)
	}

	log.Println(res.Body)

	return *PageData

}

