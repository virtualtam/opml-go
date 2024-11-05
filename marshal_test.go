// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var (
	specDocumentCategory = Document{
		XMLName: xml.Name{
			Local: "opml",
		},
		Version: "2.0",
		Head: Head{
			Title:       "Illustrating the category attribute",
			DateCreated: mustParseTimeGMT("Mon, 31 Oct 2005 19:23:00 GMT"),
		},
		Body: Body{
			Outlines: []Outline{
				{
					Text: "The Mets are the best team in baseball.",
					Categories: []string{
						"/Philosophy/Baseball/Mets",
						"/Tourism/New York",
					},
					Created: mustParseTimeGMT("Mon, 31 Oct 2005 18:21:33 GMT"),
				},
			},
		},
	}
)

func TestMarshalFileSpec(t *testing.T) {
	cases := []struct {
		tname             string
		document          Document
		referenceFileName string
	}{
		{
			tname:             "category",
			document:          specDocumentCategory,
			referenceFileName: "category.opml",
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			referenceFilePath := filepath.Join("testdata", "marshal", tc.referenceFileName)

			wantBytes, err := os.ReadFile(referenceFilePath)
			if err != nil {
				t.Fatalf("failed to read reference output file: %q", err)
			}

			gotBytes, err := Marshal(&tc.document)

			if err != nil {
				t.Fatalf("want no error, got %q", err)
			}

			got := string(gotBytes)
			want := string(wantBytes)

			if got != want {
				t.Errorf("\nwant:\n%s\n\ngot:\n%s", want, got)
			}
		})
	}
}

func TestUnmarshalFileSpec(t *testing.T) {
	cases := []struct {
		tname         string
		inputFileName string
		want          Document
	}{
		{
			tname:         "category",
			inputFileName: "category.opml",
			want:          specDocumentCategory,
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			inputFilePath := filepath.Join("testdata", "unmarshal", tc.inputFileName)

			got, err := UnmarshalFile(inputFilePath)
			if err != nil {
				t.Fatalf("want no error, got %q", err)
			}

			assertDocumentsEqual(t, *got, tc.want)
		})
	}
}

func mustParseTimeGMT(dateStr string) time.Time {
	parsed, err := time.ParseInLocation(time.RFC1123, dateStr, locationGMT)
	if err != nil {
		panic(err)
	}

	return parsed
}

func assertDocumentsEqual(t *testing.T, got, want Document) {
	t.Helper()

	if got.XMLName.Local != want.XMLName.Local {
		t.Errorf("want XMLName.Local %q, got %q", want.XMLName.Local, got.XMLName.Local)
	}

	if got.Version != want.Version {
		t.Errorf("want Version %q, got %q", want.Version, got.Version)
	}

	if !got.Head.DateCreated.Equal(want.Head.DateCreated) {
		t.Errorf("want Head.DateCreated %q, got %q", want.Head.DateCreated, got.Head.DateCreated)
	}

	if len(got.Body.Outlines) != len(want.Body.Outlines) {
		t.Fatalf("want %d Outlines, got %d", len(want.Body.Outlines), len(got.Body.Outlines))
	}

	for index, wantOutline := range want.Body.Outlines {
		gotOutline := got.Body.Outlines[index]

		if gotOutline.Text != wantOutline.Text {
			t.Errorf("want Outline %d Text %q, got %q", index, wantOutline.Text, gotOutline.Text)
		}

		if len(gotOutline.Categories) != len(wantOutline.Categories) {
			t.Errorf("want Outline %d %d Categories, got %d", index, len(wantOutline.Categories), len(gotOutline.Categories))
		}

		for cIndex, wantCategory := range wantOutline.Categories {
			gotCategory := gotOutline.Categories[cIndex]

			if gotCategory != wantCategory {
				t.Errorf("want Outline %d Category %d %q, got %q", index, cIndex, wantCategory, gotCategory)
			}
		}

		if !gotOutline.Created.Equal(wantOutline.Created) {
			t.Errorf("want Outline %d Created %q, got %q", index, wantOutline.Created, gotOutline.Created)
		}
	}
}
