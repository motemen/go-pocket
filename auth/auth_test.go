package auth_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/motemen/go-pocket/api"
	"github.com/motemen/go-pocket/auth"
	. "github.com/onsi/gomega"
)

func TestObtainRequestToken(t *testing.T) {
	RegisterTestingT(t)

	theCode := "4a334434-a4ac-38fa-a747-4049b4"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(`{"code":"%s"}`, theCode)))
	}))
	defer ts.Close()

	api.Origin = ts.URL

	res, err := auth.ObtainRequestToken("", "http://www.example.com/")

	Expect(err).To(BeNil())
	Expect(res.Code).To(Equal(theCode))
}
