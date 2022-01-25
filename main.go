package main

import "example.com/rivercrossing/state"

// Det er satt restriksjoner ved if-setninger. Rekkefølgen i main-funksjonen er riktig.
// Forsøk og endre eller kommenter ut felt, Da vil du ikke klare å komme i mål.

func main() {
	state.MoveTo("Restart")
	state.MoveTo("HS og kylling til båt")
	state.MoveTo("HS og kylling til Øst")
	state.MoveTo("HS til båt")
	state.MoveTo("HS til Vest")
	state.MoveTo("HS og korn til båt")
	state.MoveTo("HS og korn til Øst")

	// eksempel på feilkommando - fjern "//" i linje 19 for å se.
	// state.MoveTo("Demo")
}
