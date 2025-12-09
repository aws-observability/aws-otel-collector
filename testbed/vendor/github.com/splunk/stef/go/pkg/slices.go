package pkg

func EnsureLen[T byte | int64 | float64 | any](data []T, dataLen int) []T {
	if cap(data) < dataLen {
		return append(data[0:cap(data)], make([]T, dataLen-cap(data))...)
	}
	return data[0:dataLen]
}
