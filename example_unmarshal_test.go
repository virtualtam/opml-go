// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/virtualtam/opml-go"
)

func ExampleUnmarshal() {
	blob := `<?xml version="1.0" encoding="UTF-8"?>
<opml version="2.0">
  <head>
    <title>Feed subscriptions</title>
    <ownerName>Jane Doe</ownerName>
    <ownerEmail>jane@thedo.es</ownerEmail>
  </head>
  <body>
    <outline text="Linux" title="Linux">
      <outline text="Bits from Debian" htmlUrl="https://bits.debian.org/feeds/atom.xml" title="Bits from Debian" type="rss" xmlUrl="https://bits.debian.org/feeds/atom.xml"></outline>
      <outline text="KXStudio News" htmlUrl="https://kx.studio/News" title="KXStudio News" type="rss" xmlUrl="https://kx.studio/News/?action=feed"></outline>
    </outline>
    <outline text="Social News" title="Social News">
      <outline text="Hacker News" htmlUrl="https://news.ycombinator.com/" title="Hacker News" type="rss" xmlUrl="https://news.ycombinator.com/rss"></outline>
      <outline text="Lobsters" htmlUrl="https://lobste.rs" title="Lobsters" type="rss" xmlUrl="https://lobste.rs/rss"></outline>
      <outline text="Phoronix" htmlUrl="https://www.phoronix.com/" title="Phoronix" type="rss" xmlUrl="https://www.phoronix.com/rss.php"></outline>
    </outline>
  </body>
</opml>
`

	document, err := opml.Unmarshal([]byte(blob))
	if err != nil {
		fmt.Println("failed to unmarshal file:", err)
		os.Exit(1)
	}

	jsonData, err := json.MarshalIndent(document, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal data as JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))

	// Output:
	// {
	//   "version": "2.0",
	//   "head": {
	//     "title": "Feed subscriptions",
	//     "owner_name": "Jane Doe",
	//     "owner_email": "jane@thedo.es"
	//   },
	//   "body": {
	//     "outlines": [
	//       {
	//         "text": "Linux",
	//         "title": "Linux",
	//         "outlines": [
	//           {
	//             "text": "Bits from Debian",
	//             "html_url": "https://bits.debian.org/feeds/atom.xml",
	//             "title": "Bits from Debian",
	//             "type": "rss",
	//             "xml_url": "https://bits.debian.org/feeds/atom.xml"
	//           },
	//           {
	//             "text": "KXStudio News",
	//             "html_url": "https://kx.studio/News",
	//             "title": "KXStudio News",
	//             "type": "rss",
	//             "xml_url": "https://kx.studio/News/?action=feed"
	//           }
	//         ]
	//       },
	//       {
	//         "text": "Social News",
	//         "title": "Social News",
	//         "outlines": [
	//           {
	//             "text": "Hacker News",
	//             "html_url": "https://news.ycombinator.com/",
	//             "title": "Hacker News",
	//             "type": "rss",
	//             "xml_url": "https://news.ycombinator.com/rss"
	//           },
	//           {
	//             "text": "Lobsters",
	//             "html_url": "https://lobste.rs",
	//             "title": "Lobsters",
	//             "type": "rss",
	//             "xml_url": "https://lobste.rs/rss"
	//           },
	//           {
	//             "text": "Phoronix",
	//             "html_url": "https://www.phoronix.com/",
	//             "title": "Phoronix",
	//             "type": "rss",
	//             "xml_url": "https://www.phoronix.com/rss.php"
	//           }
	//         ]
	//       }
	//     ]
	//   }
	// }
}
