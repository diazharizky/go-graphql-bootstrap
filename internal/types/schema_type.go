package types

type Schema string

func (s Schema) String() string {
	return string(s)
}