package cfi

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {

	c, err := From("EMXXXM")
	fmt.Println(err)
	fmt.Println(c.Format(Tag))
	fmt.Println(c.Format(Short))
	fmt.Println(c.Format(Long))

	fmt.Println()
	fmt.Println()
	fmt.Println()

	c = New(" qw12 ")
	fmt.Println(c.Format(Tag))
	fmt.Println(c.Format(Short))
	fmt.Println(c.Format(Long))
}
