// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"encoding/xml"
	"path/filepath"
	"testing"
)

var (
	feedReaderDocumentFeedly = Document{
		XMLName: xml.Name{Local: "opml"},
		Version: Version1,
		Head: Head{
			Title: "My subscriptions in feedly Cloud",
		},
		Body: Body{
			Outlines: []Outline{
				{
					Text:  "Programming",
					Title: "Programming",
					Outlines: []Outline{
						{
							Text:    "Elixir Lang",
							Title:   "Elixir Lang",
							Type:    OutlineTypeSubscription,
							HtmlUrl: "http://elixir-lang.org",
							XmlUrl:  "https://feeds.feedburner.com/ElixirLang",
						},
						{
							Text:    "Zephyr Project",
							Title:   "Zephyr Project",
							Type:    OutlineTypeSubscription,
							HtmlUrl: "https://www.zephyrproject.org",
							XmlUrl:  "https://www.zephyrproject.org/feed/",
						},
						{
							Text:    "Python Insider",
							Title:   "Python Insider",
							Type:    OutlineTypeSubscription,
							HtmlUrl: "https://pythoninsider.blogspot.com/",
							XmlUrl:  "https://feeds.feedburner.com/PythonInsider",
						},
					},
				},
				{
					Text:  "Games",
					Title: "Games",
					Outlines: []Outline{
						{
							Text:    "BeamNG.drive",
							Title:   "BeamNG.drive",
							Type:    OutlineTypeSubscription,
							HtmlUrl: "https://www.beamng.com/game/",
							XmlUrl:  "https://www.beamng.com/game/index.xml",
						},
						{
							Text:    "Vintage Story",
							Title:   "Vintage Story",
							Type:    OutlineTypeSubscription,
							HtmlUrl: "https://www.vintagestory.at/blog.html/",
							XmlUrl:  "https://www.vintagestory.at/blog.html/?rss=1",
						},
					},
				},
			},
		},
	}

	feedReaderDocumentNewsblur = Document{
		XMLName: xml.Name{Local: "opml"},
		Version: Version1_1,
		Head: Head{
			Title:        "NewsBlur Feeds",
			DateCreated:  mustDecodeRFC1123Time("Thu, 07 Nov 2024 20:18:01.109756 GMT"),
			DateModified: mustDecodeRFC1123Time("Thu, 07 Nov 2024 20:18:01.109756 GMT"),
		},
		Body: Body{
			Outlines: []Outline{
				{
					Text:  "Security",
					Title: "Security",
					Outlines: []Outline{
						{
							Text:    "Google Online Security Blog",
							Title:   "Google Online Security Blog",
							Type:    OutlineTypeSubscription,
							Version: RSSVersion1,
							HtmlUrl: "http://security.googleblog.com/",
							XmlUrl:  "http://www.blogger.com/feeds/1176949257541686127/posts/default?max-results=25&redirect=false&start-index=26",
						},
						{
							Text:    "Blog on Library",
							Title:   "Blog on Library",
							Type:    OutlineTypeSubscription,
							Version: RSSVersion1,
							HtmlUrl: "https://www.openssl.org/blog/",
							XmlUrl:  "https://openssl-library.org:443/post/atom.xml",
						},
						{
							Text:    "Schneier on Security",
							Title:   "Schneier on Security",
							Type:    OutlineTypeSubscription,
							Version: RSSVersion1,
							HtmlUrl: "https://www.schneier.com",
							XmlUrl:  "https://www.schneier.com/feed/atom/",
						},
					},
				},
				{
					Text:  "Self-Hosted",
					Title: "Self-Hosted",
				},
				{
					Text:  "Cryptography",
					Title: "Cryptography",
				},
				{
					Text:  "Programming",
					Title: "Programming",
					Outlines: []Outline{
						{
							Text:    "Git Rev News",
							Title:   "Git Rev News",
							Type:    OutlineTypeSubscription,
							Version: RSSVersion1,
							HtmlUrl: "https://git.github.io/rev_news/",
							XmlUrl:  "https://git.github.io/feed.xml",
						},
						{
							Text:    "The Go Programming Language Blog",
							Title:   "The Go Programming Language Blog",
							Type:    OutlineTypeSubscription,
							Version: RSSVersion1,
							HtmlUrl: "tag:blog.golang.org,2013:blog.golang.org",
							XmlUrl:  "https://go.dev/blog/feed.atom",
						},
					},
				},
			},
		},
	}
)

func TestUnmarshalFeedReader(t *testing.T) {
	cases := []struct {
		tname         string
		inputFileName string
		want          Document
	}{
		{
			tname:         "feedly",
			inputFileName: "feedly.opml",
			want:          feedReaderDocumentFeedly,
		},
		{
			tname:         "newsblur",
			inputFileName: "newsblur.opml",
			want:          feedReaderDocumentNewsblur,
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			inputFilePath := filepath.Join("testdata", "feedreader", tc.inputFileName)

			got, err := UnmarshalFile(inputFilePath)
			if err != nil {
				t.Fatalf("want no error, got %q", err)
			}

			assertDocumentsEqual(t, *got, tc.want)
		})
	}
}
