package state

import "testing"

// Brukt til å teste om MoveTo-funksjon fungerte som den skal, fil feile nå.
func Test1(t *testing.T) {
	wanted := "[V---\\[kylling rev korn mann]\\ \\________________________/ _____________________/---Ø]"
	state := MoveTo("vest")
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func Test2(t *testing.T) {
	wanted := "[V---\\_______________________\\ \\[kylling rev korn mann]/ ______________________/---Ø]"
	state := MoveTo("Båt")
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func Test3(t *testing.T) {
	wanted := "[V---\\_______________________\\ \\______________________/ [kylling rev korn mann]/---Ø]"
	state := MoveTo("Øst")
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
