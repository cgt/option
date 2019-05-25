package option

type None struct{}

func (None) IsNone() bool         { return true }
func (None) IsSome() bool         { return false }
func (None) Value() string        { return "NONE" }
func (None) Map(f MapFunc) Option { return None{} }

type Some struct {
	value string
}

func (Some) IsNone() bool       { return false }
func (Some) IsSome() bool       { return true }
func (some Some) Value() string { return some.value }
func (some Some) Map(f MapFunc) Option {
	return Some{f(some.value)}
}

type Option interface {
	IsNone() bool
	IsSome() bool
	Value() string
	Map(MapFunc) Option
}

type MapFunc func(string) string
