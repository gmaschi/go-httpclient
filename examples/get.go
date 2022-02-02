package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func Get() (*Endpoints, error){
	url := "https://api.github.com"
	res, err := httpClient.Get(url, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status code: %d\n", res.StatusCode))
	fmt.Println(fmt.Sprintf("Status: %d\n", res.Status))
	fmt.Println(fmt.Sprintf("Body: %s\n", res.String()))

	endpoints := Endpoints{}

	err = res.UnmarshalJson(&endpoints)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repository url: %s\n", endpoints.RepositoryUrl))
	return &endpoints, nil
}
