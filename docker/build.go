package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/docker/docker/pkg/archive"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/mitchellh/go-homedir"
)

func testImageBuild() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err, " :unable to init client")
	}

	var wg sync.WaitGroup
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			dockerFile := fmt.Sprintf("%d.Dockerfile", i)
			imageBuildResponse, err := cli.ImageBuild(ctx, GetContext("."),
				types.ImageBuildOptions{
					Dockerfile: dockerFile,
					Tags:       []string{fmt.Sprintf("devbuild_%d", i)},
				})
			if err != nil {
				log.Fatal(err, " :unable to build docker image"+string(1))
			}
			defer imageBuildResponse.Body.Close()
			_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
			if err != nil {
				log.Fatal(err, " :unable to read image build response "+string(1))
			}
			log.Printf("Docker %d is built", i)
		}(i)
	}
	wg.Wait()
}

func GetContext(filePath string) io.Reader {
	filePath, _ = homedir.Expand(filePath)
	ctx, err := archive.TarWithOptions(filePath, &archive.TarOptions{})
	if err != nil {
		panic(err)
	}

	return ctx
}

func main() {
	testImageBuild()
}
