package model

import "fmt"

type Tag struct {
	Typ  int
	Name string
}

func (b *Tag) String() string {
	return fmt.Sprintf("%s", b.Name)
}
