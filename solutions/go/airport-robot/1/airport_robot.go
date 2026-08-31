package airportrobot

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

func SayHello (name string, greeter Greeter) string {
    return "I can speak " + greeter.LanguageName() + ": "  + greeter.Greet(name)
}

func (l Italian) LanguageName() string {
    return "Italian"
}

func (l Italian) Greet(name string) string {
    return "Ciao " + name + "!"
}

type Italian struct {
}

func (l Portuguese) LanguageName() string {
    return "Portuguese"
}

func (l Portuguese) Greet(name string) string {
    return "Olá " + name + "!"
}

type Portuguese struct {
}