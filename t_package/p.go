package tpackage

type student struct {
	name string
	age  int
}

func NewStudent(s string, age int) *student {
	return &student{
		s,
		age,
	}
}

func (s student) GetName() string {
	return s.name
}

func (s *student) SetName(name string) {
	s.name = name
}
