package test

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/bryanvaz/go-templ-lucide-icons"
)

var (
	snapshotSmile string
)

func TestRender(t *testing.T) {
	snapshotSmileBytes, err := os.ReadFile("./_snapshot_smile.svg")
	if err != nil {
		panic(err)
	}
	snapshotSmile = string(snapshotSmileBytes)

	t.Run("should render smile", func(t *testing.T) {
		sb := new(strings.Builder)
		err := icons.Smile().Render(context.Background(), sb)
		if err != nil {
			t.Fatalf("Error rendering icon: %v", err)
		}
		rendered := sb.String()

		if strings.TrimSpace(rendered) != strings.TrimSpace(snapshotSmile) {
			t.Fatalf("Render does not match snapshot: got '%v', want '%v'", rendered, snapshotSmile)
		}
	})
	t.Run("should render with default attributes", func(t *testing.T) {
		sb := new(strings.Builder)
		err := icons.Smile().Render(context.Background(), sb)
		if err != nil {
			t.Fatalf("Error rendering icon: %v", err)
		}
		rendered := sb.String()
		type SVG struct {
			XMLName        xml.Name `xml:"svg"`
			Xmlns          string   `xml:"xmlns,attr"`
			Width          string   `xml:"width,attr"`
			Height         string   `xml:"height,attr"`
			ViewBox        string   `xml:"viewBox,attr"`
			Fill           string   `xml:"fill,attr"`
			Stroke         string   `xml:"stroke,attr"`
			StrokeWidth    string   `xml:"stroke-width,attr"`
			StrokeLinecap  string   `xml:"stroke-linecap,attr"`
			StrokeLinejoin string   `xml:"stroke-linejoin,attr"`
		}
		var svg SVG
		err = xml.Unmarshal([]byte(rendered), &svg)
		if err != nil {
			t.Fatalf("Error unmarshaling icon: %v", err)
		}

		if svg.Xmlns != defaultAttributes["xmlns"] {
			t.Fatalf("Xmlns does not match: got %v, want %v", svg.Xmlns, defaultAttributes["xmlns"])
		}
		if svg.Width != defaultAttributes["width"] {
			t.Fatalf("Width does not match: got %v, want %v", svg.Width, defaultAttributes["width"])
		}
		if svg.Height != defaultAttributes["height"] {
			t.Fatalf("Height does not match: got %v, want %v", svg.Height, defaultAttributes["height"])
		}
		if svg.ViewBox != defaultAttributes["viewBox"] {
			t.Fatalf("ViewBox does not match: got %v, want %v", svg.ViewBox, defaultAttributes["viewBox"])
		}
		if svg.Fill != defaultAttributes["fill"] {
			t.Fatalf("Fill does not match: got %v, want %v", svg.Fill, defaultAttributes["fill"])
		}
		if svg.Stroke != defaultAttributes["stroke"] {
			t.Fatalf("Stroke does not match: got %v, want %v", svg.Stroke, defaultAttributes["stroke"])
		}
		if svg.StrokeWidth != defaultAttributes["stroke-width"] {
			t.Fatalf("StrokeWidth does not match: got %v, want %v", svg.StrokeWidth, defaultAttributes["stroke-width"])
		}
		if svg.StrokeLinecap != defaultAttributes["stroke-linecap"] {
			t.Fatalf("StrokeLinecap does not match: got %v, want %v", svg.StrokeLinecap, defaultAttributes["stroke-linecap"])
		}
		if svg.StrokeLinejoin != defaultAttributes["stroke-linejoin"] {
			t.Fatalf("StrokeLinejoin does not match: got %v, want %v", svg.StrokeLinejoin, defaultAttributes["stroke-linejoin"])
		}
	})
	t.Run("should adjust the size, stroke color and stroke width", func(t *testing.T) {
		sb := new(strings.Builder)
		err := icons.Smile(templ.Attributes{
			"size":         "48",
			"color":        "red",
			"stroke-width": "4",
		}).Render(context.Background(), sb)
		if err != nil {
			t.Fatalf("Error rendering icon: %v", err)
		}
		rendered := sb.String()
		t.Logf("rendered: %s", rendered)
		type SVG struct {
			XMLName     xml.Name `xml:"svg"`
			Width       string   `xml:"width,attr"`
			Height      string   `xml:"height,attr"`
			Stroke      string   `xml:"stroke,attr"`
			StrokeWidth string   `xml:"stroke-width,attr"`
		}
		var svg SVG
		err = xml.Unmarshal([]byte(rendered), &svg)
		if err != nil {
			t.Fatalf("Error unmarshaling icon: %v", err)
		}

		if svg.Width != "48" {
			t.Fatalf("Width does not match: got %v, want %v", svg.Width, "48")
		}
		if svg.Height != "48" {
			t.Fatalf("Height does not match: got %v, want %v", svg.Height, "48")
		}
		if svg.Stroke != "red" {
			t.Fatalf("Stroke does not match: got %v, want %v", svg.Stroke, "red")
		}
		if svg.StrokeWidth != "4" {
			t.Fatalf("StrokeWidth does not match: got %v, want %v", svg.StrokeWidth, "4")
		}
	})
	t.Run("should add a class to the element", func(t *testing.T) {
		sb := new(strings.Builder)
		err := icons.Smile(templ.Attributes{
			"class": "my-icon",
		}).Render(context.Background(), sb)
		if err != nil {
			t.Fatalf("Error rendering icon: %v", err)
		}
		rendered := sb.String()
		type SVG struct {
			XMLName xml.Name `xml:"svg"`
			Class   string   `xml:"class,attr"`
		}
		var svg SVG
		err = xml.Unmarshal([]byte(rendered), &svg)
		if err != nil {
			t.Fatalf("Error unmarshaling icon: %v", err)
		}
		classes := strings.Split(svg.Class, " ")
		found := false
		for _, class := range classes {
			if class == "my-icon" {
				found = true
			}
		}
		if !found {
			t.Fatalf("Class attribute does not have inserted class: got %v, want %v", svg.Class, "my-icon")
		}
	})
	t.Run("should pass children to the icon slot", func(t *testing.T) {
		helloText := "Hello, world!"
		contents := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			_, err := io.WriteString(w, fmt.Sprintf("<text>%s</text>", helloText))
			return err
		})
		ctx := templ.WithChildren(context.Background(), contents)
		sb := new(strings.Builder)
		err := icons.Smile().Render(ctx, sb)
		if err != nil {
			t.Fatalf("Error rendering icon: %v", err)
		}
		rendered := sb.String()
		type Node struct {
			XMLName xml.Name
			Attrs   []xml.Attr `xml:",any,attr"`
			Content []byte     `xml:",innerxml"`
			Nodes   []Node     `xml:",any"`
		}

		// Decode the XML
		decoder := xml.NewDecoder(strings.NewReader(rendered))
		var root Node
		if err := decoder.Decode(&root); err != nil {
			t.Fatalf("Error decoding XML: %v", err)
		}

		innerNodes := root.Nodes
		nodeFound := false
		for _, node := range innerNodes {
			if node.XMLName.Local == "text" {
				if string(node.Content) == helloText {
					nodeFound = true
					break
				}
			}
		}
		if !nodeFound {
			t.Fatalf("Node not found in XML: %v", innerNodes)
		}
	})
	t.Run("should render the alias icon", func(t *testing.T) {
		sb := new(strings.Builder)
		err := icons.Pen(templ.Attributes{
			"size":         "24",
			"color":        "red",
			"stroke-width": "4",
		}).Render(context.Background(), sb)
		if err != nil {
			t.Fatalf("Error rendering icon: %v", err)
		}
		type SVG struct {
			InnerXML string `xml:",innerxml"`
		}
		var penIcon SVG
		err = xml.Unmarshal([]byte(sb.String()), &penIcon)
		if err != nil {
			t.Fatalf("Error decoding XML: %v", err)
		}
		penIconHtml := penIcon.InnerXML

		sb = new(strings.Builder)
		err = icons.Edit2(templ.Attributes{
			"size":         "24",
			"color":        "red",
			"stroke-width": "4",
		}).Render(context.Background(), sb)
		if err != nil {
			t.Fatalf("Error rendering icon: %v", err)
		}
		var edit2Icon SVG
		err = xml.Unmarshal([]byte(sb.String()), &edit2Icon)
		if err != nil {
			t.Fatalf("Error decoding XML: %v", err)
		}
		edit2IconHtml := edit2Icon.InnerXML

		if penIconHtml != edit2IconHtml {
			t.Fatalf("Pen and Edit2 icons are not the same. Pen: %v, Edit2: %v", penIconHtml, edit2IconHtml)
		}
	})
	t.Run("should not scale the strokeWidth when absoluteStrokeWidth is set", func(t *testing.T) {
		sb := new(strings.Builder)
		err := icons.Smile(templ.Attributes{
			"size":                "48",
			"stroke":              "red",
			"absoluteStrokeWidth": true,
		}).Render(context.Background(), sb)
		if err != nil {
			t.Fatalf("Error rendering icon: %v", err)
		}
		rendered := sb.String()
		type SVG struct {
			XMLName     xml.Name `xml:"svg"`
			Stroke      string   `xml:"stroke,attr"`
			Width       string   `xml:"width,attr"`
			Height      string   `xml:"height,attr"`
			StrokeWidth string   `xml:"stroke-width,attr"`
		}
		var svg SVG
		err = xml.Unmarshal([]byte(rendered), &svg)
		if err != nil {
			t.Fatalf("Error unmarshaling icon: %v", err)
		}

		if svg.Stroke != "red" {
			t.Fatalf("Stroke does not match: got %v, want %v", svg.Stroke, "red")
		}
		if svg.Width != "48" {
			t.Fatalf("Width does not match: got %v, want %v", svg.Width, "48")
		}
		if svg.Height != "48" {
			t.Fatalf("Height does not match: got %v, want %v", svg.Height, "48")
		}
		if svg.StrokeWidth != "1" {
			t.Fatalf("StrokeWidth does not match: got %v, want %v", svg.StrokeWidth, "1")
		}
	})
}
