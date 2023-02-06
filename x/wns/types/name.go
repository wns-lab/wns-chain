package types

func (name *DomainName) Valid() bool {
	for _, c := range name.Name {
		if !(c >= 'a' && c <= 'z') || !(c >= 'A' && c <= 'Z') || !(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}
