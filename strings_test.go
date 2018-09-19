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
}
