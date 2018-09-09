package nuagex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// User represents NuageX user
type User struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Token    string
}

// LoadCredentials : load user credentials from YAML file
func (u *User) LoadCredentials(fn string) *User {
	fmt.Printf("Loading user credentials from '%s' file\n", fn)
	yamlFile, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Printf("LoadCredentials error   #%v ", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, u)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return u
}

// LoginResponse : represents the credentials data reported back from API server
type LoginResponse struct {
	Token     string `json:"accessToken"`
	User      string `json:"user"`
	ExpiresIn int    `json:"expiresIn"`
}

// Login logs in a user with the passed in `login User` struct
func (u *User) Login() (*User, error) {
	fmt.Printf("Logging '%s' user in...\n", u.Username)
	body, _ := json.Marshal(u)
	URL := buildURL("/auth/login")
	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, _ := client.Do(req)

	if response.StatusCode != 200 {
		fmt.Printf("Failed to login a user! Aborting...\n")
		os.Exit(1)
	}
	defer response.Body.Close()

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Printf("User '%s' logged in...\n", u.Username)
	var loginResponse LoginResponse

	json.Unmarshal(body, &loginResponse)

	u.Token = loginResponse.Token

	return u, nil
}
