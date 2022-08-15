package app

import (
	"fmt"
	"image/color"

	"github.com/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func generateReportHist(data interruptData) (string, error) {
	p := plot.New()
	p.Title.Text = "Interrupt Frequency"
	v := make(plotter.Values, len(data.triggerTimes))

	for i := range v {
		v[i] = data.triggerTimes[i]
	}

	h, err := plotter.NewHist(v, len(v)+1)
	if err != nil {
		return "", err
	}

	h.Normalize(1)
	p.Add(h)

	norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(float64(len(v) + 1))
	p.Add(norm)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("%s/frequency-%s.png", data.folderPath, data.sessionName)); err != nil {
		return "", err
	}
	return fmt.Sprintf("frequency-%s.png", data.sessionName), nil
}
