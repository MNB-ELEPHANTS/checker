package checker

import (
	"context"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func (c *Checker) UpdateList(ctx context.Context) {
	buf, err := os.ReadFile("websites.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var websites map[string]uint

	err = yaml.Unmarshal(buf, &websites)

	if err != nil {
		log.Fatal(err)
	}

	(*c).WebSites = websites
}
