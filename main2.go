package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Multidementional example
	spendings := fetch()

	for i, daily := range spendings {

		var total int
		for _, spending := range daily {
			total += spending
		}

		fmt.Printf("Day %d: %d\n", i+1, total)
	}

}

func fetch() [][]int {

	content := `200 100
	 25 10 45 60
	  5 15 35
	  95 10
	  50 25`

	lines := strings.Split(content, "\n")

	spendings := make([][]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {
			spending, _ := strconv.Atoi(field)

			spendings[i][j] = spending
		}

	}

	return spendings
}

// //Copy function
// data := []float64{10, 25, 30, 50}

// copy(data, []float64{99, 100})
// // n := copy(data, []float64{10, 5, 15, 0, 20})
// // fmt.Printf("%d probabilities copied\n", n)

// // data = append(data[:0], []float64{10, 5, 15, 0, 20}...)

// // saved := make([]float64, len(data))
// // copy(saved, data)

// saved := append([]float64(nil), data...)

// data[0] = 0

// fmt.Println("probabilities (saved)", saved)
// fmt.Println("probabilities (data)", data)
// fmt.Printf("Is it gonna rain? %.f%% chance.\n",
// 	(data[0]+data[1]+data[2]+data[3])/float64(len(data)))

//Make function
// tasks := []string{"jump", "run", "read"}

// upTasks := make([]string, 0, len(tasks))
// fmt.Println("upTasks", upTasks)

// for _, task := range tasks {
// 	upTasks = append(upTasks, strings.ToUpper(task))
// 	fmt.Println("upTasks", upTasks)
// }

//
// nums := []int{1, 2, 3, 4}

// odds := append(nums[:2:2], 5, 7)
// evens := append(nums[2:4], 6, 8)
// fmt.Println("nums", nums)
// fmt.Println("odds", odds)
// fmt.Println("evens", evens)

// ages, oldCap := []int{1}, 1.

// for len(ages) < 5e5 {
// 	ages = append(ages, 1)

// 	c := float64(cap(ages))
// 	if c != oldCap {
// 		fmt.Printf("len:%-10d cap:%-10g growth:%.2f\n",
// 			len(ages), c, c/oldCap)

// 	}

// 	oldCap = c

// }

// var nums []int
// fmt.Println("no backing array", nums)

// nums = append(nums, 1, 3)
// fmt.Println("allocates", nums)

// nums = append(nums, 1, 3)
// fmt.Println("allocates", nums)

// nums = append(nums, 2)
// fmt.Println("free capacity", nums)

// nums = append(nums, 4)
// fmt.Println("no allocation", nums)

// nums = append(nums, nums[2:]...)
// fmt.Println("nums <- nums[2:]", nums)

// nums = append(nums[2:], 7, 9)
// fmt.Println("nums[2:] <- 7,9", nums)

// nums = nums[:6]
// fmt.Println("nums: extend", nums)

// var games []string
// fmt.Println("games", games)

// games = []string{}
// fmt.Println("games", games)

// games = []string{"pacman", "mario", "tetris", "doom"}
// fmt.Println("games", games)

// part := games
// fmt.Println("part", part)

// part = games[:0]
// fmt.Println("part[:0]", part)
// fmt.Println("part[:cap]", part[:cap(part)])

// for cap(part) != 0 {

// 	part = part[1:cap(part)]
// 	fmt.Println("part", part)

// }

// games = games[len(games):]
// fmt.Println("games", games)
// fmt.Println("part", part)

// func main() {

// 	data := collection{"slices", "are", "awesome", "period"}
// 	change(data)
// 	fmt.Println("mains data", data)
// 	fmt.Printf("mains data slice adress: %p\n", &data)
// }

// func change(data collection) {
// 	data[2] = "brilliant"
// 	fmt.Println("changes data", data)
// 	fmt.Printf("change data slice adress: %p\n", &data)

// }

// grades := []float64{40, 10, 20, 50, 60, 70}

// // var newGrades []float64
// newGrades := append([]float64(nil), grades...)

// front := newGrades[:3]
// front2 := front[:3]
// front3 := front

// sort.Float64s(front)

// fmt.Println("grades", grades[:])
// fmt.Println("newGrades", newGrades[:])
// fmt.Println("front", front)
// fmt.Println("front2", front2)
// fmt.Println("front3", front3)

// var todo []string

// todo = append(todo, "sing")
// todo = append(todo, "run", "code", "play")

// fmt.Println("todo", todo)

// items := []string{

// 	"pacman", "mario", "tetris", "doom", "galaga", "frogger", "asteriods",
// 	"simcity", "metriod", "defender", "rayman", "tempest", "ultima",
// }

// const pageSize = 4

// l := len(items)
// for from := 0; from < l; from += pageSize {

// 	to := from + pageSize
// 	if to > l {
// 		to = l
// 	}

// 	currentPage := items[from:to]

// 	head := fmt.Sprintf("Page #%d", (from/pageSize)+1)
// 	fmt.Println(head, currentPage)

// }

// fmt.Println("items", items)

// top3 := items[:3]
// fmt.Println("top 3 items", top3)

// l := len(items)

// last4 := items[l-4:]
// fmt.Println("last 4 items", last4)

// mid := last4[1:4]
// fmt.Println("mid items", mid)

// fmt.Printf("slicing : %T %[1]q\n", items[2:3])
// fmt.Printf("indexing : %T %[1]q\n", items[2])
