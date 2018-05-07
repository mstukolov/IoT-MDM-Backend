package main

import (
	"github.com/minio/minio-go"
	"log"
)

func main() {
	endpoint := "http://127.0.0.1:9000"
	accessKeyID := "ZN4P2S95HGGX4RK1GK5I"
	secretAccessKey := "G0Fv17Z+xiBbtldD9io0WzhK1898MH89iZjvFlfm"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup
}