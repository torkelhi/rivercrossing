package state

import "fmt"

var HsState = [7]string{
	"[V---\\[kylling rev korn HS]\\ \\___________________/ _________________/---Ø]",
	"[V---\\_____[rev korn]______\\ \\___[kyling & HS]___/ _________________/---Ø]",
	"[V---\\_____[rev korn]______\\ \\___________________/ [kylling & HS]___/---Ø]",
	"[V---\\_____[rev korn]______\\ \\_______[HS]________/ [kylling]________/---Ø]",
	"[V---\\___[rev korn HS]_____\\ \\___________________/ [kylling]________/---Ø]",
	"[V---\\_______[rev]_________\\ \\____[HS & korn]____/ [kylling]________/---Ø]",
	"[V---\\_______[rev]_________\\ \\___________________/ [kylling korn HS]/---Ø]"}

var Status = HsState[0]

func MoveTo(MoveTo string) string {

	if MoveTo == "Restart" || MoveTo == "HS og kylling til Vest" {
		Status = HsState[0]
		fmt.Println(HsState[0] + "Init State")
		return HsState[0]
	}

	if MoveTo == "HS og kylling til båt" && (Status == HsState[0] || Status == HsState[2]) {
		Status = HsState[1]
		fmt.Println(HsState[1])
		return HsState[1]
	}

	if MoveTo == "HS og kylling til Øst" && (Status == HsState[1] || Status == HsState[3]) {
		Status = HsState[2]
		fmt.Println(HsState[2])
		return HsState[2]
	}

	if MoveTo == "HS til båt" && (Status == HsState[2] || Status == HsState[4]) {
		Status = HsState[3]
		fmt.Println(HsState[3])
		return HsState[3]
	}

	if MoveTo == "HS til Vest" && (Status == HsState[3] || Status == HsState[5]) {
		Status = HsState[4]
		fmt.Println(HsState[4])
		return HsState[4]
	}

	if MoveTo == "HS og korn til båt" && (Status == HsState[4] || Status == HsState[6]) {
		Status = HsState[5]
		fmt.Println(HsState[5])
		return HsState[5]
	}

	if MoveTo == "HS og korn til Øst" && Status == HsState[5] {
		Status = HsState[6]
		fmt.Println(HsState[6] + "Du fikk alle over trygt og forlot reven, godt jobbet")
		return HsState[6]
	}

	fmt.Println("HS (eventuelt og ting/dyr) til (båt/Vest/Øst) || Ulovlig handling")
	return ""
}
