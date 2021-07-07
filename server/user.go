package main

import (
	"encoding/json"
	"fmt"
	user "gRPC/proto/gen"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"
)

type Server struct {
	user.UnimplementedUserServiceServer
}

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Followers int64  `json:"followers"`
	Gists     int64  `json:"public_gists"`
	URL       string `json:"url"`
}

//GetUser is get user on github
func (s *Server) GetUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	log.Printf("Receive message from client: %s", in.Username)

	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", in.Username))
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	usr := User{}
	jsonErr := json.Unmarshal(body, &usr)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &user.UserResponse{
		Id:        usr.ID,
		Name:      usr.Name,
		Avatarurl: usr.AvatarURL,
		Statistics: &user.Statistics{
			Followers: usr.Followers,
			Gists:     usr.Gists,
		},
		ListURLs: []string{usr.URL},
	}, nil
}