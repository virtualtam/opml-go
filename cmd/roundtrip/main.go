// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/virtualtam/opml-go"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("missing input filename")
	}

	filePath := os.Args[1]

	document, err := opml.UnmarshalFile(filePath)
	if err != nil {
		fmt.Println("failed to unmarshal file:", err)
		os.Exit(1)
	}

	m, err := opml.Marshal(document)
	if err != nil {
		fmt.Println("failed to marshal document:", err)
		os.Exit(1)
	}

	fmt.Print(string(m))
}
