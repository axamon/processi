package main

import (
	"fmt"
	"processi/frasi"
)

var testo = `L'energia atomica di fusione deve essere usata! 
Ci sono molti tipi di energia atomica ma quella atomica di fusione è la migliore.
Insomma fate come volete ma l'energia atomica di fusione ci deve essere!
Solo con "energia atomica" ecco lo slogan, si potranno vincere le sfide future.
`

func main() {

	periodi := frasi.GetPeriodi(testo)
	for i, periodo := range periodi {
		fmt.Println(i, periodo)
	}

	fmt.Println(frasi.Importanti(testo))

}
