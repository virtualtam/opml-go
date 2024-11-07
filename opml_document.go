// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

type (
	OutlineType string
	RSSVersion  string
)

const (
	Version1_1 string = "1.1"
	Version2   string = "2.0"

	OutlineTypeInclusion    OutlineType = "include"
	OutlineTypeLink         OutlineType = "link"
	OutlineTypeSubscription OutlineType = "rss"

	RSSVersion1 RSSVersion = "RSS"
	RSSVersion2 RSSVersion = "RSS2"
)

// A Document represents an OPML Document.
//
// See https://opml.org/spec2.opml for details on the OPML specification.
type Document struct {
	XMLName xml.Name `xml:"opml" json:"-"`
	Version string   `xml:"version,attr" json:"version"`
	Head    Head     `xml:"head" json:"head"`
	Body    Body     `xml:"body" json:"body"`
}

// A Head contains the metadata for the OPML Document.
type Head struct {
	// The title of the document.
	Title string

	// The date the document was created.
	DateCreated time.Time

	// The date indicating when the document was last modified.
	DateModified time.Time

	// The owner of the document.
	OwnerName string

	// The email address of the owner of the document.
	OwnerEmail string

	// A list of line numbers that are expanded.
	// The line numbers in the list indicate which headlines to expand.
	ExpansionState []int

	// A number, saying which  line of the outline is displayed on the top line of the window.
	VertScrollState int

	// The pixel location of the top edge of the window.
	WindowTop int

	// The pixel location of the left edge of the window.
	WindowLeft int

	// The pixel location of the bottom edge of the window.
	WindowBottom int

	// The pixel location of the right edge of the window.
	WindowRight int
}

func (h *Head) MarshalJSON() ([]byte, error) {
	mHead := newMarshalableHead(h)

	return json.Marshal(mHead)
}

func (h *Head) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	mHead := newMarshalableHead(h)

	return e.EncodeElement(mHead, start)
}

func (h *Head) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var mHead marshalableHead
	if err := d.DecodeElement(&mHead, &start); err != nil {
		return err
	}

	head, err := mHead.toHead()
	if err != nil {
		return err
	}

	*h = head

	return nil
}

type marshalableHead struct {
	Title              string `xml:"title,omitempty" json:"title"`
	DateCreatedStr     string `xml:"dateCreated,omitempty" json:"date_created,omitempty"`
	DateModifiedStr    string `xml:"dateModified,omitempty" json:"date_modified,omitempty"`
	OwnerName          string `xml:"ownerName,omitempty" json:"owner_name,omitempty"`
	OwnerEmail         string `xml:"ownerEmail,omitempty" json:"owner_email,omitempty"`
	ExpansionStatesStr string `xml:"expansionState,omitempty" json:"expansion_state,omitempty"`
	VertScrollState    int    `xml:"vertScrollState,omitempty" json:"vert_scroll_state,omitempty"`
	WindowTop          int    `xml:"windowTop,omitempty" json:"window_top,omitempty"`
	WindowLeft         int    `xml:"windowLeft,omitempty" json:"window_left,omitempty"`
	WindowBottom       int    `xml:"windowBottom,omitempty" json:"window_bottom,omitempty"`
	WindowRight        int    `xml:"windowRight,omitempty" json:"window_right,omitempty"`
}

func newMarshalableHead(h *Head) marshalableHead {
	mHead := marshalableHead{
		Title:           h.Title,
		OwnerName:       h.OwnerName,
		OwnerEmail:      h.OwnerEmail,
		VertScrollState: h.VertScrollState,
		WindowTop:       h.WindowTop,
		WindowLeft:      h.WindowLeft,
		WindowBottom:    h.WindowBottom,
		WindowRight:     h.WindowRight,
	}

	if !h.DateCreated.IsZero() {
		mHead.DateCreatedStr = encodeRFC1123Time(h.DateCreated)
	}
	if !h.DateModified.IsZero() {
		mHead.DateModifiedStr = encodeRFC1123Time(h.DateModified)
	}

	if len(h.ExpansionState) > 0 {
		var statesStr []string
		for _, state := range h.ExpansionState {
			statesStr = append(statesStr, strconv.Itoa(state))
		}

		mHead.ExpansionStatesStr = strings.Join(statesStr, ", ")
	}

	return mHead
}

func (mHead *marshalableHead) toHead() (Head, error) {
	h := Head{
		Title:           mHead.Title,
		OwnerName:       mHead.OwnerName,
		OwnerEmail:      mHead.OwnerEmail,
		VertScrollState: mHead.VertScrollState,
		WindowTop:       mHead.WindowTop,
		WindowLeft:      mHead.WindowLeft,
		WindowBottom:    mHead.WindowBottom,
		WindowRight:     mHead.WindowRight,
	}

	if mHead.DateCreatedStr != "" {
		dateCreated, err := decodeTime(mHead.DateCreatedStr)
		if err != nil {
			return Head{}, err
		}

		h.DateCreated = dateCreated
	}

	if mHead.DateModifiedStr != "" {
		dateModified, err := decodeTime(mHead.DateModifiedStr)
		if err != nil {
			return Head{}, err
		}

		h.DateModified = dateModified
	}

	if mHead.ExpansionStatesStr != "" {
		var expansionStates []int

		statesStr := strings.Split(mHead.ExpansionStatesStr, ",")
		for _, stateStr := range statesStr {
			stateStr = strings.TrimSpace(stateStr)

			if stateStr == "" {
				continue
			}

			state, err := strconv.Atoi(stateStr)
			if err != nil {
				return Head{}, err
			}

			expansionStates = append(expansionStates, state)
		}

		h.ExpansionState = expansionStates
	}

	return h, nil
}

// A Body contains one or more Outline elements.
type Body struct {
	Outlines []Outline `xml:"outline" json:"outlines"`
}

// An Outline represents a text element, a subscription list item or a directory.
type Outline struct {
	// The Text that is displayed when an outliner opens the OPML document.
	Text string

	// Special: The Type indicates how the attributes of the Outline are interpreted.
	Type OutlineType

	// Special: Indicates whether a breakpoint is set on this outline.
	IsBreakpoint bool

	// Special: Indicates whether the outline is commented.
	//
	// If an outline is commented, all subordinate outlines are considered to also be commented.
	IsComment bool

	// Special: A list of category strings.
	Categories []string

	// Special: The date when the outline node was created.
	Created time.Time

	// Inclusion: The HTTP address of the included link.
	Url string

	// Subscription: Version of RSS/Atom that is being supplied by the feed.
	Version RSSVersion

	// Subscription: Title is the top-level title from the feed.
	Title string

	// Subscription: Description is the top-level description element from the feed.
	Description string

	// Subscription: Language is the top-level language element for the feed.
	Language string

	// Subscription: HtmlUrl is the top-level link element from the feed.
	HtmlUrl string

	// Subscription: XmlUrl is the address of the feed.
	XmlUrl string

	// Directory: Subordinated outlines, arbitrarily structured.
	Outlines []Outline
}

func (o *Outline) MarshalJSON() ([]byte, error) {
	mOutline := newMarshalableOutline(o)

	return json.Marshal(mOutline)
}

func (o *Outline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	mOutline := newMarshalableOutline(o)

	return e.EncodeElement(mOutline, start)
}

func (o *Outline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var mOutline marshalableOutline
	if err := d.DecodeElement(&mOutline, &start); err != nil {
		return err
	}

	outline, err := mOutline.toOutline()
	if err != nil {
		return err
	}

	*o = outline

	return nil
}

type marshalableOutline struct {
	Text string `xml:"text,attr" json:"text"`

	CategoriesStr string      `xml:"category,attr,omitempty" json:"categories,omitempty"`
	CreatedStr    string      `xml:"created,attr,omitempty" json:"created,omitempty"`
	Description   string      `xml:"description,attr,omitempty" json:"description,omitempty"`
	HtmlUrl       string      `xml:"htmlUrl,attr,omitempty" json:"html_url,omitempty"`
	IsBreakpoint  bool        `xml:"isBreakpoint,attr,omitempty" json:"is_breakpoint,omitempty"`
	IsComment     bool        `xml:"isComment,attr,omitempty" json:"is_comment,omitempty"`
	Language      string      `xml:"language,attr,omitempty" json:"language,omitempty"`
	Title         string      `xml:"title,attr,omitempty" json:"title,omitempty"`
	Type          OutlineType `xml:"type,attr,omitempty" json:"type,omitempty"`
	Url           string      `xml:"url,attr,omitempty" json:"url,omitempty"`
	Version       RSSVersion  `xml:"version,attr,omitempty" json:"version,omitempty"`
	XmlUrl        string      `xml:"xmlUrl,attr,omitempty" json:"xml_url,omitempty"`

	Outlines []Outline `xml:"outline" json:"outlines,omitempty"`
}

func newMarshalableOutline(o *Outline) marshalableOutline {
	mOutline := marshalableOutline{
		Text: o.Text,
		Type: o.Type,

		CategoriesStr: strings.Join(o.Categories, ","),
		Url:           o.Url,

		Version:     o.Version,
		Title:       o.Title,
		Description: o.Description,
		HtmlUrl:     o.HtmlUrl,
		Language:    o.Language,
		XmlUrl:      o.XmlUrl,

		IsBreakpoint: o.IsBreakpoint,
		IsComment:    o.IsComment,

		Outlines: o.Outlines,
	}

	if !o.Created.IsZero() {
		mOutline.CreatedStr = encodeRFC1123Time(o.Created)
	}

	return mOutline
}

func (mo *marshalableOutline) toOutline() (Outline, error) {
	outline := Outline{
		// Text fields
		Text: mo.Text,
		Type: mo.Type,

		// Special fields

		IsBreakpoint: mo.IsBreakpoint,
		IsComment:    mo.IsComment,

		// Inclusion fields
		Url: mo.Url,

		// Subscription fields
		Version:     mo.Version,
		Title:       mo.Title,
		Description: mo.Description,
		Language:    mo.Language,
		HtmlUrl:     mo.HtmlUrl,
		XmlUrl:      mo.XmlUrl,

		// Directory fields
		Outlines: mo.Outlines,
	}

	if mo.CategoriesStr != "" {
		outline.Categories = strings.Split(mo.CategoriesStr, ",")
	}

	if mo.CreatedStr != "" {
		created, err := decodeTime(mo.CreatedStr)
		if err != nil {
			return Outline{}, err
		}

		outline.Created = created
	}

	return outline, nil
}
