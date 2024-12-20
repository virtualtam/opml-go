// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
)

var (
	specDocumentCategory = Document{
		XMLName: xml.Name{
			Local: "opml",
		},
		Version: Version2,
		Head: Head{
			Title:       "Illustrating the category attribute",
			DateCreated: mustDecodeRFC1123Time("Mon, 31 Oct 2005 19:23:00 GMT"),
		},
		Body: Body{
			Outlines: []Outline{
				{
					Text: "The Mets are the best team in baseball.",
					Categories: []string{
						"/Philosophy/Baseball/Mets",
						"/Tourism/New York",
					},
					Created: mustDecodeRFC1123Time("Mon, 31 Oct 2005 18:21:33 GMT"),
				},
			},
		},
	}

	specDocumentDirectory = Document{
		XMLName: xml.Name{
			Local: "opml",
		},
		Version: Version2,
		Head: Head{
			Title:           "scriptingNewsDirectory.opml",
			DateCreated:     mustDecodeRFC1123Time("Thu, 13 Oct 2005 15:34:07 GMT"),
			DateModified:    mustDecodeRFC1123Time("Tue, 25 Oct 2005 21:33:57 GMT"),
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
					Created: mustDecodeRFC1123Time("Sun, 16 Oct 2005 05:56:10 GMT"),
					Type:    OutlineTypeLink,
					Url:     "http://hosting.opml.org/dave/mySites.opml",
				},
				{
					Text:    "News.Com top 100 OPML",
					Created: mustDecodeRFC1123Time("Tue, 25 Oct 2005 21:33:28 GMT"),
					Type:    OutlineTypeLink,
					Url:     "http://news.com.com/html/ne/blogs/CNETNewsBlog100.opml",
				},
				{
					Text:    "BloggerCon III Blogroll",
					Created: mustDecodeRFC1123Time("Mon, 24 Oct 2005 05:23:52 GMT"),
					Type:    OutlineTypeLink,
					Url:     "http://static.bloggercon.org/iii/blogroll.opml",
				},
				{
					Text: "TechCrunch reviews",
					Type: OutlineTypeLink,
					Url:  "http://hosting.opml.org/techcrunch.opml.org/TechCrunch.opml",
				},
				{
					Text: "Tod Maffin's directory of Public Radio podcasts",
					Type: OutlineTypeLink,
					Url:  "http://todmaffin.com/radio.opml",
				},
				{
					Text: "Adam Curry's iPodder.org directory",
					Type: OutlineTypeLink,
					Url:  "http://homepage.mac.com/dailysourcecode/DSC/ipodderDirectory.opml",
				},
				{
					Text:    "Memeorandum",
					Created: mustDecodeRFC1123Time("Thu, 13 Oct 2005 15:19:05 GMT"),
					Type:    OutlineTypeLink,
					Url:     "http://tech.memeorandum.com/index.opml",
				},
				{
					Text:    "DaveNet archive",
					Created: mustDecodeRFC1123Time("Wed, 12 Oct 2005 01:39:56 GMT"),
					Type:    OutlineTypeLink,
					Url:     "http://davenet.opml.org/index.opml",
				},
			},
		},
	}

	specDocumentPlacesLived = Document{
		XMLName: xml.Name{
			Local: "opml",
		},
		Version: Version2,
		Head: Head{
			Title:           "placesLived.opml",
			DateCreated:     mustDecodeRFC1123Time("Mon, 27 Feb 2006 12:09:48 GMT"),
			DateModified:    mustDecodeRFC1123Time("Mon, 27 Feb 2006 12:11:44 GMT"),
			OwnerName:       "Dave Winer",
			ExpansionState:  []int{1, 2, 5, 10, 13, 15},
			VertScrollState: 1,
			WindowTop:       242,
			WindowLeft:      329,
			WindowBottom:    665,
			WindowRight:     547,
		},
		Body: Body{
			Outlines: []Outline{
				{
					Text: "Places I've lived",
					Outlines: []Outline{
						{
							Text: "Boston",
							Outlines: []Outline{
								{Text: "Cambridge"},
								{Text: "West Newton"},
							},
						},
						{
							Text: "Bay Area",
							Outlines: []Outline{
								{Text: "Mountain View"},
								{Text: "Los Gatos"},
								{Text: "Palo Alto"},
								{Text: "Woodside"},
							},
						},
						{
							Text: "New Orleans",
							Outlines: []Outline{
								{Text: "Uptown"},
								{Text: "Metairie"},
							},
						},
						{
							Text: "Wisconsin",
							Outlines: []Outline{
								{Text: "Madison"},
							},
						},
						{
							Text: "Florida",
							Type: OutlineTypeInclusion,
							Url:  "http://hosting.opml.org/dave/florida.opml",
						},
						{
							Text: "New York",
							Outlines: []Outline{
								{Text: "Jackson Heights"},
								{Text: "Flushing"},
								{Text: "The Bronx"},
							},
						},
					},
				},
			},
		},
	}

	specDocumentSimpleScript = Document{
		XMLName: xml.Name{
			Local: "opml",
		},
		Version: Version2,
		Head: Head{
			Title:           "workspace.userlandsamples.doSomeUpstreaming",
			DateCreated:     mustDecodeRFC1123Time("Mon, 11 Feb 2002 22:48:02 GMT"),
			DateModified:    mustDecodeRFC1123Time("Sun, 30 Oct 2005 03:30:17 GMT"),
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
		Version: Version2,
		Head: Head{
			Title:           "states.opml",
			DateCreated:     mustDecodeRFC1123Time("Tue, 15 Mar 2005 16:35:45 GMT"),
			DateModified:    mustDecodeRFC1123Time("Thu, 14 Jul 2005 23:41:05 GMT"),
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
											Created: mustDecodeRFC1123Time("Tue, 12 Jul 2005 23:56:35 GMT"),
										},
										{
											Text:    "Las Vegas",
											Created: mustDecodeRFC1123Time("Tue, 12 Jul 2005 23:56:37 GMT"),
										},
										{
											Text:    "Ely",
											Created: mustDecodeRFC1123Time("Tue, 12 Jul 2005 23:56:39 GMT"),
										},
										{
											Text:    "Gerlach",
											Created: mustDecodeRFC1123Time("Tue, 12 Jul 2005 23:56:47 GMT"),
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
		Version: Version2,
		Head: Head{
			Title:           "mySubscriptions.opml",
			DateCreated:     mustDecodeRFC1123Time("Sat, 18 Jun 2005 12:11:52 GMT"),
			DateModified:    mustDecodeRFC1123Time("Tue, 02 Aug 2005 21:42:48 GMT"),
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
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion2,
					Title:       "CNET News.com",
					Description: "Tech news and business reports by CNET News.com. Focused on information technology, core topics include computers, hardware, software, networking, and Internet media.",
					Language:    "unknown",
					HtmlUrl:     "http://news.com.com/",
					XmlUrl:      "http://news.com.com/2547-1_3-0-5.xml",
				},
				{
					Text:        "washingtonpost.com - Politics",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion2,
					Title:       "washingtonpost.com - Politics",
					Description: "Politics",
					Language:    "unknown",
					HtmlUrl:     "http://www.washingtonpost.com/wp-dyn/politics?nav=rss_politics",
					XmlUrl:      "http://www.washingtonpost.com/wp-srv/politics/rssheadlines.xml",
				},
				{
					Text:        "Scobleizer: Microsoft Geek Blogger",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion2,
					Title:       "Scobleizer: Microsoft Geek Blogger",
					Description: "Robert Scoble's look at geek and Microsoft life.",
					Language:    "unknown",
					HtmlUrl:     "http://radio.weblogs.com/0001011/",
					XmlUrl:      "http://radio.weblogs.com/0001011/rss.xml",
				},
				{
					Text:        "Yahoo! News: Technology",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion2,
					Title:       "Yahoo! News: Technology",
					Description: "Technology",
					Language:    "unknown",
					HtmlUrl:     "http://news.yahoo.com/news?tmpl=index&cid=738",
					XmlUrl:      "http://rss.news.yahoo.com/rss/tech",
				},
				{
					Text:        "Workbench",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion2,
					Title:       "Workbench",
					Description: "Programming and publishing news and comment",
					Language:    "unknown",
					HtmlUrl:     "http://www.cadenhead.org/workbench/",
					XmlUrl:      "http://www.cadenhead.org/workbench/rss.xml",
				},
				{
					Text:        "Christian Science Monitor | Top Stories",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion1,
					Title:       "Christian Science Monitor | Top Stories",
					Description: "Read the front page stories of csmonitor.com.",
					Language:    "unknown",
					HtmlUrl:     "http://csmonitor.com",
					XmlUrl:      "http://www.csmonitor.com/rss/top.rss",
				},
				{
					Text:        "Dictionary.com Word of the Day",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion1,
					Title:       "Dictionary.com Word of the Day",
					Description: "A new word is presented every day with its definition and example sentences from actual published works.",
					Language:    "unknown",
					HtmlUrl:     "http://dictionary.reference.com/wordoftheday/",
					XmlUrl:      "http://www.dictionary.com/wordoftheday/wotd.rss",
				},
				{
					Text:        "The Motley Fool",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion1,
					Title:       "The Motley Fool",
					Description: "To Educate, Amuse, and Enrich",
					Language:    "unknown",
					HtmlUrl:     "http://www.fool.com",
					XmlUrl:      "http://www.fool.com/xml/foolnews_rss091.xml",
				},
				{
					Text:        "InfoWorld: Top News",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion2,
					Title:       "InfoWorld: Top News",
					Description: "The latest on Top News from InfoWorld",
					Language:    "unknown",
					HtmlUrl:     "http://www.infoworld.com/news/index.html",
					XmlUrl:      "http://www.infoworld.com/rss/news.xml",
				},
				{
					Text:        "NYT > Business",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion2,
					Title:       "NYT > Business",
					Description: "Find breaking news & business news on Wall Street, media & advertising, international business, banking, interest rates, the stock market, currencies & funds.",
					Language:    "unknown",
					HtmlUrl:     "http://www.nytimes.com/pages/business/index.html?partner=rssnyt",
					XmlUrl:      "http://www.nytimes.com/services/xml/rss/nyt/Business.xml",
				},
				{
					Text:     "NYT > Technology",
					Type:     "rss",
					Version:  RSSVersion2,
					Title:    "NYT > Technology",
					Language: "unknown",
					HtmlUrl:  "http://www.nytimes.com/pages/technology/index.html?partner=rssnyt",
					XmlUrl:   "http://www.nytimes.com/services/xml/rss/nyt/Technology.xml",
				},
				{
					Text:        "Scripting News",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion2,
					Title:       "Scripting News",
					Description: "It's even worse than it appears.",
					Language:    "unknown",
					HtmlUrl:     "http://www.scripting.com/",
					XmlUrl:      "http://www.scripting.com/rss.xml",
				},
				{
					Text:        "Wired News",
					Type:        OutlineTypeSubscription,
					Version:     RSSVersion1,
					Title:       "Wired News",
					Description: "Technology, and the way we do business, is changing the world we know. Wired News is a technology - and business-oriented news service feeding an intelligent, discerning audience. What role does technology play in the day-to-day living of your life? Wired News tells you. How has evolving technology changed the face of the international business world? Wired News puts you in the picture.",
					Language:    "unknown",
					HtmlUrl:     "http://www.wired.com/",
					XmlUrl:      "http://www.wired.com/news_drop/netcenter/netcenter.rdf",
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
			tname:             "places lived",
			document:          specDocumentPlacesLived,
			referenceFileName: "placesLived.opml",
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
			referenceFilePath := filepath.Join("testdata", "spec", "marshal", tc.referenceFileName)

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
			tname:         "places lived",
			inputFileName: "placesLived.opml",
			want:          specDocumentPlacesLived,
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
			inputFilePath := filepath.Join("testdata", "spec", "unmarshal", tc.inputFileName)

			got, err := UnmarshalFile(inputFilePath)
			if err != nil {
				t.Fatalf("want no error, got %q", err)
			}

			AssertDocumentsEqual(t, *got, tc.want)
		})
	}
}
