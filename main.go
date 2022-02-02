package main

import (
	"fmt"
	"github.com/gmaschi/go-httpclient/gohttp"
	"log"
	"net/http"
	"time"
)

func createTestClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(5 * time.Second).
		DisableTimeouts(true).
		SetMaxIdleConnections(3).
		Build()
	return client
}

var testClient = createTestClient()

func getUrl() {
	url := "https://api.github.com"
	headers := make(http.Header)

	res, err := testClient.Get(url, headers)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
	fmt.Println(res.String())
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(firstName, lastName string) User {
	return User{firstName, lastName}
}

func postUserUrl(user User) {
	client := gohttp.NewBuilder().Build()
	url := "https://api.github.com"
	headers := make(http.Header)

	res, err := client.Post(url, user, headers)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.StatusCode)
	fmt.Println(res.Bytes())
	fmt.Println(res.String())
}

func main() {
	getUrl()
	//user := createUser("userA", "userB")
	//postUserUrl(user)
}
