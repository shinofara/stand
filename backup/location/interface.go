package location

type Location interface {
	Save(string, string) error
	Clean() error
}
