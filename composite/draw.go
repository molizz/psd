package composite

import (
	"image"
	"image/color"

	"golang.org/x/image/draw"

	"github.com/molizz/psd"
	"github.com/molizz/psd/blend"
)

var blendModes = map[psd.BlendMode]blend.Drawer{
	psd.BlendModeNormal:       blend.Normal,
	psd.BlendModeDarken:       blend.Darken,
	psd.BlendModeMultiply:     blend.Multiply,
	psd.BlendModeColorBurn:    blend.ColorBurn,
	psd.BlendModeLinearBurn:   blend.LinearBurn,
	psd.BlendModeDarkerColor:  blend.DarkerColor,
	psd.BlendModeLighten:      blend.Lighten,
	psd.BlendModeScreen:       blend.Screen,
	psd.BlendModeColorDodge:   blend.ColorDodge,
	psd.BlendModeLinearDodge:  blend.Add,
	psd.BlendModeLighterColor: blend.LighterColor,
	psd.BlendModeOverlay:      blend.Overlay,
	psd.BlendModeSoftLight:    blend.SoftLight,
	psd.BlendModeHardLight:    blend.HardLight,
	psd.BlendModeLinearLight:  blend.LinearLight,
	psd.BlendModeVividLight:   blend.VividLight,
	psd.BlendModePinLight:     blend.PinLight,
	psd.BlendModeHardMix:      blend.HardMix,
	psd.BlendModeDifference:   blend.Difference,
	psd.BlendModeExclusion:    blend.Exclusion,
	psd.BlendModeSubtract:     blend.Subtract,
	psd.BlendModeDivide:       blend.Divide,
	psd.BlendModeHue:          blend.Hue,
	psd.BlendModeSaturation:   blend.Saturation,
	psd.BlendModeColor:        blend.Color,
	psd.BlendModeLuminosity:   blend.Luminosity,
}

func drawWithOpacity(dst draw.Image, r image.Rectangle, src draw.Image, sp image.Point, opacity int, bm psd.BlendMode) {
	blendModes[bm].DrawMask(dst, r, src, sp, image.NewUniform(color.Alpha{uint8(opacity)}), image.Point{})
}

func removeAlpha(b *image.NRGBA) {
	d := b.Pix
	ln := len(d)
	_ = d[ln-1]
	for i := 0; i < ln; i += 4 {
		d[i+3] = 255
	}
}
