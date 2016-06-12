package location

type Location interface {
	Save(string) error
	Clean() error
}
