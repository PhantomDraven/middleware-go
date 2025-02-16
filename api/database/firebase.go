package database

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

// InitializeFirebase sets up the Firebase connection
func InitializeFirebase() {
	var app *firebase.App
	var err error

	ctx := context.Background()

	fmt.Printf("Database type cloud: %s \n", os.Getenv("ENABLE_SAAS_DB"))

	if os.Getenv("ENABLE_SAAS_DB") == "true" {
		// Use Firebase with service account
		opt := option.WithCredentialsFile("firebase/serviceAccountKey.json")
		app, err = firebase.NewApp(ctx, nil, opt)
	} else {
		// Use Firebase Emulator
		app, err = firebase.NewApp(ctx, nil)
	}

	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}

	FirebaseAuth, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth: %v", err)
	}

	log.Println("Firebase initialized successfully")
}
