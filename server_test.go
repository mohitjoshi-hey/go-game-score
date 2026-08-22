package gogamescore

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)

	got := response.Body.String()
	want := "20"

	if got != want{
		t.Errorf("got %q, want %q", got, want)
	}
}

func PlayerServer(w http.ResponseWriter, r *http.Request){
// 	w → where I write the response
// 	r → information about the incoming request
	fmt.Fprint(w, "20")
}