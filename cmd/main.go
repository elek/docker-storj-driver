package main

import (
	"context"
	"fmt"
	"github.com/distribution/distribution/v3/reference"
	"github.com/distribution/distribution/v3/registry/storage"
	storj "github.com/storj/docker-storj-driver"
	"github.com/zeebo/errs"
	"log"
	"os"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("%++v", err)
	}

}

type Name struct {
	Value string
}

func (n Name) String() string {
	return n.Value
}

func (n Name) Name() string {
	return n.Value
}

var _ reference.Named = Name{}

func run() error {
	p := map[string]interface{}{}
	p["accessgrant"] = os.Getenv("UPLINK_ACCESS")
	p["bucket"] = "bucket1"

	driver, err := storj.FromParameters(p)
	if err != nil {
		return err
	}

	ctx := context.Background()
	registry, err := storage.NewRegistry(ctx, driver)
	if err != nil {
		return err
	}

	repository, err := registry.Repository(ctx, Name{"alpine"})
	if err != nil {
		return err
	}

	tags, err := repository.Tags(ctx).All(ctx)
	if err != nil {
		return errs.Wrap(err)
	}

	for _, tag := range tags {
		fmt.Println(tag)
	}
	return nil
}
