// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml_test

import (
	"fmt"
	"os"

	"github.com/virtualtam/opml-go"
)

func ExampleMarshal() {
	document := opml.Document{
		Version: "2.0",
		Head: opml.Head{
			Title:      "Feed subscriptions",
			OwnerName:  "Jane Doe",
			OwnerEmail: "jane@thedo.es",
		},
		Body: opml.Body{
			Outlines: []opml.Outline{
				{
					Text:  "Linux",
					Title: "Linux",
					Outlines: []opml.Outline{
						{
							Type:    opml.OutlineTypeRSS,
							Text:    "Bits from Debian",
							Title:   "Bits from Debian",
							HtmlUrl: "https://bits.debian.org/feeds/atom.xml",
							XmlUrl:  "https://bits.debian.org/feeds/atom.xml",
						},
						{
							Type:    opml.OutlineTypeRSS,
							Text:    "KXStudio News",
							Title:   "KXStudio News",
							HtmlUrl: "https://kx.studio/News",
							XmlUrl:  "https://kx.studio/News/?action=feed",
						},
					},
				},
				{
					Text:  "Social News",
					Title: "Social News",
					Outlines: []opml.Outline{
						{
							Type:    opml.OutlineTypeRSS,
							Text:    "Hacker News",
							Title:   "Hacker News",
							HtmlUrl: "https://news.ycombinator.com/",
							XmlUrl:  "https://news.ycombinator.com/rss",
						},
						{
							Type:    opml.OutlineTypeRSS,
							Text:    "Lobsters",
							Title:   "Lobsters",
							HtmlUrl: "https://lobste.rs",
							XmlUrl:  "https://lobste.rs/rss",
						},
						{
							Type:    opml.OutlineTypeRSS,
							Text:    "Phoronix",
							Title:   "Phoronix",
							HtmlUrl: "https://www.phoronix.com/",
							XmlUrl:  "https://www.phoronix.com/rss.php",
						},
					},
				},
			},
		},
	}

	m, err := opml.Marshal(&document)
	if err != nil {
		fmt.Println("failed to marshal data as XML:", err)
		os.Exit(1)
	}

	fmt.Print(string(m))

	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <opml version="2.0">
	//   <head>
	//     <title>Feed subscriptions</title>
	//     <ownerName>Jane Doe</ownerName>
	//     <ownerEmail>jane@thedo.es</ownerEmail>
	//   </head>
	//   <body>
	//     <outline text="Linux" title="Linux">
	//       <outline text="Bits from Debian" htmlUrl="https://bits.debian.org/feeds/atom.xml" title="Bits from Debian" type="rss" xmlUrl="https://bits.debian.org/feeds/atom.xml"></outline>
	//       <outline text="KXStudio News" htmlUrl="https://kx.studio/News" title="KXStudio News" type="rss" xmlUrl="https://kx.studio/News/?action=feed"></outline>
	//     </outline>
	//     <outline text="Social News" title="Social News">
	//       <outline text="Hacker News" htmlUrl="https://news.ycombinator.com/" title="Hacker News" type="rss" xmlUrl="https://news.ycombinator.com/rss"></outline>
	//       <outline text="Lobsters" htmlUrl="https://lobste.rs" title="Lobsters" type="rss" xmlUrl="https://lobste.rs/rss"></outline>
	//       <outline text="Phoronix" htmlUrl="https://www.phoronix.com/" title="Phoronix" type="rss" xmlUrl="https://www.phoronix.com/rss.php"></outline>
	//     </outline>
	//   </body>
	// </opml>
}
