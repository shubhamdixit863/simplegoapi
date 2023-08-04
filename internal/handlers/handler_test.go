package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"simplegoapi/internal/dto"
	"simplegoapi/internal/mocks"
	"testing"
)

func TestHandler_AddData(t *testing.T) {
	mockService := new(mocks.MockService)

	mockService.On("AddData", mock.Anything).Return()
	h := NewHandler(mockService)

	// We will create a temporary mock http server
	user := dto.AddRequest{
		Name: "",
		Age:  0,
	}
	///
	marshal, err := json.Marshal(user)
	assert.Nil(t, err) // if this condition is true our this test would pass
	request, err := http.NewRequest("POST", "/user", bytes.NewBuffer(marshal))
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	h.AddData(rr, request)

	t.Log(rr.Body)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestHandler_GetSingleData(t *testing.T) {
	mockService := new(mocks.MockService)

	mockService.On("GetSingleData", mock.Anything).Return(
		dto.AddRequest{
			Name: "Shubham",
			Age:  23,
		}, nil)
	h := NewHandler(mockService)

	request, err := http.NewRequest("GET", "/user/12", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	h.GetSingleData(rr, request)

	// Unmarshal the json response

	var data dto.AddRequest
	err = json.Unmarshal(rr.Body.Bytes(), &data)
	assert.Nil(t, err)
	t.Log(data.Name)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Shubham", data.Name)
}
