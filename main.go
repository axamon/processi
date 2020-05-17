package main

import (
	"fmt"
	"os"
	"processi/frasi"
)

// var testo = `L'energia atomica di fusione deve essere usata!
// Ci sono molti tipi di energia atomica ma quella atomica di fusione è la migliore.
// Insomma fate come volete ma l'energia atomica di fusione ci deve essere!
// Solo con "energia atomica" ecco lo slogan, si potranno vincere le sfide future.
// `

func main() {

	periodi := frasi.GetPeriodi(os.Args[1])
	for i, periodo := range periodi {
		fmt.Println(i, periodo)
	}

	fmt.Println(frasi.Importanti(os.Args[1]))

}
