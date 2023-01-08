package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
)

var bucketName = os.Getenv("OSRM_BUILDER_VERSION")

func main() {
	logrus.Infof("fetching bucket %v objects name...", bucketName)
	objectFiles, err := GetBucketObjectNames(bucketName)
	if err != nil {
		logrus.Fatalf("could not get bucket object names: %v", err)
	}

	for _, fileName := range objectFiles {
		f, err := os.Create("/osrm-data/" + fileName)
		if err != nil {
			logrus.Fatalf("could not create file: %v", err)
		}
		logrus.Infof("downloading %v", fileName)
		fileBytes, err := downloadFile(bucketName, fileName)
		if err != nil {
			logrus.Fatalf("could not get bucket object names: %v", err)
		}
		_, err = f.Write(fileBytes)
		if err != nil {
			logrus.Fatalf("could not write file: %v", err)
		}
		logrus.Infof("created %v", fileName)
		f.Close()
	}
}

func GetBucketObjectNames(bucket string) ([]string, error) {
	objectNames := []string{}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Bucket(%q).Objects: %v", bucket, err)
		}
		objectNames = append(objectNames, attrs.Name)
	}

	return objectNames, nil
}

func downloadFile(bucket, object string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Minute*10)
	defer cancel()

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %v", object, err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	return data, nil
}
