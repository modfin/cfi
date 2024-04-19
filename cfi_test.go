package cfi

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {

	c, err := From("EPVNDB")
	fmt.Println(err)
	fmt.Println(c.Format(Tag))
	fmt.Println(c.Format(Short))
	fmt.Println(c.Format(Long))
}
