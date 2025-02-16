package database

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

var (
	FirebaseAuth *auth.Client
	FirestoreDB  *db.Client
)

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
		// Use Firestore Emulator
		app, err = firebase.NewApp(ctx, nil)
	}

	if err != nil {
		log.Fatalf("Error initializing Firebase %v", err)
	}

	FirebaseAuth, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth %v", err)
	}

	FirestoreDB, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore %v", err)
	}

	log.Printf("Firebase initialized successfully")
}
