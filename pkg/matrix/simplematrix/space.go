package simplematrix

import (
	"math"

	"github.com/mavr/totalwar/pkg/matrix"
)

type Space struct {
	x matrix.WorldSpacePiece
	y matrix.WorldSpacePiece
}

func NewSpace(x, y matrix.WorldSpacePiece) *Space {
	return &Space{
		x: x,
		y: y,
	}
}

func (s *Space) X() matrix.WorldSpacePiece {
	return s.x
}

func (s *Space) Y() matrix.WorldSpacePiece {
	return s.y
}

func (s *Space) Distance(aim matrix.WorldSpace) matrix.WorldSpacePiece {
	return matrix.WorldSpacePiece(
		math.Sqrt(math.Pow(float64(s.x-aim.X()), 2) + math.Pow(float64(s.y-aim.Y()), 2)))
}

func (s *Space) Move(aim matrix.WorldSpace, speed matrix.WorldSpacePiece) {

	tg := math.Abs(float64((s.y - aim.Y()) / (s.x - aim.X())))
	rad := math.Atan(tg)

	a := float64(speed) * math.Sin(rad) / tg
	b := float64(speed) * math.Cos(rad) * tg

	if (s.x - aim.X()) > 0 {
		s.x = matrix.WorldSpacePiece(float64(s.x) - a)
	} else {
		s.x = matrix.WorldSpacePiece(float64(s.x) + a)
	}

	if (s.y - aim.Y()) > 0 {
		s.y = matrix.WorldSpacePiece(float64(s.y) - b)
	} else {
		s.y = matrix.WorldSpacePiece(float64(s.y) + b)
	}
}

func (s *Space) update(x, y matrix.WorldSpacePiece) {
	s.x = x
	s.y = y
}
