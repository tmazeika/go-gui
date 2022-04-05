package elements

import (
	"github.com/veandco/go-sdl2/sdl"
)

type FlexFit int

const (
	FlexTight FlexFit = iota
	FlexLoose
)

type Flex struct {
	Axis         Axis
	Distribution Distribution
	Alignment    Alignment
	Fit          FlexFit
	Children     []Box
	parentData
}

func (b *Flex) GetSize(c Constraints) Size {
	toChildConstraint := func(c Constraints) Constraints {
		if b.Axis == AxisX {
			return c
		}
		return c.Flip()
	}
	fromChildSize := func(s Size) Size {
		if b.Axis == AxisX {
			return s
		}
		return s.Flip()
	}
	if b.Axis == AxisY {
		c = c.Flip()
	}
	usedWidth := Length(0)
	mostHeight := Length(0)
	cc := c.UnboundedWidth().NoMin()
	if b.Alignment == AlignStretch {
		cc = cc.TightenHeightToMax()
	}
	for _, child := range b.Children {
		if item, ok := child.(*FlexItem); ok && item.Factor != 0 {
			continue
		}
		size := fromChildSize(child.GetSize(toChildConstraint(cc)))
		child.SetParentData(size)
		usedWidth += size.Width
		mostHeight = max(mostHeight, size.Height)
	}
	factorSum := float32(0)
	for _, child := range b.Children {
		if item, ok := child.(*FlexItem); ok && item.Factor != 0 {
			factorSum += item.Factor
		}
	}
	space := float32(c.Max.Width - usedWidth)
	for _, child := range b.Children {
		item, ok := child.(*FlexItem)
		if !ok || item.Factor == 0 {
			continue
		}
		width := Length(item.Factor / factorSum * space)
		cc := NewConstraints(0, 0, width, c.Max.Height)
		if item.Fit == FlexTight {
			cc = cc.TightenWidthToMax()
		}
		if b.Alignment == AlignStretch {
			cc = cc.TightenHeightToMax()
		}
		size := fromChildSize(child.GetSize(toChildConstraint(cc)))
		child.SetParentData(size)
		if b.Fit == FlexLoose {
			usedWidth += width
		} else {
			usedWidth += cc.Max.Width
		}
		mostHeight = max(mostHeight, size.Height)
	}
	size := NewSize(usedWidth, mostHeight).GrowToSatisfy(c)
	if b.Axis == AxisY {
		size = size.Flip()
	}
	return size
}

func (b *Flex) Draw(g *sdl.Renderer, r Rect) {
	if b.Axis == AxisY {
		r = r.Flip()
	}
	left := r.Left
	top := r.Top
	childrenWidth := Length(0)
	for _, child := range b.Children {
		size := child.GetParentData().(Size)
		childrenWidth += size.Width
	}
	spaceAroundUnit := Length(0)
	spaceBetweenUnit := Length(0)
	spaceEvenlyUnit := Length(0)
	switch b.Distribution {
	case DistributeStart:
	case DistributeCenter:
		left += (r.Width - childrenWidth) / 2
	case DistributeEnd:
		left = r.Right() - childrenWidth
	case DistributeSpaceAround:
		if len(b.Children) > 0 {
			spaceAroundUnit = (r.Width - childrenWidth) / Length(len(b.Children)*2)
			left += spaceAroundUnit
		}
	case DistributeSpaceBetween:
		if len(b.Children) > 1 {
			spaceBetweenUnit = (r.Width - childrenWidth) / Length(len(b.Children)-1)
		}
	case DistributeSpaceEvenly:
		if len(b.Children) > 0 {
			spaceEvenlyUnit = (r.Width - childrenWidth) / Length(len(b.Children)+1)
			left += spaceEvenlyUnit
		}
	}
	for _, child := range b.Children {
		size := child.GetParentData().(Size)
		switch b.Alignment {
		case AlignStart:
		case AlignCenter:
			top = r.Top + (r.Height-size.Height)/2
		case AlignEnd:
			top = r.Top + r.Height - size.Height
		}
		rect := NewRect(left, top, size.Width, size.Height)
		if b.Axis == AxisY {
			rect = rect.Flip()
		}
		child.Draw(g, rect)
		left += size.Width + spaceAroundUnit*2 + spaceBetweenUnit + spaceEvenlyUnit
	}
}
