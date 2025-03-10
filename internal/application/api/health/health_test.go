package health

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const checkMark = "\u2713"

// const ballotX = "\u2717"

func TestHealth(t *testing.T) {

	t.Log("Test the /health route.")
	{
		req, err := http.NewRequest(http.MethodGet, HealthPath, nil)
		if err != nil {
			t.Errorf("Should be able to create a request: %v", err)
		}
		rw := httptest.NewRecorder()

		h := NewHealth()
		h.Handle()(rw, req)

		res := rw.Result()
		resBody, _ := io.ReadAll(res.Body)

		assert.Equal(t, http.StatusOK, rw.Code)
		assert.JSONEq(t, `{"message": "OK"}`, string(resBody))

		t.Log("Should pass", checkMark)

	}
}
