package salut

{
	println("side effect of improting salut with the blank identifier as explicit package name")
}

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
