
// Package weather fornece ferramentas para exibir as condições
// climáticas atuais e locais de várias cidades de Goblinocus.
package weather

var (
    // CurrentCondition armazenar a descricão do clima atual.
    CurrentCondition string
    // CurrentLocation  armazenaro nome da cidade atual.
	CurrentLocation  string
)

// Forecast define a localização e condição atual, retornando uma string formatada.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
