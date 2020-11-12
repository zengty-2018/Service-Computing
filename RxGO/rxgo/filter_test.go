package rxgo

import (
	"testing"
	"time"

	"github.com/pmlpml/rxgo"
	"github.com/stretchr/testify/assert"
)

/*func TestDebounce(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 4, 5, 6).Map(func(x int) int {
		switch x {
		case 1:
			time.Sleep(0 * time.Millisecond)
		case 2:
			time.Sleep(11 * time.Millisecond)
		case 3:
			time.Sleep(15 * time.Millisecond)
		case 4:
			time.Sleep(5 * time.Millisecond)
		case 5:
			time.Sleep(12 * time.Millisecond)
		case 6:
			time.Sleep(2 * time.Millisecond)
		}
		return x
	}).Debounce(10 * time.Millisecond).Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{1, 2, 4}, res, "Debounce Test Error!")
}*/

func TestDistinct(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 1, 2, 4, 4, 5).Distinct().Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{1, 2, 3, 4, 5}, res, "Distinct Error!")
}

func TestElementAt(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 4, 5, 6).ElementAt(2).Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{3}, res, "ElementAt Error!")
}

func TestFirst(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 4, 5, 6).First().Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{1}, res, "First Error!")
}

func TestIgnoreElements(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 4, 5, 6).IgnoreElements().Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{}, res, "IgnoreElements Error!")
}

func TestLast(t *testing.T) {
	res := []int{}
	rxgo.Just(1, 2, 3, 4, 5, 6).Last().Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{6}, res, "Last Error")
}

func TestSample(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 4, 5, 6).Map(func(x int) int {
		switch x {
		case 1:
			time.Sleep(0 * time.Millisecond)
		case 2:
			time.Sleep(3 * time.Millisecond)
		case 3:
			time.Sleep(2 * time.Millisecond)
		case 4:
			time.Sleep(13 * time.Millisecond)
		case 5:
			time.Sleep(7 * time.Millisecond)
		case 6:
			time.Sleep(8 * time.Millisecond)
		}
		return x
	}).Sample(10 * time.Millisecond).Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{3, 4, 5, 6}, res, "Sample Error!")
}

func TestSkip(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 4, 5, 6).Skip(2).Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{3, 4, 5, 6}, res, "Skip Error!")
}

func TestSkipLast(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 4, 5, 6).SkipLast(3).Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{1, 2, 3}, res, "SkipLast Error!")
}

func TestTake(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 4, 5, 6).Take(2).Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{1, 2}, res, "Take Error!")
}

func TestTakeLast(t *testing.T) {
	res := []int{}
	Just(1, 2, 3, 4, 5, 6).TakeLast(2).Subscribe(func(x int) {
		res = append(res, x)
	})

	assert.Equal(t, []int{5, 6}, res, "TakeLast Error!")
}
