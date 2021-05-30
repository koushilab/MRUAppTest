package MRUApp

func MakeTestInput() *MRU {
	data := MRU{}
	addMRUData()
	return &data
}

func addMRUData(data MRU) {
	content := MRU{
		State: State{
			Model: "NewTextFile",
			Item:  []Item{},
		},
	}
}
