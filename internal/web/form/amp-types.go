package form

type ampType struct {
	Name     string
	Selected bool
}

func GetAmpTypes(selected int) []ampType {
	res := []ampType{
		{Name: "PP 6L6"},
		{Name: "PP EL34"},
		{Name: "SE 6L6"},
		{Name: "SE EL34"},
		{Name: "AMT TC-3"},
		{Name: "CALIF"},
		{Name: "BRIT M"},
		{Name: "BRIT L"},
		{Name: "DEFAULT"},
		{Name: "CALIF MODERN"},
		{Name: "CALIF VINTAGE"},
		{Name: "PVH 01"},
		{Name: "PVH 02"},
		{Name: "PVH 03"},
		{Name: "PVH 04"},
	}

	//idx, err := strconv.ParseUint(string(byte(selected)), 16, 16)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	if int(selected) < len(res) {
		res[selected].Selected = true
	}
	return res
}
