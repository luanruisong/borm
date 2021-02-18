package stringx

func SnakeName(name string) string {
	oldName := []rune(name)
	newName := make([]rune, 0)
	for i, v := range oldName {
		x := v
		if v < 91 {
			x += 32
			if i > 0 {
				newName = append(newName, 95)
			}
		}
		newName = append(newName, x)
	}
	return string(newName)
}
