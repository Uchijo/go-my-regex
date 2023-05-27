package eval

type RegExp interface {
	Match(s *string) bool
}
