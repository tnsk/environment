package environment

import (
	"fmt"
	"os"
)

// Checks checks if all the environment variables are present
//
// if not present returns an error
//
// else returns nil
//
// sample usage:
//
//	if err := environment.Checks([]string{"PATH"}); err != nil {
//	 log.Fatal(err)
//	}
func Checks(envName []string) error {
	for _, e := range envName {
		if os.Getenv(e) == "" {
			return fmt.Errorf("missing enviroment variable %s", e)
		}
	}
	return nil
}
