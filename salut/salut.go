package salut

func Salut(gender int) string {
	switch gender {
	case 1:
		return "Monsieur"
	case 0:
		return "Madame"
	default:
		return ""
	}
}
