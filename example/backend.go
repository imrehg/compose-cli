package example

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/api/backend"
	"github.com/docker/api/containers"
)

type containerService struct{}

func init() {
	backend.Register("example", "example", func(ctx context.Context) (interface{}, error) {
		return New(), nil
	})
}

func New() containers.ContainerService {
	return &containerService{}
}

func (cs *containerService) List(ctx context.Context) ([]containers.Container, error) {
	return []containers.Container{
		{
			ID:    "id",
			Image: "nginx",
		},
		{
			ID:    "1234",
			Image: "alpine",
		},
	}, nil
}

func (cs *containerService) Run(ctx context.Context, r containers.ContainerConfig) error {
	fmt.Printf("Running container %q with name %q\n", r.Image, r.ID)
	return nil
}

func (cs *containerService) Exec(ctx context.Context, name string, command string, reader io.Reader, writer io.Writer) error {
	fmt.Printf("Executing command %q on container %q", command, name)
	return nil
}