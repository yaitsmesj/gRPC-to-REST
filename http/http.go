package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/yaitsmesj/gRPC-to-REST/customerror"
	t "github.com/yaitsmesj/gRPC-to-REST/types"
)

const url = "https://reqres.in/api"

type methodType string

const (
	get    methodType = "GET"
	post   methodType = "POST"
	put    methodType = "PUT"
	delete methodType = "DELETE"
)

// GetUser ...
func GetUser(id int32) (*t.UserJSON, error) {

	rURL := fmt.Sprintf("%s/users/%d", url, id)
	body, err := sendReq(rURL, nil, get)
	if err != nil {
		return nil, err
	}
	var r t.UserJSON
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, errors.Wrap(err, "Error while parsing body into User")
	}
	return &r, nil
}

// GetUserList ...
func GetUserList(page int32) (*t.UserListJSON, error) {
	rURL := fmt.Sprintf("%s/users?page=%d", url, page)
	body, err := sendReq(rURL, nil, get)
	if err != nil {
		return nil, err
	}
	var r t.UserListJSON
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, errors.Wrap(err, "Error while parsing body into User List")
	}
	return &r, nil
}

// Create ...
func Create(name string, job string) (*t.CreateJSON, error) {
	rURL := fmt.Sprintf("%s/users", url)
	jbody, _ := json.Marshal(map[string]string{
		"name": name,
		"job":  job,
	})
	pbody := bytes.NewBuffer(jbody)
	body, err := sendReq(rURL, pbody, post)
	if err != nil {
		return nil, err
	}
	var r t.CreateJSON
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, errors.Wrap(err, "Error while parsing body into User List")
	}
	return &r, nil
}

// Update ...
func Update(id int32, name string, job string) (*t.UpdateJSON, error) {
	rURL := fmt.Sprintf("%s/users/%d", url, id)
	jbody, _ := json.Marshal(map[string]string{
		"name": name,
		"job":  job,
	})
	pbody := bytes.NewBuffer(jbody)
	body, err := sendReq(rURL, pbody, put)
	if err != nil {
		return nil, err
	}
	var r t.UpdateJSON
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, errors.Wrap(err, "Error while parsing body into User List")
	}
	return &r, nil
}

// Delete ...
func Delete(id int32) error {
	rURL := fmt.Sprintf("%s/users/%d", url, id)
	_, err := sendReq(rURL, nil, delete)
	if err != nil {
		return err
	}
	return nil
}

func sendReq(rURL string, pBody *bytes.Buffer, method methodType) ([]byte, error) {
	var resp *http.Response
	var req *http.Request
	var err error
	switch method {
	case get:
		resp, err = http.Get(rURL)
	case post:
		resp, err = http.Post(rURL, "application/json", pBody)
	case put:
		req, err = http.NewRequest(http.MethodPut, rURL, pBody)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err = http.DefaultClient.Do(req)
	case delete:
		req, err = http.NewRequest(http.MethodDelete, rURL, nil)
		resp, err = http.DefaultClient.Do(req)
	}

	if err != nil {
		return nil, errors.Wrap(err, "Error while executing HTTP request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		return nil, &customerror.StatusCodeError{StatusCode: resp.StatusCode, Body: body}
	}
	return body, nil
}
