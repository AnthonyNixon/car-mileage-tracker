package storage

import (
	"AnthonyNixon/car-mileage-tracker/cmd/models/storage_file"
	"AnthonyNixon/car-mileage-tracker/cmd/utls/httperr"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
)

var client *storage.Client
var ctx context.Context
var bucketname string
var bucket *storage.BucketHandle

func Initialize() {
	log.Print("Initializing Storage")
	ctx = context.Background()
	myClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize storage client: %s", err)
	}

	client = myClient

	bucketname = os.Getenv("CAR_MILEAGE_TRACKER_BUCKET_NAME")
	if bucketname == "" {
		log.Fatal("CAR_MILEAGE_TRACKER_BUCKET_NAME not set.")
	}

	bucket = client.Bucket(bucketname)
}

func SaveStorageFile(storageFile storage_file.StorageFile) (returnErr httperr.HttpErr) {
	objectName := fmt.Sprintf("%s.json", storageFile.CarId)
	wc := bucket.Object(objectName).NewWriter(ctx)
	wc.ContentType = "application/json"

	content, err := json.Marshal(storageFile)
	if err != nil {
		return httperr.New(http.StatusInternalServerError, "Failed to marshal json", err.Error())
	}

	if _, err := wc.Write(content); err != nil {
		return httperr.New(http.StatusInternalServerError, "failed to write content", err.Error())
	}

	if err := wc.Close(); err != nil {
		return httperr.New(http.StatusInternalServerError, "failed to close writer", err.Error())
	}

	return nil
}

func GetStorageFileByID(id string) (storageFile storage_file.StorageFile, returnErr httperr.HttpErr) {
	objectName := fmt.Sprintf("%s.json", id)

	obj := bucket.Object(objectName)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return storageFile, httperr.New(http.StatusNotFound, "Object does not exist", err.Error())
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return storageFile, httperr.New(http.StatusInternalServerError, "Failed to read object", err.Error())
	}

	err = json.Unmarshal(data, &storageFile)
	if err != nil {
		return storageFile, httperr.New(http.StatusInternalServerError, "Failed to read object", err.Error())
	}

	return
}

func SaveFile(content []byte, filename string) (returnErr httperr.HttpErr) {
	wc := bucket.Object(filename).NewWriter(ctx)
	wc.ContentType = "application/json"

	if _, err := wc.Write(content); err != nil {
		return httperr.New(http.StatusInternalServerError, "failed to write content", err.Error())
	}

	if err := wc.Close(); err != nil {
		return httperr.New(http.StatusInternalServerError, "failed to close writer", err.Error())
	}

	return nil
}
