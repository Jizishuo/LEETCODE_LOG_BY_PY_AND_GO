package Travel_termino_1436

func destCity(paths [][]string) string {
	f := make(map[string]bool)
	e := make(map[string]bool)
	for _, path := range paths {
		f[path[0]] = true
		e[path[1]] = true
	}

	for key, _ := range e {
		if _, ok := f[key]; !ok {
			return key
		}
	}
	return ""
}