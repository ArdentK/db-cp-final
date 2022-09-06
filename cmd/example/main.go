package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ArdentK/db-cp-final/config"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	usernamePtr := flag.String("username", "test_user6", "username")
	passwordPtr := flag.String("pass", "qwerty", "password")
	rolePtr := flag.String("user", "user", "admin")

	flag.Parse()

	if err := createUser(*usernamePtr, *passwordPtr, *rolePtr); err != nil {
		log.Fatal(err)
	}

	resp, err := authorize(*usernamePtr, *passwordPtr, *rolePtr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)

	// ctx := context.WithValue(context.Background(),
	// 	"Authorization", "Bearer"+resp.Token)

	res, err := http.Get("http://localhost:8080/competitions")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", res)

	// user, err := parser.ParseToken(token, []byte(viper.GetString("auth.signing_key")))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Successfully created and authorized user: %+v", user)

}

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func createUser(username, password, role string) error {
	reqBody := &userRequest{
		Username: username,
		Password: password,
		Role:     role,
	}

	resp := new(response)
	return request(reqBody, resp, "http://localhost:8080/auth/sign-up")
}

func authorize(username, password, role string) (*response, error) {
	reqBody := &userRequest{
		Username: username,
		Password: password,
		Role:     role,
	}

	resp := new(response)
	if err := request(reqBody, resp, "http://localhost:8080/auth/sign-in"); err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

func request(req *userRequest, res *response, endpoint string) error {
	reqBodyBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		endpoint,
		"application/json",
		bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, res); err != nil {
		return err
	}

	if res.Status == "error" {
		return errors.New(fmt.Sprintf("error occured on user creation: %s", res.Message))
	}

	return nil
}