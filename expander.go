package linq

var _ Expander[int] = &Expended[int, int]{}
var _ Expander3[int] = &Expended3[int, int, int]{}
var _ Expander4[int] = &Expended4[int, int, int, int]{}

//var _ Expended[int, int] = &Expended[int, int]{}
//var _ Expended3[int, int, int] = &Expended3[int, int, int]{}
//var _ Expended4[int, int, int, int] = &Expended4[int, int, int, int]{}
var _ OrderedExpander[int] = &orderedExtender[int, int]{}
var _ OrderedExpended[int, int] = &orderedExtender[int, int]{}

func (q QueryG[T]) Expend(e Expander[T]) Expander[T] {
	e.Expend(q)
	return e
}

func (q QueryG[T]) Expend3(e Expander3[T]) Expander3[T] {
	e.Expend(q)
	return e
}

func (q QueryG[T]) Expend4(e Expander4[T]) Expander4[T] {
	e.Expend(q)
	return e
}

func (q OrderedQueryG[T]) Expend(e OrderedExpander[T]) OrderedExpander[T] {
	e.Expend(q)
	return e
}

type OrderedExpander[T any] interface {
	Expend(q OrderedQueryG[T]) any
}

type Expander[T any] interface {
	Expend(q QueryG[T]) any
	Expander()
}

type Expander3[T any] interface {
	Expend(q QueryG[T]) any
	Expander3()
}

type Expander4[T any] interface {
	Expend(q QueryG[T]) any
	Expander4()
}

//type Expended[T1, T2 any] interface {
//	Select(selector func(T1) T2) QueryG[T2]
//	SelectIndexed(selector func(int, T1) T2) QueryG[T2]
//	SelectMany(selector func(T1) QueryG[T2]) QueryG[T2]
//	SelectManyIndexed(selector func(int, T1) QueryG[T2]) QueryG[T2]
//	DistinctBy(selector func(T1) T2) QueryG[T1]
//	OrderBy(selector func(T1) T2) OrderedQueryG[T1]
//	OrderByDescending(selector func(T1) T2) OrderedQueryG[T1]
//	ExceptBy(q QueryG[T1], selector func(T1) T2) QueryG[T1]
//	Expander()
//}
//
//type Expended3[T1, T2, T3 any] interface {
//	Zip(q2 QueryG[T2],
//		resultSelector func(T1, T2) T3) QueryG[T3]
//	SelectManyBy(selector func(T1) QueryG[T2],
//		resultSelector func(T2, T1) T3) QueryG[T3]
//	SelectManyByIndexed(selector func(int, T1) QueryG[T2],
//		resultSelector func(T2, T1) T3) QueryG[T3]
//	GroupBy(keySelector func(T1) T2,
//		elementSelector func(T1) T3) QueryG[GroupG[T2, T3]]
//	Expander3()
//}
//
//type Expended4[T1, T2, T3, T4 any] interface {
//	Join(inner QueryG[T2],
//		outerKeySelector func(T1) T3,
//		innerKeySelector func(T2) T3,
//		resultSelector func(outer T1, inner T2) T4) QueryG[T4]
//	GroupJoin(inner QueryG[T2],
//		outerKeySelector func(T1) T3,
//		innerKeySelector func(T2) T3,
//		resultSelector func(outer T1, inners []T2) T4) QueryG[T4]
//	Expander4()
//}

type OrderedExpended[T1 any, T2 comparable] interface {
	ThenBy(selector func(T1) T2) OrderedQueryG[T1]
	ThenByDescending(selector func(T1) T2) OrderedQueryG[T1]
}

type Expended[T1, T2 any] struct {
	q QueryG[T1]
}

func (*Expended[T1, T2]) Expander() {}

func (e *Expended[T1, T2]) Expend(q QueryG[T1]) any {
	e.q = q
	return e
}

func To2[T1, T2 any]() Expander[T1] {
	return &Expended[T1, T2]{}
}

type Expended3[T1, T2, T3 any] struct {
	q QueryG[T1]
}

func (*Expended3[T1, T2, T3]) Expander3() {}

func (e *Expended3[T1, T2, T3]) Expend(q QueryG[T1]) any {
	e.q = q
	return e
}

func To3[T1, T2, T3 any]() Expander3[T1] {
	return &Expended3[T1, T2, T3]{}
}

func To4[T1, T2, T3, T4 any]() Expander4[T1] {
	return &Expended4[T1, T2, T3, T4]{}
}

func OrderedTo2[T1 any, T2 comparable]() OrderedExpander[T1] {
	return &orderedExtender[T1, T2]{}
}

type orderedExtender[T1 any, T2 comparable] struct {
	q OrderedQueryG[T1]
}

func (o *orderedExtender[T1, T2]) Expend(q OrderedQueryG[T1]) any {
	o.q = q
	return o
}

type Expended4[T1, T2, T3, T4 any] struct {
	q QueryG[T1]
}

func (*Expended4[T1, T2, T3, T4]) Expander4() {}

func (e *Expended4[T1, T2, T3, T4]) Expend(q QueryG[T1]) any {
	e.q = q
	return e
}
