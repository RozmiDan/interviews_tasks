package main

import (
	"fmt"
)

//-------------------------------- fst_task

// func main() {
// 	println(handle().Error())
// }

// type MyError struct{}

// func (m MyError) Error() string {
// 	return "some error"
// }

// func handle() error {
// 	return MyError{}
// }

//-------------------------------- scnd_task

// func main() {
// 	fmt.Println(uniqRandn(10))
// }

func uniqRandn(n int) []int {
	mp := make(map[int]struct{}, n)
	resultArr := make([]int, 0, n)

	for len(resultArr) < n {
		randVal := rand.Int()
		if _, ok := mp[randVal]; !ok {
			mp[randVal] = struct{}{}
			resultArr = append(resultArr, randVal)
		}
	}

	return resultArr
}

//-------------------------------- thd_task

// func main() {
// 	s1, s2 := []int{1, 2, 3}, []int{3, 4, 5, 7, 8}
// 	fmt.Println(zip(s1, s2))
// }

// func zip(s1, s2 []int) [][]int {
// 	minLen := min(len(s1), len(s2))
// 	resStr := make([][]int, 0, minLen)

// 	for i := 0; i < minLen; i++ {
// 		curVal := []int{s1[i], s2[i]}
// 		resStr = append(resStr, curVal)
// 	}

// 	return resStr
// }

//-------------------------------- fth_task

// func merge(cs ...<-chan int) <-chan int {
// 	resCh := make(chan int, len(cs))
// 	wg := &sync.WaitGroup{}

// 	for _, ch := range cs {
// 		wg.Add(1)

// 		go func(channel <-chan int) {
// 			defer wg.Done()
// 			for val := range channel {
// 				resCh <- val
// 			}
// 		}(ch)
// 	}

// 	go func() {
// 		defer close(resCh)
// 		wg.Wait()
// 	}()

// 	return resCh
// }

// -------------------------------- sxth_task

// type X struct {
// 	V int
// }

// func (x *X) S() {
// 	fmt.Println(x.V)
// }

// func main() {
// 	x := X{123}
// 	defer x.S()
// 	x.V = 456
// }

// -------------------------------- svnth_task

// func main() {
// 	var urls = []string{
// 		"http://ozon.ru",
// 		"http://ya.ru",
// 		"http://yandex.ru",
// 		"http://mail.ru",
// 	}

// 	var client http.Client

// 	for _, url := range urls {
// 		resp, err := client.Get(url)
// 		if err == nil {
// 			if resp.StatusCode == 200 {
// 				fmt.Println("ok")
// 			} else {
// 				fmt.Println("not ok")
// 			}
// 		} else {
// 			fmt.Println("some error")
// 		}
// 	}

// }

func main() {

	s := "test"
	fmt.Println(s[0])
	s[0] = "R"
	fmt.Println(s)


	// str := []rune(s)
	// str[0] = 'R'
	// fmt.Println(string(str))
}
