package main

//例子：
//
//   map
//
//   初始化   map[string]int
//
//   stu01  score 随机生成（0-99)
//   ..
//   stu10
//
//   ------
//
//   取出map中的key存入切片；
//   对切片进行排序，
//   最后输出切片的内容；；
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	map1 := make(map[string]int, 50)
	for i := 1; i <= 10; i++ {
		key := fmt.Sprintf("%d", i)
		value := rand.Intn(100)
		map1[key] = value

		keys := make([]string, 0, 200)
		for key := range map1 {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println(key, map1[key])
		}
	}

}
