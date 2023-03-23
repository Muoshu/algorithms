package Link

import (
	"fmt"
	"testing"
)

func TestDoubleLink(t *testing.T) {
	var dl = NewDoubleLink()
	dl.Append(1)
	dl.Append(2)
	dl.Append(3)
	dl.Append(4)
	d := dl.GetSize()
	fmt.Println(d)
	dl.Print()

}
