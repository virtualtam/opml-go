// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
)

func newTextOutline(t *testing.T, fake faker.Faker) Outline {
	t.Helper()

	text := fake.Lorem().Sentence(3)

	return Outline{
		Text:  text,
		Title: text,
	}
}

func newDirectoryOutline(t *testing.T, fake faker.Faker, nChildOutlines int) Outline {
	t.Helper()

	outline := newTextOutline(t, fake)

	for i := 0; i < nChildOutlines; i++ {
		outline.Outlines = append(outline.Outlines, newTextOutline(t, fake))
	}

	return outline
}

func newInclusionOutline(t *testing.T, fake faker.Faker) Outline {
	t.Helper()

	outline := newTextOutline(t, fake)
	outline.Type = OutlineTypeInclusion
	outline.Url = fake.Internet().URL()

	return outline
}

func newLinkOutline(t *testing.T, fake faker.Faker) Outline {
	t.Helper()

	outline := newTextOutline(t, fake)
	outline.Type = OutlineTypeLink
	outline.Url = fake.Internet().URL()

	return outline
}

func newSubscriptionOutline(t *testing.T, fake faker.Faker) Outline {
	t.Helper()

	outline := newTextOutline(t, fake)
	outline.Type = OutlineTypeSubscription
	outline.Version = RSSVersion1
	outline.HtmlUrl = fake.Internet().URL()
	outline.XmlUrl = fmt.Sprintf("%s/feed", outline.HtmlUrl)

	return outline
}

func TestOutlineIsDirectory(t *testing.T) {
	fake := faker.New()

	cases := []struct {
		tname   string
		outline Outline
		want    bool
	}{
		{
			tname:   "directory",
			outline: newDirectoryOutline(t, fake, 3),
			want:    true,
		},
		{
			tname:   "inclusion",
			outline: newInclusionOutline(t, fake),
		},
		{
			tname:   "link",
			outline: newLinkOutline(t, fake),
		},
		{
			tname:   "subscription",
			outline: newSubscriptionOutline(t, fake),
		},
		{
			tname:   "text",
			outline: newTextOutline(t, fake),
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			got := tc.outline.IsDirectory()

			if got != tc.want {
				t.Errorf("want IsDirectory %t, got %t", tc.want, got)
			}
		})
	}
}

func TestOutlineOutlineType(t *testing.T) {
	fake := faker.New()

	cases := []struct {
		tname   string
		outline Outline
		want    OutlineType
	}{
		{
			tname:   "directory",
			outline: newDirectoryOutline(t, fake, 3),
			want:    OutlineTypeText,
		},
		{
			tname:   "inclusion",
			outline: newInclusionOutline(t, fake),
			want:    OutlineTypeInclusion,
		},
		{
			tname:   "link",
			outline: newLinkOutline(t, fake),
			want:    OutlineTypeLink,
		},
		{
			tname:   "subscription",
			outline: newSubscriptionOutline(t, fake),
			want:    OutlineTypeSubscription,
		},
		{
			tname:   "text",
			outline: newTextOutline(t, fake),
			want:    OutlineTypeText,
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			got := tc.outline.OutlineType()

			if got != tc.want {
				t.Errorf("want Type %q, got %q", tc.want, got)
			}
		})
	}
}
