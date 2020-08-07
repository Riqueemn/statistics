package statistics

import "math"

//Funções com um simbolo "*" logo acima da função significa que está em teste

func Soma(valores []float64, qtd int) float64 {
	soma := 0.0
	for i := 0; i < qtd; i++ {
		soma += valores[i]
	}

	return soma
}

func Media(valores []float64, qtd int) float64 {
	media := 0.0
	soma := Soma(valores, qtd)
	media = soma / float64(qtd)

	return media
}

func DesvioPadrao(valores []float64, qtd int) float64 {
	media := Media(valores, qtd)
	soma := Amount(valores, media, qtd)
	desvp := math.Sqrt(soma / float64(qtd))

	return desvp
}

func DesvioPadraoEM(valores []float64, media float64, qtd int) float64 {
	soma := Amount(valores, media, qtd)
	desvp := math.Sqrt(soma / float64(qtd))

	return desvp
}

func MaxDeTresValores(a float64, b float64, c float64) float64 {
	return math.Max(a, math.Max(b, c))
}

func Amount(valores []float64, x float64, qtd int) float64 {
	soma := 0.0
	for i := 0; i < qtd; i++ {
		j := valores[i] - x
		soma += math.Pow((j), 2)
	}

	return soma
}

func SomaPonderada(valores []float64, pesos []float64) float64 {
	soma := 0.0
	qtd := len(pesos)
	for i := 0; i < qtd; i++ {
		soma += valores[i] * pesos[i]
	}

	return soma
}

func MediaPonderada(valores []float64, pesos []float64) float64 {
	soma := SomaPonderada(valores, pesos)
	qtd := len(pesos)
	d := Soma(pesos, qtd)

	return soma / d
}

func MaiorValor(valores []float64) float64 {
	i := 0
	maior := valores[i]
	for i = 1; i < len(valores); i++ {
		if valores[i] > maior {
			maior = valores[i]
		}
	}

	return maior
}

func MenorValor(valores []float64) float64 {
	i := 0
	menor := valores[i]
	for i = 1; i < len(valores); i++ {
		if valores[i] < menor {
			menor = valores[i]
		}
	}

	return menor
}
