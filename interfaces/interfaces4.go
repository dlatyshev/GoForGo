package interfaces

type Empty interface{}

type A struct{}
type B struct{}

func (a A) Method() {}
func (b B) Method() {}

func DescribeEmpty(e Empty) {
	switch e.(type) {
	case A:
		println("A")
	case B:
		println("B")
	default:
		println("Unknown type")
	}

}

func ShowInterfaces4Example() {
	var a A
	var b B

	DescribeEmpty(a)
	DescribeEmpty(b)
}
