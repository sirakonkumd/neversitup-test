package smileService

func FindSmile(arr []string) int {
	face := 0
	if len(arr) == 0 {
		return face
	}
	for _, v := range arr {
		if (len(v) < 2 || len(v) > 3) ||
			(v[0] != ':' && v[0] != ';') ||
			(v[len(v)-1] != ')' && v[len(v)-1] != 'D') {
			continue
		}

		if len(v) == 3 {
			if v[1] != '-' && v[1] != '~' {
				continue
			}
		}
		face++
	}
	return face
}
