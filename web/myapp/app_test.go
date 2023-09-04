package myapp

import (
    "io/ioutil"
	"testing"
	"net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/assert"
    "encoding/json"
	"strconv"
    "strings"
)

func TestIndex(t *testing.T) {
    assert := assert.New(t)

    // NOTE: http mock up 서버를 구성합니다.
    ts := httptest.NewServer(NewHandler())
    defer ts.Close()

    resp, err := http.Get(ts.URL)
    assert.NoError(err)
    assert.Equal(http.StatusOK, resp.StatusCode)
    data, _ := ioutil.ReadAll(resp.Body)
    assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
    assert := assert.New(t)

    ts := httptest.NewServer(NewHandler())
    defer ts.Close()

    resp, err := http.Get(ts.URL + "/users")
    assert.NoError(err)
    assert.Equal(http.StatusOK, resp.StatusCode)
    data, _ := ioutil.ReadAll(resp.Body)
    assert.Contains(string(data), "Get UserInfo")
}

func TestGetUserInfo(t *testing.T) {
    assert := assert.New(t)

    ts := httptest.NewServer(NewHandler())
    defer ts.Close()

    resp, err := http.Get(ts.URL + "/users/42")
    assert.NoError(err)
    assert.Equal(http.StatusOK, resp.StatusCode)
    data, _ := ioutil.ReadAll(resp.Body)
    assert.Contains(string(data), "No User ID:42")
}

func TestCreateUserInfo(t *testing.T) {
    assert := assert.New(t)

    ts := httptest.NewServer(NewHandler())
    defer ts.Close()

    resp, err := http.Post(ts.URL + "/users", "application/json",
            strings.NewReader(`{"first_name": "yam", "last_name": "kim", "email": "dev.yamkim@gmail.com"}`))
    assert.NoError(err)
    assert.Equal(http.StatusCreated, resp.StatusCode)
    // 생성된 유저 확인
    createdUser := new(User)
    err = json.NewDecoder(resp.Body).Decode(createdUser)
    assert.NoError(err)
    assert.NotEqual(0, createdUser.ID)

    // 생성된 유저의 정보 조회 후 비교
    id := createdUser.ID
    resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
    assert.NoError(err)
    assert.Equal(http.StatusOK, resp.StatusCode)
    resUser := new(User)
    err = json.NewDecoder(resp.Body).Decode(resUser)
    assert.Equal(createdUser.ID, resUser.ID)
}

func TestDeleteUser(t *testing.T) {
    assert := assert.New(t)

    ts := httptest.NewServer(NewHandler())
    defer ts.Close()

    // NOTE: 없는 유저를 삭제하는 경우
    req, _ := http.NewRequest("DELETE", ts.URL + "/users/1", nil)
    resp, err := http.DefaultClient.Do(req)
    assert.NoError(err)
    assert.Equal(http.StatusOK, resp.StatusCode)
    data, _ := ioutil.ReadAll(resp.Body)
    assert.Contains(string(data), "No User ID:1")

    // NOTE: 유저 생성
    resp, err = http.Post(ts.URL + "/users", "application/json",
            strings.NewReader(`{"first_name": "yam", "last_name": "kim", "email": "dev.yamkim@gmail.com"}`))
    assert.NoError(err)
    assert.Equal(http.StatusCreated, resp.StatusCode)

    user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

    // NOTE: 생성한 유저를 삭제하는 경우
    req, _ = http.NewRequest("DELETE", ts.URL + "/users/" + strconv.Itoa(user.ID), nil)
    resp, err = http.DefaultClient.Do(req)
    assert.NoError(err)
    assert.Equal(http.StatusOK, resp.StatusCode)
    data, _ = ioutil.ReadAll(resp.Body)
    assert.Contains(string(data), "Deleted User ID:" + strconv.Itoa(user.ID))
}

func TestFooHandler_WithoutJson(t *testing.T) {
    assert := assert.New(t)

    res := httptest.NewRecorder()
    req := httptest.NewRequest("GET", "/foo", nil)

    mux := NewHttpHandler()
    mux.ServeHTTP(res, req)

    assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
    assert := assert.New(t)

    res := httptest.NewRecorder()
    req := httptest.NewRequest("POST", "/foo", 
        strings.NewReader(`{"first_name": "yam", "last_name": "kim", "email": "test@example.com"}`))

    mux := NewHttpHandler()
    mux.ServeHTTP(res, req)

    assert.Equal(http.StatusCreated, res.Code)

    user := new(User)
    err := json.NewDecoder(res.Body).Decode(user)
    assert.Nil(err)
    assert.Equal("yam", user.FirstName)
    assert.Equal("kim", user.LastName)
}