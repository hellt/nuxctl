package nux

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

const url string = "https://experience.nuagenetworks.net/api"

// User : NuageX user credentials
type User struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
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

// FlavorResponse : Image Flavor response JSON object mapping
type FlavorResponse struct {
	ID                   string        `json:"_id"`
	Name                 string        `json:"name"`
	CPU                  int           `json:"cpu"`
	Memory               int           `json:"memory"`
	V                    int           `json:"__v"`
	AllowedGroups        []string      `json:"allowedGroups"`
	AllowedOrganizations []interface{} `json:"allowedOrganizations"`
	Disk                 int           `json:"disk"`
	Created              time.Time     `json:"created"`
}

// LabResponse : NuageX Lab response JSON object mapping
type LabResponse struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   string `json:"status"`
}

func buildURL(path string) string {
	return url + path
}

// UserLogin : logs in a user with the credentials against NuageX API
func UserLogin(login User) (string, error) {
	fmt.Printf("Logging '%s' user in...\n", login.Username)
	body, _ := json.Marshal(login)
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
	fmt.Printf("User '%s' logged in...\n", login.Username)
	var result LoginResponse

	json.Unmarshal(body, &result)

	return result.Token, nil
}

// SendHTTPRequest : Request for a given method and url along with body parameters.
func SendHTTPRequest(method, url, token string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}
	return body, nil
}

// GetAllFlavors : Get all image flavors
func GetAllFlavors(token string) ([]FlavorResponse, error) {
	URL := buildURL("/flavors")
	b, err := SendHTTPRequest("GET", URL, token, nil)

	if err != nil {
		return []FlavorResponse{}, err
	}

	var result []FlavorResponse

	json.Unmarshal(b, &result)

	return result, nil
}

// CreateLab : Create a Lab in NuageX
func CreateLab(token string, reqb []byte) (LabResponse, error) {
	URL := buildURL("/labs")
	b, err := SendHTTPRequest("POST", URL, token, reqb)
	if err != nil {
		return LabResponse{}, err
	}
	var result LabResponse
	json.Unmarshal(b, &result)
	return result, nil
}
