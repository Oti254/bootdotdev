package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	var isVulgar bool
	for i, word := range msg {
		for _, bad := range badWords {
			if word == bad {
				isVulgar = true
				return i
			} else {
				isVulgar = false
			}
		}
	}
	if isVulgar == false {
		return -1
	}
	return 0

}
