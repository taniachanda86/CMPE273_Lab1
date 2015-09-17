package fibonacci

import "testing"


type fibTestPair struct{
	input int
	output int
}

var fibTests = []fibTestPair{
	{1,1},
	{2,1},
	{3,2},
	{4,3},
	{5,5},
	{6,8},
	{7,13},
	{10,55},
	{15,610},
}

func TestFibonacci(t *testing.T){
	for _, i:= range fibTests{
		actual := Fibonacci(i.input)
		if actual != i.output{
			t.Error(
					"For ", i.input,
					"expected is: ", i.output,
					"got ", actual,
				)
		}
	}
}