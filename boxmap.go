package gosudoku

var BoxMap = map[string][]string{
	"BOX_A1": {"BOX_A4", "Box_D1"},
	"BOX_A4": {"BOX_A1", "BOX_A7", "BOX_D4"},
	"BOX_A7": {"BOX_A4", "BOX_D7"},
	"BOX_D1": {"BOX_A1", "BOX_D4", "BOX_G1"},
	"BOX_D4": {"BOX_A4", "BOX_D1", "BOX_D7", "BOX_G4"},
	"BOX_D7": {"BOX_D4", "BOX_A7", "BOX_G7"},
	"BOX_G1": {"BOX_D1", "BOX_G4"},
	"BOX_G4": {"BOX_G1", "BOX_D4", "BOX_G7"},
	"BOX_G7": {"BOX_G4", "BOX_D7"},
}
