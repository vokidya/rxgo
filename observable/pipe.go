package observable

type UnaryFunction func(source Observable[any]) Observable[any]

func Pipe[T any, R any](source Observable[T], operator func(source Observable[T]) Observable[R]) Observable[R] {
	return operator(source)
}

func Pipe2[T any, R1 any, R2 any](source Observable[T], operator1 func(source Observable[T]) Observable[R1], operator2 func(source Observable[R1]) Observable[R2]) Observable[R2] {
	return operator2(operator1(source))
}

func Pipe3[T any, R1 any, R2 any, R3 any](
	source Observable[T],
	operator1 func(source Observable[T]) Observable[R1],
	operator2 func(source Observable[R1]) Observable[R2],
	operator3 func(source Observable[R2]) Observable[R3],
) Observable[R3] {
	return operator3(Pipe2(source, operator1, operator2))
}

func Pipe4[T any, R1 any, R2 any, R3 any, R4 any](
	source Observable[T],
	operator1 func(source Observable[T]) Observable[R1],
	operator2 func(source Observable[R1]) Observable[R2],
	operator3 func(source Observable[R2]) Observable[R3],
	operator4 func(source Observable[R3]) Observable[R4],
) Observable[R4] {
	return operator4(Pipe3(source, operator1, operator2, operator3))
}

func Pipe5[T any, R1 any, R2 any, R3 any, R4 any, R5 any](
	source Observable[T],
	operator1 func(source Observable[T]) Observable[R1],
	operator2 func(source Observable[R1]) Observable[R2],
	operator3 func(source Observable[R2]) Observable[R3],
	operator4 func(source Observable[R3]) Observable[R4],
	operator5 func(source Observable[R4]) Observable[R5],
) Observable[R5] {
	return operator5(Pipe4(source, operator1, operator2, operator3, operator4))
}
