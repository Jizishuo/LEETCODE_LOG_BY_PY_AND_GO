package Design_the_Goal_parser_1678

func interpret(command string) string {
	res := ""
	for i:=0;i<len(command);i++ {
		if command[i] == 'G' {
			res += "G"
		} else if command[i] == '(' {
			if command[i+1] == 'a' {
				res += "al"
				i += 3
			} else {
				res += "o"
				i += 1
			}
		}
	}
	return res
}