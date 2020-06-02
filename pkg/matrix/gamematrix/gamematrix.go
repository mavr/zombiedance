package gamematrix

import (
	"github.com/mavr/totalwar/pkg/gameplay"
	"github.com/mavr/totalwar/pkg/matrix"
	"github.com/mavr/totalwar/pkg/matrix/simplematrix"
)

func New(p gameplay.Gameplay) matrix.Matrix {
	return simplematrix.NewMatrix(p) 
}