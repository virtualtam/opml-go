// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html/charset"
)

// Marshal returns the XML encoding of a Document.
func Marshal(d *Document) ([]byte, error) {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	writer.WriteString(xml.Header)

	encoder := xml.NewEncoder(writer)
	encoder.Indent("", "  ")

	if err := encoder.Encode(d); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

// Unmarshal unmarshals a []byte representation of an OPML file and returns the
// corresponding Document.
func Unmarshal(buf []byte) (*Document, error) {
	r := bytes.NewReader(buf)
	return unmarshal(r)
}

// UnmarshalFile unmarshals an OPML file and returns the corresponding Document.
func UnmarshalFile(filePath string) (*Document, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return &Document{}, err
	}
	defer file.Close()

	return unmarshal(file)
}

// Unmarshal unmarshals a string representation of an OPML file and returns the
// corresponding Document.
func UnmarshalString(data string) (*Document, error) {
	r := strings.NewReader(data)
	return unmarshal(r)
}

func unmarshal(r io.Reader) (*Document, error) {
	decoder := xml.NewDecoder(r)
	decoder.CharsetReader = charset.NewReaderLabel

	document := &Document{}

	if err := decoder.Decode(document); err != nil {
		return &Document{}, err
	}

	return document, nil
}
