package v1

import (
	"encoding/json"
	"fmt"
	"github.com/siskinc/todolist-go/routes/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestUserRegister_POST(t *testing.T) {
	httptest.NewRecorder()
	r := router.Router
	// 使用抓取到的处理程序执行一个 GET 请求。
	data, _ := json.Marshal(map[string]string{
		"username":    "123456789",
		"password":    "123456789",
		"re_password": "123456789",
		"email":       "123456789",
	})

	body := strings.NewReader(string(data))
	fmt.Println(body)
	req, _ := http.NewRequest("POST", "/v1/user/register", body)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	fmt.Println("req", req)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println("w", w)
}
