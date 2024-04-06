package selectracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the faster one", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		// by prefixing a function call with defer
		// it will be called at the end of the containing function
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Did not expect an error but got one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("Returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := makeDelayedServer(25*time.Millisecond)

		defer server.Close()

		timeout := 20*time.Millisecond
		_, err := ConfigurableRacer(server.URL, server.URL, timeout)

		if err == nil{
			t.Error("expected an error but didnt get one")
		}

	})
}

// new server expects a function to handle it, we send in anon function
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
