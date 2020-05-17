package main

import (
	"fmt"
	"processi/frasi"
	"sort"
)

var testo = `Febbraio 2020
CONDIZIONI GENERALI DELL’OFFERTA TIM UNICA – PROPOSTA DI ATTIVAZIONE LINEA FISSA PRINCIPALE
A) Caratteristiche
L’iniziativa TIM Unica è gratuita e riservata ai clienti TIM di rete fissa che abbiano domiciliato i
pagamenti nella fattura TIM e consente, mediante il servizio “TIM Ricarica Automatica”, attivato su
tutte le linee mobili facenti parte di TIM Unica, di addebitare, sulla medesima fattura di rete fissa, gli
importi delle offerte attive sulle stesse linee mobili.
Ogni linea mobile facente parte di TIM Unica, inoltre, riceverà ogni mese un bonus di Giga illimitati,
aggiuntivi rispetto all’offerta dati già attiva sulla linea mobile. I GB sono validi in Italia di cui 6
GB/mese fruibili anche nei Paesi UE. Superati i 6 GB/mese, il cliente continuerà a navigare con i GB
eventualmente disponibili sulla propria linea mobile e se utilizzabili per navigare nei paesi UE. In
caso contrario il traffico dati si bloccherà fino alla successiva disponibilità mensile.
I GB illimitati sono soggetti esclusivamente ad uso personale. Il cliente è tenuto ad utilizzarli in modo
lecito, secondo buona fede e correttezza e deve astenersi, pertanto, dal conseguire vantaggi diversi da
quelli connessi alla normale fruizione per comunicazione interpersonale; per maggiori informazioni
sull’uso personale si rinvia alle Norme d’uso del servizio prepagato, disponibili e consultabili sul sito
TIM o nei negozi TIM.
TIM Unica si rinnova ogni mese gratuitamente e potrà essere disattivata senza alcun costo in qualsiasi
momento chiamando il 187, oppure accedendo alla sezione MyTIM del sito TIM o dal proprio
apparato attraverso l’App MyTIM.
B) Condizioni generali
La modalità di addebito degli importi delle offerte attive sulle linee mobili di TIM Unica avviene
mediante attivazione, sulle medesime, del servizio “TIM Ricarica Automatica”.
Tale servizio prevede l’erogazione automatica di una ricarica, che viene addebitata sulla fattura della
linea di rete fissa indicata in fase di attivazione di TIM Unica al verificarsi di una delle seguenti
condizioni: 1) “Ricarica Automatica su base rinnovo offerta”: entro le 24 ore precedenti il rinnovo di
ogni offerta attiva sulla linea mobile, verrà erogata una ricarica di importo pari al costo del relativo
canone. Sono escluse le offerte che hanno il costo di rinnovo inferiore a 2€ e le offerte rateizzate per
l’acquisto di un prodotto e 2) “Ricarica Automatica su base soglia”: viene erogata una ricarica di 5
euro quando il credito telefonico residuo scende sotto la soglia di 3 euro, fino ad un massimo di 20
euro mensili per linea. Superata tale soglia le linee mobili dovranno provvedere autonomamente a
ricaricare il proprio credito telefonico.
Possono far parte di TIM Unica:
- fino a 6 linee mobili, di cui una linea mobile deve essere necessariamente intestata al titolare della
linea fissa e rappresenterà la “linea principale”;
- le altre linee mobili, indicate come “linee aggiuntive”, possono essere intestate ad un soggetto
diverso dal titolare della linea mobile principale.
Le linee mobili aggiuntive sono invitate dalla linea mobile principale a fare parte della TIM Unica e
per aderire le stesse dovranno manifestare esplicita accettazione.
Eventuali attivazioni di nuove offerte sulle linee mobili facenti parte di TIM Unica saranno anch’esse
ricaricate con il servizio TIM Ricarica Automatica ed addebitate nella fattura della linea fissa.
Ogni linea mobile può fare parte di una sola TIM Unica.
In qualsiasi momento l’intestatario della linea di rete fissa può procedere con l’esclusione di una o
più linee mobili aggiuntive da TIM Unica. In tal caso la linea mobile interessata riceverà una notifica
via SMS dell’avvenuta esclusione. In qualsiasi momento l’intestatario della linea mobile aggiuntiva
può decidere di non usufruire più di TIM Unica; tale scelta verrà notificata all’intestatario della
fattura (SMS sulla linea mobile principale e mail, se disponibile). In entrambi i casi, la linea mobile
che non è più parte di TIM Unica non avrà più attivo il servizio TIM Ricarica Automatica, gli importi
non saranno più addebitati nella fattura della linea di rete fissa e non beneficerà dei Giga 
Febbraio 2020
illimitati/mese aggiuntivi. Inoltre, poiché sarà ripristinata la modalità di pagamento in essere prima
dell’adesione a TIM Unica, la linea mobile dovrà provvedere autonomamente al pagamento dei
canoni delle proprie offerte attive.
Nel caso in cui, invece, è la linea mobile principale a non far più parte di TIM Unica mantenendo
attiva la linea di rete fissa:
- la TIM Unica continuerà a rimanere attiva se è presente un indirizzo mail, validato da TIM e dal
cliente, associato alla linea di rete fissa, ma non sarà possibile aggiungere altre linee mobili (anche se
non è stato raggiunto il numero massimo di linee previste);
- la TIM Unica cessa se non è presente una mail, validata da TIM e dal cliente, e con essa termineranno
i benefici specificati nel paragrafo A. In tal caso sulle offerte attive sulla linea mobile principale sarà
ripristinata la modalità di pagamento in essere prima dell’adesione a TIM Unica e la linea mobile
principale dovrà, quindi, provvedere autonomamente al pagamento dei canoni delle proprie offerte
attive.
Il cliente prende atto ed accetta che le ricariche verranno erogate a condizione che il cliente titolare
della linea di rete fissa assolva ai pagamenti delle relative fatture.
In caso di mancata erogazione della ricarica, al cliente della linea mobile verrà inviato un sms con
l’invito a verificare il proprio credito telefonico al fine di evitare la disattivazione delle offerte attive
sulla propria linea mobile.
Il cliente, aderendo a TIM Unica, dichiara di aver preso completa visione delle caratteristiche di TIM
Unica, disponibili anche presso i negozi TIM. I benefici della TIM Unica saranno disponibili entro
10 giorni dall’attivazione della stessa e TIM provvederà ad informare il cliente tramite l’invio di SMS.
I servizi radiomobili inclusi sono ad uso esclusivamente personale. Il cliente si impegna a non
utilizzare TIM Unica in violazione di qualsivoglia legge o regolamento o comunque in maniera
abusiva, in modo da trarne direttamente o indirettamente indebito profitto o in modo da arrecare a
TIM o a terzi ingiustificato danno.
Per visualizzare i dettagli di TIM Unica, i consumi del bonus di Giga illimitati e gestire le linee
aggiuntive facenti parte di TIM Unica è possibile consultare il sito TIM accedendo alla sezione
MyTIM o, dal proprio apparato, all’App MyTIM.
Tutti i prezzi sono IVA inclusa.
`

func main() {

	periodi := frasi.GetPeriodi(testo) //os.Args[1])
	for i, periodo := range periodi {
		fmt.Println(i, periodo)
	}

	fmt.Println()

	is := frasi.Importanti(testo) //os.Args[1])

	var valori []float64
	for _, v := range is {
		valori = append(valori, v)
	}

	sort.Float64s(valori)

	//sort.Slice(valori, func(i, j int) bool { return valori[i] > valori[j] })

	for i := 0; i < len(is); i++ {
		for k, v := range is {
			if v == valori[i] {
				fmt.Println(k, v)
			}
		}
	}

}
