package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // запрос на 5 ед. кафе

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code) //проверка HTTP status code
	assert.NotEmpty(t, responseRecorder.Body)              //проверка, что тело не пустое
	body := responseRecorder.Body.String()
	cafeList := strings.Split(body, ",")
	cafeListCount := len(cafeList)
	assert.Equal(t, totalCount, cafeListCount) // проверка вывода всех доступных кафе при забросе превышающем их количество

}

func TestMainHandlerWhenCityDoesntExist(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=vladimir", nil) // запрос c не существующим городом

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code) //проверка HTTP status code
	assert.NotEmpty(t, responseRecorder.Body)                     //проверка, что тело не пустое
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body) //проверка вывода сообщения о не правильном значении "город"
}
