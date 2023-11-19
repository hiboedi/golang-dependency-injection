package simple

type Config struct {
	Name string
}

type Appliacation struct {
	*Config
}

func NewApplication() *Appliacation {
	return &Appliacation{
		Config: &Config{
			Name: "Programmer Zaman Now",
		},
	}
}
