package utils

type Board[T any] [][]T

func (b Board[T]) Height() int {
	return len(b)
}

func (b Board[T]) Surface() int {
	return b.Height() * b.Width()
}

func (b Board[T]) Width() int {
	return len(b[0])
}
