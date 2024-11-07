// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import "testing"

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
	assertOutlinesEqual(t, got.Body.Outlines, want.Body.Outlines)
}

func assertOutlinesEqual(t *testing.T, gotOutlines, wantOutlines []Outline) {
	t.Helper()

	if len(gotOutlines) != len(wantOutlines) {
		t.Fatalf("want %d Outlines, got %d", len(wantOutlines), len(gotOutlines))
	}

	for index, wantOutline := range wantOutlines {
		gotOutline := gotOutlines[index]

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
		if gotOutline.IsBreakpoint != wantOutline.IsBreakpoint {
			t.Errorf("want Outline %d IsBreakpoint %t, got %t", index, wantOutline.IsBreakpoint, gotOutline.IsBreakpoint)
		}
		if gotOutline.IsComment != wantOutline.IsComment {
			t.Errorf("want Outline %d IsComment %t, got %t", index, wantOutline.IsComment, gotOutline.IsComment)
		}
		if gotOutline.Type != wantOutline.Type {
			t.Errorf("want Outline %d Type %q, got %q", index, wantOutline.Type, gotOutline.Type)
		}

		if gotOutline.Url != wantOutline.Url {
			t.Errorf("want Outline %d URL %q, got %q", index, wantOutline.Url, gotOutline.Url)
		}

		if gotOutline.Version != wantOutline.Version {
			t.Errorf("want Outline %d Version %q, got %q", index, wantOutline.Version, gotOutline.Version)
		}
		if gotOutline.Title != wantOutline.Title {
			t.Errorf("want Outline %d Title %q, got %q", index, wantOutline.Title, gotOutline.Title)
		}
		if gotOutline.Description != wantOutline.Description {
			t.Errorf("want Outline %d Description %q, got %q", index, wantOutline.Description, gotOutline.Description)
		}
		if gotOutline.Language != wantOutline.Language {
			t.Errorf("want Outline %d Language %q, got %q", index, wantOutline.Language, gotOutline.Language)
		}
		if gotOutline.HtmlUrl != wantOutline.HtmlUrl {
			t.Errorf("want Outline %d HtmlUrl %q, got %q", index, wantOutline.HtmlUrl, gotOutline.HtmlUrl)
		}
		if gotOutline.XmlUrl != wantOutline.XmlUrl {
			t.Errorf("want Outline %d XmlUrl %q, got %q", index, wantOutline.XmlUrl, gotOutline.XmlUrl)
		}

		assertOutlinesEqual(t, gotOutline.Outlines, wantOutline.Outlines)
	}
}
