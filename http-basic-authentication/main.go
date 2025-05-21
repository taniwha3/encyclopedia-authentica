package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func BasicAuth(handler http.HandlerFunc, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We get the "Authentication" HTTP Header, which may contain the credential.
		auth := r.Header.Get("Authorization")

		// If the header is empty or is incorrectly formatted, we fail the authorization.
		// We expect the format to be:
		// Authorization: Basic <credential>
		// where <credential> is a Base64 encoded secret.
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"Restricted\"")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		// IMPORTANT: HTTP Basic Authentication is defined in RFC 7617.
		// The RFC explicitly calls for "standard base64" encoding.
		encodedCredentials := strings.TrimPrefix(auth, "Basic ")
		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization format"))
			return
		}

		// The credential is formatted "username:password", so we split
		// it. However, we must be cautious since the password itself
		// may have ":" characters
		credentials := strings.SplitN(string(decodedCredentials), ":", 2)
		if len(credentials) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid authorization format"))
			return
		}
		if username != credentials[0] || password != credentials[1] {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid credentials"))
			return
		}
		handler(w, r)
	}
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You successfully logged in"))
}

func main() {
	// Load up the handler and run it.
	// Test it with curl.
	// curl -u admin:secret http://localhost:8080/protected
	http.HandleFunc("/protected", BasicAuth(protectedHandler, "admin", "secret"))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
