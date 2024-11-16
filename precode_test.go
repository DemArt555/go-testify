package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	//totalCount := len(cafeList)
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // запрос на 5 ед. кафе

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code) // проверка HTTP status code
	body := responseRecorder.Body.String()
	cafeList := strings.Split(body, ",")
	cafeListCount := len(cafeList)
	assert.Len(t, cafeList, cafeListCount) // проверка вывода всех доступных кафе при запросе превышающем их количество

}

func TestMainHandlerWhenCityDoesntExist(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=5&city=vladimir", nil) // запрос c не существующим городом

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code) // проверка HTTP status code
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body) //проверка вывода сообщения о не правильном значении "город"
}

func TestQueryCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // запрос на 5 ед. кафе

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code) //проверка HTTP status code
	assert.NotEmpty(t, responseRecorder.Body)              //проверка, что тело не пустое

}
