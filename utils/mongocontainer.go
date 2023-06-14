package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/testcontainers/testcontainers-go"
)

func LoadTestEnvs() error {
	keys := map[string]string{
		"MONGOURI":   "",
		"DBUSER":     "root",
		"DBPASS":     "password",
		"DBNAME":     "testdb",
		"DBCOLL":     "testcoll",
		"DBPORT":     "27017/tcp",
		"SERVERPORT": "8000",
	}

	for a := range keys {
		err := os.Setenv(a, keys[a])
		if err != nil {
			return err
		}
	}

	fmt.Println("Succesfully loaded envs")
	return nil
}

func CreateMongoContainer() (testcontainers.Container, error) {
	envs := map[string]string{
		"MONGO_INITDB_ROOT_USERNAME": os.Getenv("DBUSER"),
		"MONGO_INITDB_ROOT_PASSWORD": os.Getenv("DBPASS"),
		"MONGO_INITDB_DATABASE":      os.Getenv("DBNAME"),
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{},
		Env:          envs,
		AutoRemove:   true,
		Name:         "testMongoContainer",
	}

	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	endpoint, err := mongoC.Endpoint(ctx, "")
	err = os.Setenv("MONGOURI", fmt.Sprintf("mongodb://%s:%s@%s", user, pass, endpoint))
	if err != nil {
		return nil, err
	}

	fmt.Printf("Succesfully created container %s\n", mongoC.GetContainerID())
	return mongoC, err
}

func TerminateMongoContainer(ctx context.Context, c testcontainers.Container) error {
	id := c.GetContainerID()

	if err := c.Terminate(ctx); err != nil {
		return err
	}

	// envs are cleared to not interfere with other test containers
	os.Clearenv()
	fmt.Printf("Succesfully terminated container %s\n", id)
	return nil
}
