package golog

import (
	"fmt"
	"time"
)

const lvlErrori = 0     // log di errore
const lvlAttenzione = 0 // log di attenzione

var TitoloLogE = "ERRORE"      // log di errore
var TitoloLogW = "Attenzione!" // log di attenzione

var Prefisso = "-- LOG: " // esportato, personalizzabile
var SepTitoloEMsg = ">"
var Suffisso = "--" // =
var LivelloMax = 0  // =. default: solo importanti (0). Sarebbe il massimo
// livello di cui log viene mostrato, dove 0 è il più importante, e il 3 debug

var SepDataOra = " | "
var StampaOra = true
var FormatoOra = "15:04:05"
var StampaData = true
var FormatoData = "02-01-2006"

// log comune (debug e non, lo decidi da "lvl", da 0 a 3)
func Log(titolo string, msg string, lvl int) {
	logga(titolo, msg, colorePerTitolo(titolo), lvl)
}

// log di errore
func LogE(err error) {
	if err != nil { // FIXME: ridondante (?)
		logga(TitoloLogE, err.Error(), Rosso, 0)
	}
}

// log di attenzione
func LogW(msg string) {
	logga(TitoloLogW, msg, Giallo, 0)
}

// prefissa e suffissa automaticamente, gestisce colori, comprende lvl
func logga(titolo string, msg string, coloreTitolo string, lvl int) {
	if lvl > LivelloMax {
		if StampaOra {
			fmt.Printf("%s%s", ora(), SepDataOra)
		}

		if StampaData {
			fmt.Printf("%s%s", data(), SepDataOra)
		}

		fmt.Println(Prefisso, coloreTitolo, titolo, chiusuraColore,
			SepTitoloEMsg, msg, Suffisso)
	}
}

// ritorna (str) ora formattata come FormatoOra
func ora() string {
	return time.Now().Format(FormatoOra)
}

// ritorna (str) data formattata come FormatoData
func data() string {
	return time.Now().Format(FormatoData)
}
