package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/oauth2/google"
)

func main() {
	var email, credentialsFile, scopeString string
	flag.StringVar(&email, "email", "", "Superadmin email")
	flag.StringVar(&credentialsFile, "file", "", "Path to the credentials.json file")
	flag.StringVar(&scopeString, "scopes", "", "Comma-separated list of scopes")
	flag.Parse()

	if email == "" || credentialsFile == "" || scopeString == "" {
		log.Fatal("Please provide email, credentials file, and scopes using -email, -file, and -scopes flags.")
	}

	scopes := strings.Split(scopeString, ",")

	// Read the credentials file
	credentialsJSON, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Error reading credentials file: %v", err)
	}

	// Parse the JSON credentials to a Config object
	conf, err := google.JWTConfigFromJSON(credentialsJSON, scopes...)
	if err != nil {
		log.Fatalf("Error parsing credentials: %v", err)
	}

	// Impersonate the superadmin email if needed
	conf.Subject = email

	// You can use the 'conf' to create an HTTP client with the OAuth2 token
	// For this example, we'll just print the token.
	tokenSource := conf.TokenSource(context.Background())
	token, err := tokenSource.Token()
	if err != nil {
		log.Fatalf("Error fetching token: %v", err)
	}
	fmt.Printf(token.AccessToken)
}

