package statistics

import "math"

func Media(valores []float64, qtd int) float64 {
	media := 0.0
	s := 0.0
	for i := 0; i < qtd; i++ {
		s += valores[i]
	}
	media = s / float64(qtd)

	return media
}

func DesvPadrao(valores []float64, qtd int) float64 {
	desvp := 0.0
	media := Media(valores, qtd)
	s := 0.0
	for i := 0; i < qtd; i++ {
		s += math.Pow((valores[i] - media), 2)
	}
	desvp = math.Sqrt(s / float64(qtd))

	return desvp
}
