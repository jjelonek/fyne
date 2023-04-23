package canvas

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/driver/common"
)

// Declare conformity with CanvasObject interface
var _ fyne.CanvasObject = (*Rectangle)(nil)

// Rectangle describes a colored rectangle primitive in a Fyne canvas
type Rectangle struct {
	baseObject

	FillColor   color.Color // The rectangle fill color
	StrokeColor color.Color // The rectangle stroke color
	StrokeWidth float32     // The stroke width of the rectangle
}

// Hide will set this rectangle to not be visible
func (r *Rectangle) Hide() {
	r.baseObject.Hide()

	common.Repaint(r)
}

// Move the rectangle to a new position, relative to its parent / canvas
func (r *Rectangle) Move(pos fyne.Position) {
	r.baseObject.Move(pos)

	common.Repaint(r)
}

// Refresh causes this rectangle to be redrawn with its configured state.
func (r *Rectangle) Refresh() {
	Refresh(r)
}

// Resize on a rectangle updates the new size of this object.
// If it has a stroke width this will cause it to Refresh.
func (r *Rectangle) Resize(s fyne.Size) {
	if s == r.Size() {
		return
	}

	r.baseObject.Resize(s)
	if r.StrokeWidth == 0 {
		return
	}

	Refresh(r)
}

// NewRectangle returns a new Rectangle instance
func NewRectangle(color color.Color) *Rectangle {
	return &Rectangle{
		FillColor: color,
	}
}
