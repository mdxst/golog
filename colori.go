package golog

import (
	"hash/adler32"
	"strconv"
	"strings"
)

const (
	placeholderColoreInSettatoreColore = "X"            // nella str. sotto
	settatoreColore                    = "\x1b[38;5;Xm" // rimpiazzare x con un numero da 30 [nero]
	// a 97 [bianco]
	chiusuraColore         = "\033[0m"
	troncamentoTitoloCache = 12 // in caratteri, per la mappa "cacheColori"
	// [chiave]

	rosso  = "\033[31m"
	viola  = "\033[35m"
	giallo = "\033[33m"
	verde  = "\033[32m"
	ciano  = "\033[36m"
)

var cacheColori map[string]string // map[titolo cappato a X car]colore
// (stringa ansi escape sequence), dove X è "troncamentoTitoloCache"

// dato un titolo, lo tronca a X (troncamentoTitoloCache) caratteri, vede se
// c'è già (come chiave) nella mappa "cacheColori":
//   - se sì, ne tira fuori il valore (un colore già stabilito prima per lui,
//     randomicamente)
//   - sennò, generiamo un colore random e lo immagazziniamo in cache
func colorePerTitolo(t string) string {
	// primaditutto verifichiamo se la mappa è stata creata
	if cacheColori == nil {
		// è nil!! inizializziamola!
		cacheColori = make(map[string]string)
	}

	// cappiamo il titolo a X caratteri
	if len(t) > troncamentoTitoloCache {
		t = t[0:troncamentoTitoloCache] // TODO: testare!!! (out of
		//	bounds??)
	}

	// controlliamo che non ci sia già t come chiave
	if val, ok := cacheColori[t]; ok {
		// c'è già
		return val
	} else {
		c := coloreRandom(t) // generiamo
		cacheColori[t] = c   // mettiamo in cache
		return c             // restituiamo
	}
}

// RITORNA GIÀ UNA ESCAPE ANSI PRONTA!
//
// non proprio random.... a stesso seed corrisponderà sempre stesso coloreRandom()
func coloreRandom(seed string) string {
	return strings.ReplaceAll(settatoreColore,
		placeholderColoreInSettatoreColore,
		strconv.Itoa(nDaHash(hash(seed), 60, 254)),
	)
}

func hash(s string) uint32 {
	return adler32.Checksum([]byte(s))
}

func nDaHash(h uint32, min, max int) int {
	// operazioni a caso
	n := (int(h/8) << 3) * 2

	// facciamolo rientrare tra <min> e <max>
	return (n % max) + min
}
