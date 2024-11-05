// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package opml

import (
	"encoding/xml"
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
	Title       string
	DateCreated time.Time
}

type xmlHead struct {
	Title          string `xml:"title,omitempty"`
	DateCreatedStr string `xml:"dateCreated,omitempty"`
}

func (h *Head) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	dateCreatedStr := formatRFC1123Time(h.DateCreated)

	xmlHead := xmlHead{
		Title:          h.Title,
		DateCreatedStr: dateCreatedStr,
	}

	return e.EncodeElement(xmlHead, start)
}

func (h *Head) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var xmlHead xmlHead
	if err := d.DecodeElement(&xmlHead, &start); err != nil {
		return err
	}

	dateCreated, err := parseRFC1123Time(xmlHead.DateCreatedStr)
	if err != nil {
		return err
	}

	h.Title = xmlHead.Title
	h.DateCreated = dateCreated

	return nil
}

// A Body contains one or more Outline elements.
type Body struct {
	Outlines []Outline `xml:"outline"`
}

// An Outline represents a text element, a subsciption list item or a directory.
type Outline struct {
	Text       string
	Categories []string
	Created    time.Time
}

type xmlOutline struct {
	Text       string `xml:"text,attr"`
	Categories string `xml:"category,attr,omitempty"`
	CreatedStr string `xml:"created,attr,omitempty"`
}

func (o *Outline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	createdStr := formatRFC1123Time(o.Created)

	xmlOutline := xmlOutline{
		Text:       o.Text,
		Categories: strings.Join(o.Categories, ","),
		CreatedStr: createdStr,
	}

	return e.EncodeElement(xmlOutline, start)
}

func (o *Outline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var xmlOutline xmlOutline
	if err := d.DecodeElement(&xmlOutline, &start); err != nil {
		return err
	}

	categories := strings.Split(xmlOutline.Categories, ",")

	created, err := parseRFC1123Time(xmlOutline.CreatedStr)
	if err != nil {
		return err
	}

	o.Text = xmlOutline.Text
	o.Categories = categories
	o.Created = created

	return nil
}
