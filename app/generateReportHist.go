package app

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func generateReportHist(data interruptData) (string, error) {
	p := plot.New()
	p.Title.Text = "Interrupt Frequency"

	pxys := make(plotter.XYs, len(data.triggerTimes))

	for i, triggerTime := range data.triggerTimes {
		pxys[i].X = float64(i)
		pxys[i].Y = triggerTime
	}
	s, err := plotter.NewScatter(pxys)
	if err != nil {
		return "", err
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	s, err = plotter.NewScatter(plotter.XYs{
		{X: 1, Y: float64(data.smallestInterrupt)},
	})
	if err != nil {
		return "", err
	}
	s.GlyphStyle.Shape = draw.PyramidGlyph{}
	s.Color = color.RGBA{B: 255, A: 255}
	p.Add(s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("%s/frequency-%s.png", data.folderPath, data.sessionName)); err != nil {
		return "", err
	}
	return fmt.Sprintf("frequency-%s.png", data.sessionName), nil
}
