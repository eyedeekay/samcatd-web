// +build webface

package samcatweb

import (
    "testing"
)

func TestStringsLib(t *testing.T){
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
}
