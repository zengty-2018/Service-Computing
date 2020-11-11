package rxgo

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type filOperator struct {
	opFunc func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool)
}

func (filop filOperator) op(ctx context.Context, o *Observable) {
	// must hold defintion of flow resourcs here, such as chan etc., that is allocated when connected
	// this resurces may be changed when operation routine is running.
	in := o.pred.outflow
	out := o.outflow
	//fmt.Println(o.name, "operator in/out chan ", in, out)
	var wg sync.WaitGroup

	if o.computation {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-time.After(o.sample_time):
					if o.flip != nil {
						buffer, _ := o.flip.([]interface{})
						for _, v := range buffer {
							o.sendToFlow(ctx, v, out)
						}
						o.flip = nil
					}
				}
			}
		}()
	}

	go func() {
		end := false
		for x := range in {
			if end {
				continue
			}
			// can not pass a interface as parameter (pointer) to gorountion for it may change its value outside!
			xv := reflect.ValueOf(x)
			// send an error to stream if the flip not accept error
			if e, ok := x.(error); ok && !o.flip_accept_error {
				o.sendToFlow(ctx, e, out)
				continue
			}
			// scheduler
			switch threading := o.threading; threading {
			case ThreadingDefault:
				if filop.opFunc(ctx, o, xv, out) {
					end = true
				}
			case ThreadingIO:
				fallthrough
			case ThreadingComputing:
				wg.Add(1)
				go func() {
					defer wg.Done()
					if filop.opFunc(ctx, o, xv, out) {
						end = true
					}
				}()
			default:
			}
		}

		if o.flip != nil {
			buffer, _ := o.flip.([]interface{}) //通过断言实现类型转换
			for _, v := range buffer {
				o.sendToFlow(ctx, v, out)
			}
		}

		wg.Wait() //waiting all go-routines completed
		o.closeFlow(out)
	}()
}

func (parent *Observable) newFilterObservable(name string) (o *Observable) {
	//new Observable
	o = newObservable()
	o.Name = name

	//chain Observables
	parent.next = o
	o.pred = parent
	o.root = parent.root

	//set options
	o.buf_len = BufferLen
	return o
}

func (parent *Observable) Debounce(timeSpan time.Duration) (o *Observable) {
	o = parent.newFilterObservable("debounce")
	o.times = 0
	o.time_span = timeSpan

	o.operator = debounceOperator
	return o
}

var debounceOperator = filOperator{
	opFunc: func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool) {
		o.times++
		go func() {
			temp := o.times
			time.Sleep(o.time_span)
			select {
			case <-ctx.Done():
				return
			default:
				if temp == o.times && !end {
					end = o.sendToFlow(ctx, item.Interface(), out)
				}
			}
		}()
		return false
	},
}

func (parent *Observable) Distinct() (o *Observable) {
	o = parent.newFilterObservable("distinct")
	o.item_map = make(map[string]bool)

	o.operator = distinctOperator
	return o
}

var distinctOperator = filOperator{opFunc: func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool) {
	var str = fmt.Sprintf("%v", item)

	if o.item_map[str] {
		return
	}

	o.item_map[str] = true
	o.sendToFlow(ctx, item.Interface(), out)

	return
},
}

func (parent *Observable) ElementAt(pos int) (o *Observable) {
	o = parent.newFilterObservable("elementAt")
	o.times = 0
	o.item_num = pos
	o.operator = elementAtOperator
	return o
}

var elementAtOperator = filOperator{opFunc: func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool) {
	if o.times == o.item_num {
		end = o.sendToFlow(ctx, item.Interface(), out)
	}
	o.times++
	return
},
}

func (parent *Observable) First() (o *Observable) {
	o = parent.newFilterObservable("first")
	o.times = 0
	o.operator = firstOperator
	return o
}

var firstOperator = filOperator{opFunc: func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool) {
	if o.times == 0 {
		o.sendToFlow(ctx, item.Interface(), out)
	}
	o.times++

	return
},
}

func (parent *Observable) IgnoreElements() (o *Observable) {
	o = parent.newFilterObservable("ignoreElements")
	o.operator = ignoreElementsOperator
	return o
}

var ignoreElementsOperator = filOperator{opFunc: func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool) {
	return
},
}

func (parent *Observable) Last() (o *Observable) {
	o = parent.newTransformObservable("last")

	o.operator = lastOperator
	return o
}

var lastOperator = filOperator{func(ctx context.Context, o *Observable, x reflect.Value, out chan interface{}) (end bool) {
	var slice []interface{}
	o.flip = append(slice, x.Interface())
	return
},
}

func (parent *Observable) Sample(timespan time.Duration) (o *Observable) {
	o = parent.newTransformObservable("sample")
	o.computation = true
	o.sample_time = timespan
	o.operator = sampleOperator
	return o
}

var sampleOperator = filOperator{func(ctx context.Context, o *Observable, x reflect.Value, out chan interface{}) (end bool) {
	var slice []interface{}
	o.flip = append(slice, x.Interface())
	return
},
}

func (parent *Observable) Skip(num int) (o *Observable) {
	o = parent.newFilterObservable("skip")
	o.times = 0
	o.skip_num = num
	o.operator = skipOperator
	return o
}

var skipOperator = filOperator{func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool) {
	o.times++
	if o.times <= o.skip_num {
		return
	} else {
		o.sendToFlow(ctx, item.Interface(), out)
	}
	return
},
}

func (parent *Observable) SkipLast(num int) (o *Observable) {
	o = parent.newFilterObservable("skipLast")
	o.times = 0
	o.skip_last_num = num
	o.last = nil
	o.operator = skipLastOperator
	return o
}

var skipLastOperator = filOperator{func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool) {
	o.times++
	o.last = append(o.last, item.Interface())

	if o.times > o.skip_last_num {
		o.sendToFlow(ctx, o.last[0], out)
		o.last = o.last[1:]
		fmt.Println(o.last)
	}

	return false
},
}

func (parent *Observable) Take(num int) (o *Observable) {
	o = parent.newFilterObservable("take")
	o.times = 0
	o.take_num = num
	o.operator = takeOperator
	return o
}

var takeOperator = filOperator{func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool) {
	o.times++

	if o.times > o.take_num {
		return
	} else {
		o.sendToFlow(ctx, item.Interface(), out)
	}

	return false
},
}

func (parent *Observable) TakeLast(num int) (o *Observable) {
	o = parent.newFilterObservable("takeLast")
	o.times = 0
	o.take_last_num = num
	o.operator = takeLastOperator
	return o
}

var takeLastOperator = filOperator{func(ctx context.Context, o *Observable, item reflect.Value, out chan interface{}) (end bool) {
	o.last = append(o.last, item.Interface())

	if o.times < o.take_last_num {

		//o.sendToFlow(ctx, o.last[o.times], out)

	} else {
		o.last = o.last[1:]
	}
	//fmt.Println(o.last)

	o.times++

	o.flip = o.last

	return false
},
}
