package main

func reformat(message string, formatter func(string) string) string {
	text := formatter(message)
	text = formatter(text)
	text = formatter(text)
	
	return "TEXTIO: " + text

}

