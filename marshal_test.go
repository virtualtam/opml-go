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

	specDocumentDirectory = Document{
		XMLName: xml.Name{
			Local: "opml",
		},
		Version: "2.0",
		Head: Head{
			Title:           "scriptingNewsDirectory.opml",
			DateCreated:     mustParseTimeGMT("Thu, 13 Oct 2005 15:34:07 GMT"),
			DateModified:    mustParseTimeGMT("Tue, 25 Oct 2005 21:33:57 GMT"),
			OwnerName:       "Dave Winer",
			OwnerEmail:      "dwiner@yahoo.com",
			VertScrollState: 1,
			WindowTop:       105,
			WindowLeft:      466,
			WindowBottom:    386,
			WindowRight:     964,
		},
		Body: Body{
			Outlines: []Outline{
				{
					Text:    "Scripting News sites",
					Created: mustParseTimeGMT("Sun, 16 Oct 2005 05:56:10 GMT"),
					Type:    "link",
					URL:     "http://hosting.opml.org/dave/mySites.opml",
				},
				{
					Text:    "News.Com top 100 OPML",
					Created: mustParseTimeGMT("Tue, 25 Oct 2005 21:33:28 GMT"),
					Type:    "link",
					URL:     "http://news.com.com/html/ne/blogs/CNETNewsBlog100.opml",
				},
				{
					Text:    "BloggerCon III Blogroll",
					Created: mustParseTimeGMT("Mon, 24 Oct 2005 05:23:52 GMT"),
					Type:    "link",
					URL:     "http://static.bloggercon.org/iii/blogroll.opml",
				},
				{
					Text: "TechCrunch reviews",
					Type: "link",
					URL:  "http://hosting.opml.org/techcrunch.opml.org/TechCrunch.opml",
				},
				{
					Text: "Tod Maffin's directory of Public Radio podcasts",
					Type: "link",
					URL:  "http://todmaffin.com/radio.opml",
				},
				{
					Text: "Adam Curry's iPodder.org directory",
					Type: "link",
					URL:  "http://homepage.mac.com/dailysourcecode/DSC/ipodderDirectory.opml",
				},
				{
					Text:    "Memeorandum",
					Created: mustParseTimeGMT("Thu, 13 Oct 2005 15:19:05 GMT"),
					Type:    "link",
					URL:     "http://tech.memeorandum.com/index.opml",
				},
				{
					Text:    "DaveNet archive",
					Created: mustParseTimeGMT("Wed, 12 Oct 2005 01:39:56 GMT"),
					Type:    "link",
					URL:     "http://davenet.opml.org/index.opml",
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
		{
			tname:             "directory",
			document:          specDocumentDirectory,
			referenceFileName: "directory.opml",
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
		{
			tname:         "directory",
			inputFileName: "directory.opml",
			want:          specDocumentDirectory,
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

	// OPML metadata
	if got.XMLName.Local != want.XMLName.Local {
		t.Errorf("want XMLName.Local %q, got %q", want.XMLName.Local, got.XMLName.Local)
	}
	if got.Version != want.Version {
		t.Errorf("want Version %q, got %q", want.Version, got.Version)
	}

	// Head
	if !got.Head.DateCreated.Equal(want.Head.DateCreated) {
		t.Errorf("want Head > DateCreated %q, got %q", want.Head.DateCreated, got.Head.DateCreated)
	}
	if !got.Head.DateModified.Equal(want.Head.DateModified) {
		t.Errorf("want Head > DateModified %q, got %q", want.Head.DateModified, got.Head.DateModified)
	}
	if got.Head.OwnerName != want.Head.OwnerName {
		t.Errorf("want Head > OwnerName %q, got %q", want.Head.OwnerName, got.Head.OwnerName)
	}
	if got.Head.OwnerEmail != want.Head.OwnerEmail {
		t.Errorf("want Head > OwnerEmail %q, got %q", want.Head.OwnerEmail, got.Head.OwnerEmail)
	}
	if got.Head.VertScrollState != want.Head.VertScrollState {
		t.Errorf("want Head > VertScrollState %d, got %d", want.Head.VertScrollState, got.Head.VertScrollState)
	}
	if got.Head.WindowTop != want.Head.WindowTop {
		t.Errorf("want Head > WindowTop %d, got %d", want.Head.WindowTop, got.Head.WindowTop)
	}
	if got.Head.WindowLeft != want.Head.WindowLeft {
		t.Errorf("want Head > WindowLeft %d, got %d", want.Head.WindowLeft, got.Head.WindowLeft)
	}
	if got.Head.WindowBottom != want.Head.WindowBottom {
		t.Errorf("want Head > WindowBottom %d, got %d", want.Head.WindowBottom, got.Head.WindowBottom)
	}
	if got.Head.WindowRight != want.Head.WindowRight {
		t.Errorf("want Head > WindowRight %d, got %d", want.Head.WindowRight, got.Head.WindowRight)
	}

	// Body
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
		if gotOutline.Type != wantOutline.Type {
			t.Errorf("want Outline %d Type %q, got %q", index, wantOutline.Type, gotOutline.Type)
		}
		if gotOutline.URL != wantOutline.URL {
			t.Errorf("want Outline %d URL %q, got %q", index, wantOutline.URL, gotOutline.URL)
		}
	}
}
