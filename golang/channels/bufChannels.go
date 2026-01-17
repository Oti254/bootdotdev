package main

func addEmailsToQueue(email []string) chan string {
	ch := make(chan string, len(email))
	for _, mail := range email {
		ch <- mail
	}
	return ch
}
