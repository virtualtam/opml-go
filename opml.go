// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
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
	Title string `json:"title"`

	// The date the document was created.
	DateCreated time.Time `json:"date_created,omitempty"`

	// The date indicating when the document was last modified.
	DateModified time.Time `json:"date_modified,omitempty"`

	// The owner of the document.
	OwnerName string `json:"owner_name,omitempty"`

	// The email address of the owner of the document.
	OwnerEmail string `json:"owner_email,omitempty"`

	// A list of line numbers that are expanded.
	// The line numbers in the list indicate which headlines to expand.
	ExpansionStates []int `json:"expansion_states,omitempty"`

	// A number, saying which  line of the outline is displayed on the top line of the window.
	VertScrollState int `json:"vert_scroll_state,omitempty"`

	// The pixel location of the top edge of the window.
	WindowTop int `json:"window_top,omitempty"`

	// The pixel location of the left edge of the window.
	WindowLeft int `json:"window_left,omitempty"`

	// The pixel location of the bottom edge of the window.
	WindowBottom int `json:"window_bottom,omitempty"`

	// The pixel location of the right edge of the window.
	WindowRight int `json:"window_right,omitempty"`
}

type xmlHead struct {
	Title              string `xml:"title,omitempty"`
	DateCreatedStr     string `xml:"dateCreated,omitempty"`
	DateModifiedStr    string `xml:"dateModified,omitempty"`
	OwnerName          string `xml:"ownerName,omitempty"`
	OwnerEmail         string `xml:"ownerEmail,omitempty"`
	ExpansionStatesStr string `xml:"expansionState,omitempty"`
	VertScrollState    int    `xml:"vertScrollState,omitempty"`
	WindowTop          int    `xml:"windowTop,omitempty"`
	WindowLeft         int    `xml:"windowLeft,omitempty"`
	WindowBottom       int    `xml:"windowBottom,omitempty"`
	WindowRight        int    `xml:"windowRight,omitempty"`
}

func (h *Head) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	xmlHead := xmlHead{
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
		xmlHead.DateCreatedStr = formatRFC1123Time(h.DateCreated)
	}
	if !h.DateModified.IsZero() {
		xmlHead.DateModifiedStr = formatRFC1123Time(h.DateModified)
	}

	if len(h.ExpansionStates) > 0 {
		var statesStr []string
		for _, state := range h.ExpansionStates {
			statesStr = append(statesStr, strconv.Itoa(state))
		}

		xmlHead.ExpansionStatesStr = strings.Join(statesStr, ", ")
	}

	return e.EncodeElement(xmlHead, start)
}

func (h *Head) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var xmlHead xmlHead
	if err := d.DecodeElement(&xmlHead, &start); err != nil {
		return err
	}

	if xmlHead.DateCreatedStr != "" {
		dateCreated, err := parseRFC1123Time(xmlHead.DateCreatedStr)
		if err != nil {
			return err
		}

		h.DateCreated = dateCreated
	}

	if xmlHead.DateModifiedStr != "" {
		dateModified, err := parseRFC1123Time(xmlHead.DateModifiedStr)
		if err != nil {
			return err
		}

		h.DateModified = dateModified
	}

	if xmlHead.ExpansionStatesStr != "" {
		var expansionStates []int

		statesStr := strings.Split(xmlHead.ExpansionStatesStr, ",")
		for _, stateStr := range statesStr {
			stateStr = strings.TrimSpace(stateStr)

			if stateStr == "" {
				continue
			}

			state, err := strconv.Atoi(stateStr)
			if err != nil {
				return err
			}

			expansionStates = append(expansionStates, state)
		}

		h.ExpansionStates = expansionStates
	}

	h.Title = xmlHead.Title
	h.OwnerName = xmlHead.OwnerName
	h.OwnerEmail = xmlHead.OwnerEmail
	h.VertScrollState = xmlHead.VertScrollState
	h.WindowTop = xmlHead.WindowTop
	h.WindowLeft = xmlHead.WindowLeft
	h.WindowBottom = xmlHead.WindowBottom
	h.WindowRight = xmlHead.WindowRight

	return nil
}

// A Body contains one or more Outline elements.
type Body struct {
	Outlines []Outline `xml:"outline" json:"outlines"`
}

// An Outline represents a text element, a subscription list item or a directory.
type Outline struct {
	// The Text that is displayed when an outliner opens the OPML document.
	Text string `json:"text"`

	// Special: The Type indicates how the attributes of the Outline are interpreted.
	Type string `json:"type,omitempty"`

	// Special: Indicates whether a breakpoint is set on this outline.
	IsBreakpoint bool `json:"is_breakpoint,omitempty"`

	// Special: Indicates whether the outline is commented.
	//
	// If an outline is commented, all subordinate outlines are considered to also be commented.
	IsComment bool `json:"is_comment,omitempty"`

	// Special: A list of category strings.
	Categories []string `json:"categories,omitempty"`

	// Special: The date when the outline node was created.
	Created time.Time `json:"created,omitempty"`

	// Inclusion: The HTTP address of the included link.
	URL string `json:"url,omitempty"`

	// Subscription: Version of RSS/Atom that is being supplied by the feed.
	Version string `json:"version,omitempty"`

	// Subscription: Title is the top-level title from the feed.
	Title string `json:"title,omitempty"`

	// Subscription: Description is the top-level description element from the feed.
	Description string `json:"description,omitempty"`

	// Subscription: Language is the top-level language element for the feed.
	Language string `json:"language,omitempty"`

	// Subscription: HTMLURL is the top-level link element from the feed.
	HTMLURL string `json:"htmlurl,omitempty"`

	// Subscription: XMLURL is the address of the feed.
	XMLURL string `json:"xmlurl,omitempty"`

	// Directory: Subordinated outlines, arbitrarily structured.
	Outlines []Outline `json:"outlines,omitempty"`
}

type xmlOutline struct {
	Text string `xml:"text,attr"`

	Categories   string `xml:"category,attr,omitempty"`
	CreatedStr   string `xml:"created,attr,omitempty"`
	Description  string `xml:"description,attr,omitempty"`
	HTMLURL      string `xml:"htmlUrl,attr,omitempty"`
	IsBreakpoint bool   `xml:"isBreakpoint,attr,omitempty"`
	IsComment    bool   `xml:"isComment,attr,omitempty"`
	Language     string `xml:"language,attr,omitempty"`
	Title        string `xml:"title,attr,omitempty"`
	Type         string `xml:"type,attr,omitempty"`
	URL          string `xml:"url,attr,omitempty"`
	Version      string `xml:"version,attr,omitempty"`
	XMLURL       string `xml:"xmlUrl,attr,omitempty"`

	Outlines []Outline `xml:"outline"`
}

func (o *Outline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	xmlOutline := xmlOutline{
		Text: o.Text,
		Type: o.Type,

		Categories: strings.Join(o.Categories, ","),
		URL:        o.URL,

		Version:     o.Version,
		Title:       o.Title,
		Description: o.Description,
		HTMLURL:     o.HTMLURL,
		Language:    o.Language,
		XMLURL:      o.XMLURL,

		IsBreakpoint: o.IsBreakpoint,
		IsComment:    o.IsComment,

		Outlines: o.Outlines,
	}

	if !o.Created.IsZero() {
		xmlOutline.CreatedStr = formatRFC1123Time(o.Created)
	}

	return e.EncodeElement(xmlOutline, start)
}

func (o *Outline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var xmlOutline xmlOutline
	if err := d.DecodeElement(&xmlOutline, &start); err != nil {
		return err
	}

	o.Text = xmlOutline.Text
	o.Type = xmlOutline.Type

	o.URL = xmlOutline.URL

	if xmlOutline.Categories != "" {
		o.Categories = strings.Split(xmlOutline.Categories, ",")
	}

	if xmlOutline.CreatedStr != "" {
		created, err := parseRFC1123Time(xmlOutline.CreatedStr)
		if err != nil {
			return err
		}

		o.Created = created
	}

	o.Version = xmlOutline.Version
	o.Title = xmlOutline.Title
	o.Description = xmlOutline.Description
	o.Language = xmlOutline.Language
	o.HTMLURL = xmlOutline.HTMLURL
	o.XMLURL = xmlOutline.XMLURL

	o.IsBreakpoint = xmlOutline.IsBreakpoint
	o.IsComment = xmlOutline.IsComment

	o.Outlines = xmlOutline.Outlines

	return nil
}
