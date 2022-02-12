package utils

type IdType interface{
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
	~int8 | ~int16 | ~int32 | ~int64 | ~int 
}

type IdInterface[T IdType] interface {
	ID() T
}

func LowerBound[T IdType](array []T, target T) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func LowerBoundId[T IdType, I IdInterface[T]](array []I, target T) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid].ID() >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
