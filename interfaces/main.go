package main

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	englishBot := englishBot{}
	spanishBot := spanishBot{}	
	printGreeting(englishBot)
	printGreeting(spanishBot)
	
}

func printGreeting(b bot) {	
	greeting := b.getGreeting()
	println(greeting)
}


func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}