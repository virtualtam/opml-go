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

	specDocumentSimpleScript = Document{
		XMLName: xml.Name{
			Local: "opml",
		},
		Version: "2.0",
		Head: Head{
			Title:           "workspace.userlandsamples.doSomeUpstreaming",
			DateCreated:     mustParseTimeGMT("Mon, 11 Feb 2002 22:48:02 GMT"),
			DateModified:    mustParseTimeGMT("Sun, 30 Oct 2005 03:30:17 GMT"),
			OwnerName:       "Dave Winer",
			OwnerEmail:      "dwiner@yahoo.com",
			ExpansionState:  []int{1, 2, 4},
			VertScrollState: 1,
			WindowTop:       74,
			WindowLeft:      41,
			WindowBottom:    314,
			WindowRight:     475,
		},
		Body: Body{
			Outlines: []Outline{
				{
					Text:      "Changes",
					IsComment: true,
					Outlines: []Outline{
						{
							Text: "1/3/02; 4:54:25 PM by DW",
							Outlines: []Outline{
								{
									Text: `Change "playlist" to "radio".`,
								},
							},
						},
						{
							Text:      "2/12/01; 1:49:33 PM by DW",
							IsComment: true,
							Outlines: []Outline{
								{
									Text: "Test upstreaming by sprinkling a few files in a nice new test folder.",
								},
							},
						},
					},
				},
				{
					Text: "on writetestfile (f, size)",
					Outlines: []Outline{
						{
							Text:         "file.surefilepath (f)",
							IsBreakpoint: true,
						},
						{
							Text: `file.writewholefile (f, string.filledstring ("x", size))`,
						},
					},
				},
				{
					Text: `local (folder = user.radio.prefs.wwwfolder + "test\\largefiles\\")`,
				},
				{
					Text: "for ch = 'a' to 'z'",
					Outlines: []Outline{
						{
							Text: `writetestfile (folder + ch + ".html", random (1000, 16000))`,
						},
					},
				},
			},
		},
	}

	specDocumentStates = Document{
		XMLName: xml.Name{
			Local: "opml",
		},
		Version: "2.0",
		Head: Head{
			Title:           "states.opml",
			DateCreated:     mustParseTimeGMT("Tue, 15 Mar 2005 16:35:45 GMT"),
			DateModified:    mustParseTimeGMT("Thu, 14 Jul 2005 23:41:05 GMT"),
			OwnerName:       "Dave Winer",
			OwnerEmail:      "dave@scripting.com",
			ExpansionState:  []int{1, 6, 13, 16, 18, 20},
			VertScrollState: 1,
			WindowTop:       106,
			WindowLeft:      106,
			WindowBottom:    558,
			WindowRight:     479,
		},
		Body: Body{
			Outlines: []Outline{
				{
					Text: "United States",
					Outlines: []Outline{
						{
							Text: "Far West",
							Outlines: []Outline{
								{Text: "Alaska"},
								{Text: "California"},
								{Text: "Hawaii"},
								{
									Text: "Nevada",
									Outlines: []Outline{
										{
											Text:    "Reno",
											Created: mustParseTimeGMT("Tue, 12 Jul 2005 23:56:35 GMT"),
										},
										{
											Text:    "Las Vegas",
											Created: mustParseTimeGMT("Tue, 12 Jul 2005 23:56:37 GMT"),
										},
										{
											Text:    "Ely",
											Created: mustParseTimeGMT("Tue, 12 Jul 2005 23:56:39 GMT"),
										},
										{
											Text:    "Gerlach",
											Created: mustParseTimeGMT("Tue, 12 Jul 2005 23:56:47 GMT"),
										},
									},
								},
								{Text: "Oregon"},
								{Text: "Washington"},
							},
						},
						{
							Text: "Great Plains",
							Outlines: []Outline{
								{Text: "Kansas"},
								{Text: "Nebraska"},
								{Text: "North Dakota"},
								{Text: "Oklahoma"},
								{Text: "South Dakota"},
							},
						},
						{
							Text: "Mid-Atlantic",
							Outlines: []Outline{
								{Text: "Delaware"},
								{Text: "Maryland"},
								{Text: "New Jersey"},
								{Text: "New York"},
								{Text: "Pennsylvania"},
							},
						},
						{
							Text: "Midwest",
							Outlines: []Outline{
								{Text: "Illinois"},
								{Text: "Indiana"},
								{Text: "Iowa"},
								{Text: "Kentucky"},
								{Text: "Michigan"},
								{Text: "Minnesota"},
								{Text: "Missouri"},
								{Text: "Ohio"},
								{Text: "West Virginia"},
								{Text: "Wisconsin"},
							},
						},
						{
							Text: "Mountains",
							Outlines: []Outline{
								{Text: "Colorado"},
								{Text: "Idaho"},
								{Text: "Montana"},
								{Text: "Utah"},
								{Text: "Wyoming"},
							},
						},
						{
							Text: "New England",
							Outlines: []Outline{
								{Text: "Connecticut"},
								{Text: "Maine"},
								{Text: "Massachusetts"},
								{Text: "New Hampshire"},
								{Text: "Rhode Island"},
								{Text: "Vermont"},
							},
						},
						{
							Text: "South",
							Outlines: []Outline{
								{Text: "Alabama"},
								{Text: "Arkansas"},
								{Text: "Florida"},
								{Text: "Georgia"},
								{Text: "Louisiana"},
								{Text: "Mississippi"},
								{Text: "North Carolina"},
								{Text: "South Carolina"},
								{Text: "Tennessee"},
								{Text: "Virginia"},
							},
						},
						{
							Text: "Southwest",
							Outlines: []Outline{
								{Text: "Arizona"},
								{Text: "New Mexico"},
								{Text: "Texas"},
							},
						},
					},
				},
			},
		},
	}

	specDocumentSubscriptionList = Document{
		XMLName: xml.Name{
			Local: "opml",
		},
		Version: "2.0",
		Head: Head{
			Title:           "mySubscriptions.opml",
			DateCreated:     mustParseTimeGMT("Sat, 18 Jun 2005 12:11:52 GMT"),
			DateModified:    mustParseTimeGMT("Tue, 02 Aug 2005 21:42:48 GMT"),
			OwnerName:       "Dave Winer",
			OwnerEmail:      "dave@scripting.com",
			VertScrollState: 1,
			WindowTop:       61,
			WindowLeft:      304,
			WindowBottom:    562,
			WindowRight:     842,
		},
		Body: Body{
			Outlines: []Outline{
				{
					Text:        "CNET News.com",
					Type:        "rss",
					Version:     "RSS2",
					Title:       "CNET News.com",
					Description: "Tech news and business reports by CNET News.com. Focused on information technology, core topics include computers, hardware, software, networking, and Internet media.",
					Language:    "unknown",
					HTMLURL:     "http://news.com.com/",
					XMLURL:      "http://news.com.com/2547-1_3-0-5.xml",
				},
				{
					Text:        "washingtonpost.com - Politics",
					Type:        "rss",
					Version:     "RSS2",
					Title:       "washingtonpost.com - Politics",
					Description: "Politics",
					Language:    "unknown",
					HTMLURL:     "http://www.washingtonpost.com/wp-dyn/politics?nav=rss_politics",
					XMLURL:      "http://www.washingtonpost.com/wp-srv/politics/rssheadlines.xml",
				},
				{
					Text:        "Scobleizer: Microsoft Geek Blogger",
					Type:        "rss",
					Version:     "RSS2",
					Title:       "Scobleizer: Microsoft Geek Blogger",
					Description: "Robert Scoble's look at geek and Microsoft life.",
					Language:    "unknown",
					HTMLURL:     "http://radio.weblogs.com/0001011/",
					XMLURL:      "http://radio.weblogs.com/0001011/rss.xml",
				},
				{
					Text:        "Yahoo! News: Technology",
					Type:        "rss",
					Version:     "RSS2",
					Title:       "Yahoo! News: Technology",
					Description: "Technology",
					Language:    "unknown",
					HTMLURL:     "http://news.yahoo.com/news?tmpl=index&cid=738",
					XMLURL:      "http://rss.news.yahoo.com/rss/tech",
				},
				{
					Text:        "Workbench",
					Type:        "rss",
					Version:     "RSS2",
					Title:       "Workbench",
					Description: "Programming and publishing news and comment",
					Language:    "unknown",
					HTMLURL:     "http://www.cadenhead.org/workbench/",
					XMLURL:      "http://www.cadenhead.org/workbench/rss.xml",
				},
				{
					Text:        "Christian Science Monitor | Top Stories",
					Type:        "rss",
					Version:     "RSS",
					Title:       "Christian Science Monitor | Top Stories",
					Description: "Read the front page stories of csmonitor.com.",
					Language:    "unknown",
					HTMLURL:     "http://csmonitor.com",
					XMLURL:      "http://www.csmonitor.com/rss/top.rss",
				},
				{
					Text:        "Dictionary.com Word of the Day",
					Type:        "rss",
					Version:     "RSS",
					Title:       "Dictionary.com Word of the Day",
					Description: "A new word is presented every day with its definition and example sentences from actual published works.",
					Language:    "unknown",
					HTMLURL:     "http://dictionary.reference.com/wordoftheday/",
					XMLURL:      "http://www.dictionary.com/wordoftheday/wotd.rss",
				},
				{
					Text:        "The Motley Fool",
					Type:        "rss",
					Version:     "RSS",
					Title:       "The Motley Fool",
					Description: "To Educate, Amuse, and Enrich",
					Language:    "unknown",
					HTMLURL:     "http://www.fool.com",
					XMLURL:      "http://www.fool.com/xml/foolnews_rss091.xml",
				},
				{
					Text:        "InfoWorld: Top News",
					Type:        "rss",
					Version:     "RSS2",
					Title:       "InfoWorld: Top News",
					Description: "The latest on Top News from InfoWorld",
					Language:    "unknown",
					HTMLURL:     "http://www.infoworld.com/news/index.html",
					XMLURL:      "http://www.infoworld.com/rss/news.xml",
				},
				{
					Text:        "NYT > Business",
					Type:        "rss",
					Version:     "RSS2",
					Title:       "NYT > Business",
					Description: "Find breaking news & business news on Wall Street, media & advertising, international business, banking, interest rates, the stock market, currencies & funds.",
					Language:    "unknown",
					HTMLURL:     "http://www.nytimes.com/pages/business/index.html?partner=rssnyt",
					XMLURL:      "http://www.nytimes.com/services/xml/rss/nyt/Business.xml",
				},
				{
					Text:     "NYT > Technology",
					Type:     "rss",
					Version:  "RSS2",
					Title:    "NYT > Technology",
					Language: "unknown",
					HTMLURL:  "http://www.nytimes.com/pages/technology/index.html?partner=rssnyt",
					XMLURL:   "http://www.nytimes.com/services/xml/rss/nyt/Technology.xml",
				},
				{
					Text:        "Scripting News",
					Type:        "rss",
					Version:     "RSS2",
					Title:       "Scripting News",
					Description: "It's even worse than it appears.",
					Language:    "unknown",
					HTMLURL:     "http://www.scripting.com/",
					XMLURL:      "http://www.scripting.com/rss.xml",
				},
				{
					Text:        "Wired News",
					Type:        "rss",
					Version:     "RSS",
					Title:       "Wired News",
					Description: "Technology, and the way we do business, is changing the world we know. Wired News is a technology - and business-oriented news service feeding an intelligent, discerning audience. What role does technology play in the day-to-day living of your life? Wired News tells you. How has evolving technology changed the face of the international business world? Wired News puts you in the picture.",
					Language:    "unknown",
					HTMLURL:     "http://www.wired.com/",
					XMLURL:      "http://www.wired.com/news_drop/netcenter/netcenter.rdf",
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
		{
			tname:             "simple script",
			document:          specDocumentSimpleScript,
			referenceFileName: "simpleScript.opml",
		},
		{
			tname:             "states",
			document:          specDocumentStates,
			referenceFileName: "states.opml",
		},
		{
			tname:             "subscription list",
			document:          specDocumentSubscriptionList,
			referenceFileName: "subscriptionList.opml",
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
		{
			tname:         "simple script",
			inputFileName: "simpleScript.opml",
			want:          specDocumentSimpleScript,
		},
		{
			tname:         "states",
			inputFileName: "states.opml",
			want:          specDocumentStates,
		},
		{
			tname:         "subscription list",
			inputFileName: "subscriptionList.opml",
			want:          specDocumentSubscriptionList,
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
		if gotOutline.Type != wantOutline.Type {
			t.Errorf("want Outline %d Type %q, got %q", index, wantOutline.Type, gotOutline.Type)
		}
		if gotOutline.URL != wantOutline.URL {
			t.Errorf("want Outline %d URL %q, got %q", index, wantOutline.URL, gotOutline.URL)
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
		if gotOutline.HTMLURL != wantOutline.HTMLURL {
			t.Errorf("want Outline %d HTMLURL %q, got %q", index, wantOutline.HTMLURL, gotOutline.HTMLURL)
		}
		if gotOutline.XMLURL != wantOutline.XMLURL {
			t.Errorf("want Outline %d XMLURL %q, got %q", index, wantOutline.XMLURL, gotOutline.XMLURL)
		}

		if gotOutline.IsBreakpoint != wantOutline.IsBreakpoint {
			t.Errorf("want Outline %d IsBreakpoint %t, got %t", index, wantOutline.IsBreakpoint, gotOutline.IsBreakpoint)
		}
		if gotOutline.IsComment != wantOutline.IsComment {
			t.Errorf("want Outline %d IsComment %t, got %t", index, wantOutline.IsComment, gotOutline.IsComment)
		}

		assertOutlinesEqual(t, gotOutline.Outlines, wantOutline.Outlines)
	}
}
