package greeter

import (
	"fmt"

	"github.com/google/uuid"
)

// Hello returns a greeting for the given name with a unique ID.
func Hello(name string) string {
	id := uuid.New()
	return fmt.Sprintf("[%s] Hello, %s! Welcome to my module.", id.String(), name)
}
