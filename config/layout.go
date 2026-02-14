package config

type Layout struct {
	Scale   float64
	OffsetX int
	OffsetY int
}

func NewLayout(winW, winH int) Layout {
	sx := float64(winW) / float64(WindowWidth)
	sy := float64(winH) / float64(WindowHeight)
	scale := sx
	if sy < sx {
		scale = sy
	}
	if scale < 1.0 {
		scale = 1.0
	}

	scaledW := int(float64(WindowWidth) * scale)
	scaledH := int(float64(WindowHeight) * scale)
	ox := (winW - scaledW) / 2
	oy := (winH - scaledH) / 2

	return Layout{
		Scale:   scale,
		OffsetX: ox,
		OffsetY: oy,
	}
}
