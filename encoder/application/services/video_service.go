package services

import (
	"encoder/application/repositories"
	"encoder/domain"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"context"

	"cloud.google.com/go/storage"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(v.Video.FilePath)

	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}

	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	f, err := os.Create(os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	defer f.Close()

	log.Printf("O video %v has been stored", v.Video.ID)

	return nil
}

func (v *VideoService) Fragment() error {
	videoFilePath := os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID

	err := os.Mkdir(videoFilePath, os.ModePerm)
	if err != nil {
		return err
	}

	source := videoFilePath + ".mp4"
	target := videoFilePath + ".frag"

	cmd := exec.Command("mp4Fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	PrintOutput(output)

	return nil
}

func PrintOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("=====> Output: %s\n", string(out))
	}
}
