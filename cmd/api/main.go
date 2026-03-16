package main

import (
	"fmt"

	"github.com/Djuanzz/cashlens-backend/internal/router"
)

func main() {
	fmt.Println("===== Cashlens Backend =====")

	r := router.SetupRouter()
	r.Run(":5000")
}
