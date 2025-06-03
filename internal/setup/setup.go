package setup

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
)

type Core struct {
	Env   Environment
	Debug bool
}

type Environment struct {
	InstanceUUID string
}

func (c *Core) Shutdown() {

}

func safeEnv(env string) (string, error) {
	// Lookup env variable, and panic if not present
	res, present := os.LookupEnv(env)
	if !present {
		return "", fmt.Errorf("missing environment variable %s", env)
	}
	return res, nil
}

func getEnv(env, fallback string) string {
	if value, ok := os.LookupEnv(env); ok {
		return value
	}
	return fallback
}

func CreateCore() (*Core, []error) {
	var errs []error

	// Grab ENV Variables
	INSTANCE_UUID := uuid.New().String()
	DEBUG, err := strconv.ParseBool(getEnv("DEBUG", "false"))
	if err != nil {
		errs = append(errs, err)
	}

	// Error on missing env variables
	if len(errs) != 0 {
		return nil, errs
	}

	return &Core{
		Debug: DEBUG,
		Env: Environment{
			InstanceUUID: INSTANCE_UUID,
		},
	}, nil
}
