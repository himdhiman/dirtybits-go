package main

import (
	"fmt"

	"github.com/himdhiman/dirtybits-golang/auth/internal/config"
)

func main() {
	cfg, err := config.GetConfiguration("/Users/himanshudhiman/code/golang/dirtybits/auth/settings/config.yaml")

	if err != nil {
		fmt.Errorf("Error occureed: %w", err)
	}

	fmt.Println(cfg)
}
