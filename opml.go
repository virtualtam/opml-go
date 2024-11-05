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
type Document struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

// A Head contains the metadata for the OPML Document.
type Head struct {
	Title           string
	DateCreated     time.Time
	DateModified    time.Time
	OwnerName       string
	OwnerEmail      string
	ExpansionStates []int
	VertScrollState int
	WindowTop       int
	WindowLeft      int
	WindowBottom    int
	WindowRight     int
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
	Outlines []Outline `xml:"outline"`
}

// An Outline represents a text element, a subscription list item or a directory.
type Outline struct {
	Text string
	Type string

	Categories []string
	Created    time.Time
	URL        string

	Version     string
	Title       string
	Description string
	Language    string
	HTMLURL     string
	XMLURL      string
}

type xmlOutline struct {
	Text        string `xml:"text,attr"`
	Categories  string `xml:"category,attr,omitempty"`
	CreatedStr  string `xml:"created,attr,omitempty"`
	Description string `xml:"description,attr,omitempty"`
	HTMLURL     string `xml:"htmlUrl,attr,omitempty"`
	Language    string `xml:"language,attr,omitempty"`
	Title       string `xml:"title,attr,omitempty"`
	Type        string `xml:"type,attr,omitempty"`
	URL         string `xml:"url,attr,omitempty"`
	Version     string `xml:"version,attr,omitempty"`
	XMLURL      string `xml:"xmlUrl,attr,omitempty"`
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

	return nil
}
