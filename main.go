package main

import (
	"encoding/json"
	"os"
)

func main() {
	a := App{}
	a.Run(":8080")
}
