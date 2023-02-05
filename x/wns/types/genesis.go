package types

func ValidateGenesis(*GenesisState) error {
	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}
