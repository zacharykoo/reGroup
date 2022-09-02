package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/zacharykoo/reGroup/backend/pkg/model"
	"github.com/zacharykoo/reGroup/backend/pkg/repository"
)

const (
	KeyUserID = "userID"
)

type user struct {
	repo repository.UserRepository
}

func GetUserService(repo repository.UserRepository) UserService {
	return &user{
		repo: repo,
	}
}

func (c *user) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userID int
		var err error
		userIDString := r.URL.Query().Get(KeyUserID)
		if userIDString == "" {
			userID = 0
		} else {
			userID, err = strconv.Atoi(userIDString)
			if err != nil {
				fmt.Printf("unable to parse userID: %v", err)
				w.Write([]byte("error getting user"))
				return
			}
		}
		someUser, err := c.repo.Get(userID)
		if err != nil {
			fmt.Printf("unable to get user: %v", err)
			w.Write([]byte("error getting user"))
		}
		b, err := json.Marshal(someUser)
		if err != nil {
			fmt.Printf("unable to marshal user: %v", err)
			w.Write([]byte("error getting user"))
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(b)
		if err != nil {
			fmt.Printf("Unable to send marshal user: %v", err)
			w.Write([]byte("error cannot send user"))
		}
	}
}

func (c *user) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var user model.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			fmt.Printf("unable to unmarshal into user: %v", err)
			return
		}

		user, _ = c.repo.Create(user)
		w.Write([]byte(fmt.Sprintf("created user with userID: %v", user.UserID)))
	}
}

func (c *user) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("unable to read body: %v", err)
			return
		}
		var user model.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			fmt.Printf("unable to unmarshal into user: %v", err)
			return
		}

		user, _ = c.repo.Edit(user)
		w.Write([]byte(fmt.Sprintf("user with userID %v is editted", user.UserID)))
	}
}
