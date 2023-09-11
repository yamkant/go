package app

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"

	"testing"

	"github.com/stretchr/testify/assert"
	"example.com/m/model"
)

func getTodoBodyData(resp *http.Response) (model.Todo, error) {
	var todo model.Todo
	err := json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func getTodoBodyDataList(resp *http.Response) ([]*model.Todo, error) {
	todos := []*model.Todo{}
	err := json.NewDecoder(resp.Body).Decode(&todos)
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func getDeleteAPIResponseData(resp *http.Response) (Success, error){
	var success Success
	err := json.NewDecoder(resp.Body).Decode(&success)
	if err != nil {
		return success, err
	}
	return success, nil
}

func TestTodos(t *testing.T) {
	// NOTE: Test mock up
	getSessionID = func (r *http.Request) string {
		return "testsessionId"
	}

	os.Remove("./test.db")
	assert := assert.New(t)

	ah := MakeNewHandler("./test.db")
	defer ah.Close()

	ts := httptest.NewServer(ah)
	defer ts.Close()

	// NOTE: CREATE TEST
	for i := 0; i < 2; i++ {
		todoName := "Test todo" + strconv.Itoa(i)
		resp, err := http.PostForm(ts.URL + "/todos", url.Values{"name": {todoName}})
		assert.NoError(err)
		assert.Equal(http.StatusCreated, resp.StatusCode)
		todo, err := getTodoBodyData(resp)
		assert.NoError(err)
		assert.Equal(todo.Name, todoName)
	}


	// NOTE: GET TEST
	resp, err := http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos, err := getTodoBodyDataList(resp)
	assert.NoError(err)
	assert.Equal(len(todos), 2)

	// NOTE: PATCH TEST
	patchData := []byte(`{"completed": true}`)
	req, _ := http.NewRequest("PATCH", ts.URL + "/todos/" + strconv.Itoa(todos[0].ID), bytes.NewBuffer(patchData))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos, err = getTodoBodyDataList(resp)
	for _, todo := range todos {
		if todo.ID == todos[0].ID {
			assert.True(todo.Completed)
		}
	}

	// NOTE: DELETE TEST
	req, _ = http.NewRequest("DELETE", ts.URL + "/todos/" + strconv.Itoa(todos[0].ID), nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	result, err := getDeleteAPIResponseData(resp)
	assert.True(result.Success)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos, err = getTodoBodyDataList(resp)
	assert.Equal(len(todos), 1)
}