package fibonacci


func Fibonacci(x int) int{
	if x==0{
		return 0
	} else if x == 1{
		return 1
	} else {
		return Fibonacci(x-1) + Fibonacci(x-2)
	}
}