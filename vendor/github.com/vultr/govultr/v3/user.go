package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const path = "/v2/users"

// UserService is the interface to interact with the user management endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/users
type UserService interface { //nolint:dupl
	Create(ctx context.Context, userCreate *UserReq) (*User, *http.Response, error)
	Get(ctx context.Context, userID string) (*User, *http.Response, error)
	Update(ctx context.Context, userID string, userReq *UserReq) error
	Delete(ctx context.Context, userID string) error
	List(ctx context.Context, options *ListOptions) ([]User, *Meta, *http.Response, error)
}

var _ UserService = &UserServiceHandler{}

// UserServiceHandler handles interaction with the user methods for the Vultr API
type UserServiceHandler struct {
	client *Client
}

// User represents an user on Vultr
type User struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	APIEnabled  *bool    `json:"api_enabled"`
	APIKey      string   `json:"api_key,omitempty"`
	ACL         []string `json:"acls,omitempty"`
	ServiceUser bool     `json:"service_user"`
}

// UserReq is the user struct for create and update calls
type UserReq struct {
	Email       string   `json:"email,omitempty"`
	Name        string   `json:"name,omitempty"`
	APIEnabled  *bool    `json:"api_enabled,omitempty"`
	ACL         []string `json:"acls,omitempty"`
	Password    string   `json:"password,omitempty"`
	ServiceUser bool     `json:"service_user,omitempty"`
}

type usersBase struct {
	Users []User `json:"users"`
	Meta  *Meta  `json:"meta"`
}

type userBase struct {
	User *User `json:"user"`
}

// Create will add the specified user to your Vultr account
func (u *UserServiceHandler) Create(ctx context.Context, userCreate *UserReq) (*User, *http.Response, error) {
	req, err := u.client.NewRequest(ctx, http.MethodPost, path, userCreate)
	if err != nil {
		return nil, nil, err
	}

	user := new(userBase)
	resp, err := u.client.DoWithContext(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user.User, resp, nil
}

// Get will retrieve a specific user account
func (u *UserServiceHandler) Get(ctx context.Context, userID string) (*User, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s", path, userID)

	req, err := u.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(userBase)
	resp, err := u.client.DoWithContext(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user.User, resp, nil
}

// Update will update the given user. Empty strings will be ignored.
func (u *UserServiceHandler) Update(ctx context.Context, userID string, userReq *UserReq) error {
	uri := fmt.Sprintf("%s/%s", path, userID)
	req, err := u.client.NewRequest(ctx, http.MethodPatch, uri, userReq)
	if err != nil {
		return err
	}

	_, err = u.client.DoWithContext(ctx, req, nil)
	return err
}

// Delete will remove the specified user from your Vultr account
func (u *UserServiceHandler) Delete(ctx context.Context, userID string) error {
	uri := fmt.Sprintf("%s/%s", path, userID)

	req, err := u.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = u.client.DoWithContext(ctx, req, nil)
	return err
}

// List will list all the users associated with your Vultr account
func (u *UserServiceHandler) List(ctx context.Context, options *ListOptions) ([]User, *Meta, *http.Response, error) { //nolint:dupl
	req, err := u.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	users := new(usersBase)
	resp, err := u.client.DoWithContext(ctx, req, &users)
	if err != nil {
		return nil, nil, resp, err
	}

	return users.Users, users.Meta, resp, nil
}
