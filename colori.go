package golog

import (
	"math/rand"
	"strconv"
	"strings"
	"crypto/sha1"
)

const (
	placeholderColoreInSettatoreColore = "X"       // nella str. sotto
	settatoreColore                    = "\033[Xm" // rimpiazzare x con un numero da 30 [nero]
	// a 97 [bianco]
	chiusuraColore         = "\033[0m"
	troncamentoTitoloCache = 12 // in caratteri, per la mappa "cacheColori"
	// [chiave]

	Rosso = "\033[31m"
	Giallo = "\033[33m"
)

var cacheColori map[string]string // map[titolo cappato a X car]colore
// (stringa ansi escape sequence), dove X è "troncamentoTitoloCache"

// dato un titolo, lo tronca a X (troncamentoTitoloCache) caratteri, vede se
// c'è già (come chiave) nella mappa "cacheColori":
//   - se sì, ne tira fuori il valore (un colore già stabilito prima per lui,
//     randomicamente)
//   - sennò, generiamo un colore random e lo immagazziniamo in cache
func colorePerTitolo(t string) string {
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

func coloreRandom(seed string) string {
	return strings.ReplaceAll(settatoreColore,
		placeholderColoreInSettatoreColore,
		strconv.Itoa(numeroRandom(hash(seed), 30, 97)))
}

func hash(s string) int64 {
	daConvertire := sha1.Sum([]byte(s))
	ret := int64(0)
    	for _, b := range daConvertire {
        	ret = (ret << 8) | int64(b)
    	}
	return ret
}

func numeroRandom(seed int64, min int, max int) int {
	rand.Seed(seed)
	return min + rand.Intn(max-min)
}
