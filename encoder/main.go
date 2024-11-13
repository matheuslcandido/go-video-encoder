package main

import (
	"encoder/domain"
	"fmt"
	"time"
)

func main() {
	video := domain.NewVideo()

	video.ID = "test"
	video.FilePath = "/aws/s3/test"
	video.ResourceId = "external_id"
	video.CreatedAt = time.Now()

	err := video.Validate()

	fmt.Println(err)
}