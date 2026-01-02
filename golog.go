package golog

import (
	"fmt"
	"os"
	"strings"
	"time"
	"runtime"
)

const lvlErrori = 0     // log di errore
const lvlAttenzione = 0 // log di attenzione
const lvlInfo = 1       // log di errore
const lvlDebug = 3      // log di attenzione
const lvlSmart = 1      // LogS

var TitoloLogE = "ERRORE"      // log di errore
var TitoloLogW = "Attenzione!" // log di attenzione
var TitoloLogI = "Info"        // log di informazione
var TitoloLogD = "DEBUG"       // log di debug

var Prefisso = "-- LOG: " // esportato, personalizzabile
var SepTitoloEMsg = " > "
var Suffisso = " --"  // =
var SepMultiStr = " " // default: spazio
var LivelloMax = 0    // =. default: solo importanti (0). Sarebbe il massimo
// livello di cui log viene mostrato, dove 0 è il più importante, e il 3 debug

var SepDataOra = " | "
var StampaOra = true
var FormatoOra = "15:04:05"
var StampaData = true
var FormatoData = "02-01-2006"

// log comune (debug e non, lo decidi da "lvl", da 0 a 3)
func Log(lvl int, titolo string, msg ...string) string {
	return logga(titolo, colorePerTitolo(titolo), lvl, msg...)
}

// log di errore
func LogE(err error, quitta ...bool) string {
	if err != nil { // FIXME: ridondante (?)
		ret := logga(TitoloLogE, rosso, lvlErrori, err.Error())

		if len(quitta) > 1 {
			if quitta[0] == true {
				os.Exit(1)
			}
		}

		return ret
	}
	return ""
}

// log di attenzione
func LogW(msg ...string) string {
	return logga(TitoloLogW, giallo, lvlAttenzione, msg...)
}

// log di info
func LogI(msg ...string) string {
	return logga(TitoloLogI, verde, lvlInfo, msg...)
}

// log di DEBUG
func LogD(msg ...string) string {
	return logga(TitoloLogD, ciano, lvlDebug, msg...)
}

// log SMART: ottiene il nome della funzione chiamante e lo usa come TITOLO.
// - LIVELLO: lvlSmart (settato a 1, al pari di LogI)
// - VELOCITÀ: pesantuccio, lavora con gli array/slice/split... usare con
//	moderazione.
func LogS(msg ...string) string {
	// ottieni nome funzione chiamante
	pc, _, _, _ := runtime.Caller(1)
	nCompl := runtime.FuncForPC(pc).Name() // pacchetto.nomeFunc
	nComplSplit := strings.Split(nCompl, ".")
	var t string
	if len(nCompl) > 0 {
		t = nComplSplit[len(nComplSplit)-1]
	} else {
		t = "??"
	}
	return logga(t,
		colorePerTitolo(t),
		lvlSmart,
		msg...)
}

// prefissa e suffissa automaticamente, gestisce colori, comprende lvl
func logga(titolo, coloreTitolo string, lvl int, msg ...string) string {
	// msg è composto di più stringhe (variadico), sommiamole ora! Con
	// separatore: SepMultiStr
	var b strings.Builder
	nParole := len(msg)
	if nParole > 0 {
		for i := 0; i < nParole-1; i++ {
			b.WriteString(msg[i])
			b.WriteString(SepMultiStr)
		}
		b.WriteString(msg[nParole-1])
	}

	log := ""
	if lvl <= LivelloMax {
		log += Prefisso
		if StampaOra {
			log += fmt.Sprintf("%s%s", ora(), SepDataOra)
		}

		if StampaData {
			log += fmt.Sprintf("%s%s", data(), SepDataOra)
		}

		log += fmt.Sprint(coloreTitolo, titolo, chiusuraColore,
			SepTitoloEMsg, b.String(), Suffisso, "\n")
	}
	fmt.Print(log)
	return log
}

// ritorna (str) ora formattata come FormatoOra
func ora() string {
	return time.Now().Format(FormatoOra)
}

// ritorna (str) data formattata come FormatoData
func data() string {
	return time.Now().Format(FormatoData)
}
