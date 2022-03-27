package form

// https://media.amt-sales.com/cp-100-fw-v-pa-6-8-3/
var Amplifiers = []string{
	"Двухтактный усилитель с лампами 6L6",
	"Двухтактный усилитель с лампами EL34",
	"Однотактный усилитель с лампами 6L6",
	"Однотактный усилитель с лампами EL34",
	"AMT Electronics TubeCake TC-3",
	"Mesa/Boogie Dual Rectifier",
	"Marshall JCM-800",
	"Laney LC50 II 112 Combo",
	"Линейная АЧХ",
	"Mesa/Boogie California Modern",
	"Mesa/Boogie California Vintage",
	"Peavey 6505 Presence=0 Resonance=0",
	"Peavey 6505 Presence=5 Resonance=5",
	"Peavey 6505 Presence=8 Resonance=7",
	"Peavey 6505 Presence=9 Resonance=8",
}

var AmplifiersShort = []string{
	"PP 6L6",
	"PP EL34",
	"SE 6L6",
	"SE EL34",
	"AMT TC-3",
	"CALIF",
	"BRIT M",
	"BRIT L",
	"DEFAULT",
	"CALIF MODERN",
	"CALIF VINTAGE",
	"PVH 01",
	"PVH 02",
	"PVH 03",
	"PVH 04",
}

func GetAmplifiers(version string) []string {
	return Amplifiers
}

func GetAmplifiersShort() []string {
	return AmplifiersShort
}
