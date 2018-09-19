// +build webface

package samcatweb

import (
	"testing"
)

func TestStringsLib(t *testing.T) {
	x := dedouble("////////", "//", "/")
	if x != "/" {
		t.Fatal(x)
	}
	y := name(" \n name=test \n")
	if y != "test" {
		t.Fatal(y)
	}
	z := stringify(&[]string{"test,,,test,,test,test,"})
	if z != "test,test,test,test" {
		t.Fatal(z)
	}
	a := makeid("///test,,,//test", "test_id")
	if a != "test_id_test_test" {
		t.Fatal(a)
	}
	b := makeclass("///test___//test", "test_id")
	if b != "test,id,test,test" {
		t.Fatal(b)
	}
	c := makeurl(",,,,test___,,test", "test_id")
	if c != "test/id/test/test" {
		t.Fatal(c)
	}
}
