package model

type student struct {
	Name  string
	Score float64
}

func Newstudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}

func (s *student) Getscore() float64 {
	return s.Score
}
