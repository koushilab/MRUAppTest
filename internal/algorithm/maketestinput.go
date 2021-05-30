package MRUApp

func MakeTestInput() *MRU {
	data := MRU{}
	addMRUData(&data)
	return &data
}

func addMRUData(data *MRU) {
	content := MRU{
		State: State{
			Model: "NewTextFile",
			Item: []Item{{
				Label: "file1",
				Payload: Payload{
					Content: "Data of the file goes here",
				},
			},
			},
		},
	}
	data.State = content.State
}
