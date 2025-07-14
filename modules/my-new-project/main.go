package main

import (
	"fmt"

	"github.com/your-username/my-new-project/greeter"
)

func main() {
	message := greeter.Hello("Go Developer")
	fmt.Println(message)
}
