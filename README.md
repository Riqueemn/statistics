# statistics
Pacote com funções para a linguagem Go que retornam cálculos estatísticos (médias, desvio padrão, etc...)<br>
Criado com o intuito de compensar a precariedade de pacotes com cálculos estatísticos<br>
Fico feliz quem puder contribuir! :)

<br><br>
Ex.:<br>
    Para retornar a media de n valores<br>
    media := statistics.Media(valores, quantidade)


<br><br>
Lista de funções:
<br>
    Soma(valores []float64, qtd int) float64                    ->  Retorna a Somatóra dos valores no slice
<br>
    Media(valores []float64, qtd int) float64                   ->  "" Media dos valores no slice
<br>
    DesvioPadrao(valores []float64, qtd int) float64            ->  "" Desvio padrão dos valores no slice
<br>
    MaxDeTresValores(a float64, b float64, c float64) float64   ->  "" Máximo entre três valores
<br>
    Amount(valores []float64, x float64, qtd int) float64       ->  "" Soma dos quadrados dos valores do slice
<br>
    SomaPonderada(valores []float64, pesos []float64) float64   ->  "" Soma ponderada
<br>
    MediaPonderada(valores []float64, pesos []float64) float64  ->  "" Média ponderada
