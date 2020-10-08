package apiproxy

import (
	"github.com/vavar/go-elementary/trace"
	"github.com/pkg/errors"
	"encoding/xml"
	"encoding/json"
	"net/http"
	"fmt"
)

type User struct {
	XMLName	xml.Name	`xml:"Test_envelop"`
	Page       int64   `json:"page"`       
	PerPage    int64   `json:"per_page"`   
	Total      int64   `json:"total"`      
	TotalPages int64   `json:"total_pages"`
	Data       []Datum `json:"data"`       
	Ad         Ad      `json:"ad"`         
}

type Ad struct {
	Company string `json:"company"`
	URL     string `json:"url"`
	Text    string `json:"text"`   
}

type Datum struct {
	ID        int64  `json:"id" xml:"ID,attr"`        
	Email     string `json:"email"`     
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"` 
	Avatar    string `json:"avatar"`    
}

func List(client *http.Client, url string) (*User,error) {
	fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, errors.Wrap(err, "list user")
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "list user")
	}

	var users User
	err = json.NewDecoder(res.Body).Decode(&users)
	defer res.Body.Close()
	return &users, err
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/xml")
	logger := trace.UnWrap(r)
	logger.Info("hello world")
	user, err := List(&http.Client{}, "https://reqres.in" + r.RequestURI)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	xml.NewEncoder(w).Encode(&user)
}