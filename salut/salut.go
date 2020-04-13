package salut

// side effect of improting salut with the blank identifier as explicit package name
var RandomNumber = 123456

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
