package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 两个协程交替答应数字，直到100.
func sim2routines() {
	cs := make([]chan int, 2)
	end := make(chan bool)
	for i := range cs {
		cs[i] = make(chan int)
	}

	num := 0
	tn := 2

	for i := 0; i < tn; i++ {
		go func(i, o chan int) {
			for {
				a := <-i
				fmt.Printf("%d ", a)
				a += 1
				if a > 100 {
					end <- true
					fmt.Println()
					break
				}
				// <-time.After(1 * time.Second)
				o <- a

			}
		}(cs[i%tn], cs[(i+1)%tn])
	}

	cs[0] <- num

	<-end
}

// 生成质数列表
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			fmt.Printf("generating %d\n", i)
			ch <- i
		}
	}()

	fmt.Printf("gen ch %x\n", ch)
	return ch
}
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				fmt.Printf("prime check: %d\n", i)
				out <- i
			}
		}
	}()
	return out
}

// 累加数组中的所有元素
func add(numbers []int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}

// 并发累加数组中所有元素
func addConcurrent(goroutines int, numbers []int) int {
	s := int64(0)

	total := len(numbers)
	stride := total / goroutines
	wg := sync.WaitGroup{}

	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if end > total {
				end = total
			}
			lv := 0
			for i := start; i < end; i++ {
				lv += numbers[i]
			}
			atomic.AddInt64(&s, int64(lv))
			wg.Done()
		}(g)
	}

	wg.Wait()
	return int(s)
}

func main() {

	numbers := make([]int, 100)

	fmt.Printf("numbers: %d\n", numbers)

	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())

	fmt.Printf("NumGoRoutine: %d\n", runtime.NumGoroutine())

	sim2routines()

	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 5; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("ch1 %x\n", ch)
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
		fmt.Printf("ch2 %x\n", ch)
		fmt.Printf("routines : %d\n", runtime.NumGoroutine())
	}
	testTransitive()
}

// 测试GOGC 对局部变量 的 transitive 管理。
func testTransitive() {
	n := 5
	ps := make([]*int, n)
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(a int) {
			b := a * a
			ps[a] = &b
			wg.Done()
		}(i)
	}

	wg.Wait()

	for _, p := range ps {
		fmt.Printf("local values: %d\n", *p)
	}
}
