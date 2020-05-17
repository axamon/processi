package frasi

import "strings"

// Importanti restituisce le frasi utilizzate più volte nel testo.
func Importanti(testo string) map[string]float64 {

	// indice di significatività relativa dei segmenti.
	var IS = make(map[string]float64)

	str := re.ReplaceAllString(testo, " ")

	parole := strings.Split(str, " ")

	// calcolo frequenze delle parole singole
	var frequenze = make(map[string]int)

	for _, parola := range parole {
		frequenze[parola]++
	}

	// calcolo delle corrispondenze di gruppi di parole nel testo
	var importanti = make(map[string]int)

	for i := 0; i < len(parole); i++ {
		for j := 0; j < len(parole); j++ {
			if i == j {
				break
			}
			if j+3 < len(parole) && i+3 < len(parole) && parole[i] == parole[j] && parole[i+1] == parole[j+1] && parole[i+2] == parole[j+2] && parole[i+3] == parole[j+3] {
				importanti[parole[i]+" "+parole[i+1]+" "+parole[i+2]+" "+parole[i+3]]++
				break
			}
			if j+2 < len(parole) && i+2 < len(parole) && parole[i] == parole[j] && parole[i+1] == parole[j+1] && parole[i+2] == parole[j+2] {
				importanti[parole[i]+" "+parole[i+1]+" "+parole[i+2]]++
				break
			}
			if j+1 < len(parole) && i+1 < len(parole) && parole[i] == parole[j] && parole[i+1] == parole[j+1] {
				importanti[parole[i]+" "+parole[i+1]]++
			}

		}

		// calcolo indici di significatività delle occorrenze importanti
		for segmento, freqSegmento := range importanti {
			elementi := strings.Split(segmento, " ")
			for _, elemento := range elementi {
				IS[segmento] += float64(frequenze[elemento] / freqSegmento)
			}
			IS[segmento] /= float64(len(segmento)) // ! da cambiare
		}

	}

	return IS

}
