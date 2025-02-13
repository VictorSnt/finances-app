package entities

type ExpenseType int

const (
	FixedCost ExpenseType = iota
	Food
	Leisure
	Investment
	MedicalEmergency
)

func (e ExpenseType) String() string {
	namesSlice := [...]string{
		"Custo Fixo",
		"Alimentação",
		"Lazer",
		"Investimento",
		"Emergência Médica",
	}

	if int(e) > len(namesSlice) {
		return "N/A"
	}

	return namesSlice[e]
}
