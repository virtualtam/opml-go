// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"fmt"
	"testing"
)

// AssertDocumentsEqual asserts two OPML documents are equal.
func AssertDocumentsEqual(t *testing.T, got, want Document) {
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
	AssertOutlinesEqual(t, got.Body.Outlines, want.Body.Outlines)
}

// AssertOutlinesEqual asserts two lists of OPML outlines are equal.
func AssertOutlinesEqual(t *testing.T, gotOutlines, wantOutlines []Outline) {
	t.Helper()

	assertOutlinesEqual(t, "", gotOutlines, wantOutlines)
}

// AssertOutlinesEqual asserts two lists of OPML outlines are equal.
func assertOutlinesEqual(t *testing.T, prefix string, gotOutlines, wantOutlines []Outline) {
	t.Helper()

	if len(gotOutlines) != len(wantOutlines) {
		t.Fatalf("want %d Outlines, got %d", len(wantOutlines), len(gotOutlines))
	}

	for index, wantOutline := range wantOutlines {
		gotOutline := gotOutlines[index]

		if gotOutline.Text != wantOutline.Text {
			t.Errorf("want Outline %s%d Text %q, got %q", prefix, index, wantOutline.Text, gotOutline.Text)
		}

		if len(gotOutline.Categories) != len(wantOutline.Categories) {
			t.Errorf("want Outline %s%d %d Categories, got %d", prefix, index, len(wantOutline.Categories), len(gotOutline.Categories))
		}

		for cIndex, wantCategory := range wantOutline.Categories {
			gotCategory := gotOutline.Categories[cIndex]

			if gotCategory != wantCategory {
				t.Errorf("want Outline %s%d Category %d %q, got %q", prefix, index, cIndex, wantCategory, gotCategory)
			}
		}

		if !gotOutline.Created.Equal(wantOutline.Created) {
			t.Errorf("want Outline %s%d Created %q, got %q", prefix, index, wantOutline.Created, gotOutline.Created)
		}
		if gotOutline.IsBreakpoint != wantOutline.IsBreakpoint {
			t.Errorf("want Outline %s%d IsBreakpoint %t, got %t", prefix, index, wantOutline.IsBreakpoint, gotOutline.IsBreakpoint)
		}
		if gotOutline.IsComment != wantOutline.IsComment {
			t.Errorf("want Outline %s%d IsComment %t, got %t", prefix, index, wantOutline.IsComment, gotOutline.IsComment)
		}
		if gotOutline.Type != wantOutline.Type {
			t.Errorf("want Outline %s%d Type %q, got %q", prefix, index, wantOutline.Type, gotOutline.Type)
		}

		if gotOutline.Url != wantOutline.Url {
			t.Errorf("want Outline %s%d URL %q, got %q", prefix, index, wantOutline.Url, gotOutline.Url)
		}

		if gotOutline.Version != wantOutline.Version {
			t.Errorf("want Outline %s%d Version %q, got %q", prefix, index, wantOutline.Version, gotOutline.Version)
		}
		if gotOutline.Title != wantOutline.Title {
			t.Errorf("want Outline %s%d Title %q, got %q", prefix, index, wantOutline.Title, gotOutline.Title)
		}
		if gotOutline.Description != wantOutline.Description {
			t.Errorf("want Outline %s%d Description %q, got %q", prefix, index, wantOutline.Description, gotOutline.Description)
		}
		if gotOutline.Language != wantOutline.Language {
			t.Errorf("want Outline %s%d Language %q, got %q", prefix, index, wantOutline.Language, gotOutline.Language)
		}
		if gotOutline.HtmlUrl != wantOutline.HtmlUrl {
			t.Errorf("want Outline %s%d HtmlUrl %q, got %q", prefix, index, wantOutline.HtmlUrl, gotOutline.HtmlUrl)
		}
		if gotOutline.XmlUrl != wantOutline.XmlUrl {
			t.Errorf("want Outline %s%d XmlUrl %q, got %q", prefix, index, wantOutline.XmlUrl, gotOutline.XmlUrl)
		}

		childPrefix := fmt.Sprintf("%s%d.", prefix, index)
		assertOutlinesEqual(t, childPrefix, gotOutline.Outlines, wantOutline.Outlines)
	}
}
