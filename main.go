package main

// func main() {

// 	rand.Seed(time.Now().UnixNano())

// 	max, _ := strconv.Atoi(os.Args[1])

// 	var uniques []int
// loop:
// 	for len(uniques) < max; {
// 		n := rand.Intn(max) + 1
// 		fmt.Print(n, " ")

// 		for _, u := range uniques {

// 			if u == n {
// 				continue loop
// 			}

// 		}

// 		//
// 		uniques = append(uniques, n)

// 	}

// 	fmt.Println("\n\nuniques:", uniques)

// 	sort.Ints(uniques)
// 	fmt.Println("\nsorted", uniques)

// 	nums:= [5]int{5,4,3,2,1}
// 	sort.Ints(nums[:])
// 	fmt.Println("\nnums",nums)
// }

// var books [5]string
// books[0] = "dracula"
// books[1] = "1984"
// books[2] = "island"

// games := []string{"kokemon", "sims"}

// newBooks := [5]string{"ulysses", "fire"}
// books = newBooks

// newGames := []string{"pacman", "doom", "pong"}

// newGames = games

// games = nil
// var ok string
// for i, game := range games {
// 	if game != newGames[i] {
// 		ok = "not "
// 		break
// 	}

// 	if len(games) != len(newGames) {
// 		ok = "not "
// 	}

// }
// fmt.Printf("games and new games are %sequal\n\n", ok)
// fmt.Printf("books               : %#v\n", books)
// fmt.Printf("books               : %#v\n", newBooks)
// fmt.Printf("games               : %#v\n", games)

// fmt.Printf("games               : %#v\n", games)
// fmt.Printf("games               : %#v\n", len(games))
// fmt.Printf("games               : %#v\n", games == nil)
