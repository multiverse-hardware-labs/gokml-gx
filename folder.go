package gokml_gx

import (
	"encoding/xml"
)

/*
 * Container type Folder
 */
type Folder struct {
	XMLName     xml.Name    `xml:"Folder"`
	Id          string      `xml:"id,attr,omitempty"`
	Name        string      `xml:"name,omitempty"`
	Visibility  int         `xml:"visibility,omitempty"`
	Open        int         `xml:"open,omitempty"`
	Address     string      `xml:"address,omitempty"`
	PhoneNumber string      `xml:"phoneNumber,omitempty"`
	Description string      `xml:"description,omitempty"`
	Placemarks  []Placemark `xml:"Placemark"`
	Folders     []Folder    `xml:"Folder"`
	Elements    []interface{}
}

/*
 * Folder represents the folder structure within a document or Kml structure
 */

// NewFolder() creates a new folder
func NewFolder(id, name string) Folder {
	f := Folder{Id: id, Name: name}
	return f
}

// AddPlacemark() adds a placemark to the folder
func (f *Folder) AddPlacemark(p Placemark) {
	f.Placemarks = append(f.Placemarks, p)
}

// AddFolder() adds a sub folder to this folder
func (d *Folder) AddFolder(f Folder) {
	d.Folders = append(d.Folders, f)
}

// SetVisibility() sets the visibility of the folder
func (f *Folder) SetVisibility(visibility bool) {
	f.Visibility = boolToInt(visibility)
}

// AddElement() adds a path to the document
func (f *Folder) AddElement(e interface{}) {

	f.Elements = append(f.Elements, e)
}
