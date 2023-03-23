package Link

import (
	"fmt"
	"testing"
)

func TestSingleLink(t *testing.T) {
	var sLink = NewSingleLink()
	sLink.Append(1)
	sLink.Append(2)
	sLink.Append(3)
	sLink.Append(4)
	sLink.Append(5)
	sLink.InsertFirst(6)
	sLink.InsertFirst(7)
	sLink.Print()
	fmt.Println(sLink.IsEmpty())

}
