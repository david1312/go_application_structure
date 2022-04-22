package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// fmt.Println(countingValleys(8, "UDDDUDUU"))                                         // https://www.hackerrank.com/challenges/counting-valleys/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
	// fmt.Println(sockMerchant(9, []int32{10, 10, 10, 20, 20, 30, 30, 60, 80}))           // https://www.hackerrank.com/challenges/sock-merchant/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
	// fmt.Println(arrayManipulation(5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}})) // https://www.hackerrank.com/challenges/crush/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
	// fmt.Println(jumpingOnClouds([]int32{0, 0, 0, 0, 0, 1, 0, 0, 1, 0}))                 // https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup

	// fmt.Println(perfectSubstring2("1102021222", 2)) //avg time 1753
	// fmt.Println(perfectSubstring("1102021222", 2))  //avg time 1753
	// testTimer()
	// test := []int32{1,2,3,4,5}
	// fmt.Println(arrayManipulationHack([]string{"1 4", "2 3", "4 1"}))
	// testLogic()
	// hackerRank2()
	// coba()
	// a := "asb"
	// fmt.Println(string(a[2]))
	fmt.Println(compressedString("wwww"))
	item := [][]string{{"p1", "5", "5"}, {"p2", "20", "2"}, {"p2", "5", "100"}}
	sortParam := int32(0) // 0 name 1 relevance 2 price by kolom aja
	sortOrder := int32(0) // ascending // 1 descending
	itemperPage := int32(2)
	pageNumber := int32(1) // page number start from 0
	fmt.Println(fetchItemsToDisplay(item, sortParam, sortOrder, itemperPage, pageNumber))

	fmt.Println("done")
	var occurenceArr = make(map[int]int)
	occurenceArr[2]++
	fmt.Println(solution([]int{1, 2, 3, 9, 3, 5, 10, 100}))

	fmt.Println("done xxx")
	fmt.Println(solution2Ajaib(10, 10))

	fmt.Println("done 2")
	fmt.Println(solution3Ajaib([]int{3, 7, 2, 5, 4}))

}

func solution3Ajaib(A []int) string {
	var (
		sumR, sumG, sumB = 0, 0, 0
		result           = ""
	)
	totalColor := 3

	totalValue := 0
	for _, v := range A {
		totalValue += v
	}
	avg := totalValue / totalColor
	if totalValue%totalColor == 0 && len(A) >= totalColor {
		for _, v := range A {
			if sumR < avg && sumR+v <= avg {
				sumR += v
				result += "R"
			} else if sumG < avg && sumG+v <= avg {
				sumG += v
				result += "G"
			} else if sumB < avg && sumB+v <= avg {
				sumB += v
				result += "B"
			}
		}
		fmt.Println(avg)
		return result
	}

	return "impossible"
}

func solution2Ajaib(A int, B int) int {
	calcInt := A * B
	binary := fmt.Sprintf("%b", calcInt)
	fmt.Println(binary)
	return CountSetOfBits(calcInt)
}

func CountSetOfBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

//SELECT web_page_ipv4,count(distinct user_ipv4) AS users_cnt FROM visits  GROUP BY web_page_ipv4
func solution(A []int) int {
	var (
		max, min     = A[0], A[0]
		occurenceArr = make(map[int]int)
		result       = 1
	)
	for _, val := range A {
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}
		occurenceArr[val]++
	}

	if max <= 0 {
		return result
	}
	gotZeroOccurence := false
	for i := min; i <= max; i++ {
		if occurenceArr[i] == 0 {
			result = i
			gotZeroOccurence = true
			break
		}
	}
	if !gotZeroOccurence {
		return max + 1
	}
	return result
}

func fetchItemsToDisplay(items [][]string, sortParameter int32, sortOrder int32, itemsPerPage int32, pageNumber int32) []string {
	result := []string{}
	listProduct := []string{}
	for _, val := range items {
		listProduct = append(listProduct, val[0])
	}
	switch sortParameter {
	case 0:
		sort.Strings(listProduct)

	}
	startAddResult := ((pageNumber+1)*itemsPerPage - itemsPerPage)
	endAddResult := ((pageNumber + 1) * itemsPerPage) - 1
	fmt.Println(startAddResult, endAddResult)
	return result

}

func compressedString(message string) string { //hackerrank tokped soal 3
	result := ""
	sameString := 1
	currentString := ""
	for k, v := range message {
		valString := string(v)
		if currentString != valString {
			result += valString
		} else if (k+1) != len(message) && valString != string(message[k+1]) && currentString == valString {
			sameString++
			result += fmt.Sprintf("%v", sameString)
			sameString = 1
		} else if (k+1) != len(message) && valString == string(message[k+1]) && currentString == valString {
			sameString++
		} else if (k+1) == len(message) && currentString == valString {
			sameString++
			result += fmt.Sprintf("%v", sameString)
		}
		currentString = valString
	}
	return result
}

type resWikipedia struct {
	Parse struct {
		Title  string `json:"title"`
		PageID int    `json:"pageid"`
		Text   struct {
			Content string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
}

func getTopicCount(topic string) int { //tokopedia hackerrank soal 2
	urlRequest := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&section=0&prop=text&format=json&page=%v", topic)

	resp, err := http.Get(urlRequest)
	var data resWikipedia
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&data)
	result := strings.Count(data.Parse.Text.Content, "Pizza")

	return result
}

func hackerRank2() error {
	// client := &http.Client{}
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&section=0&prop=text&format=json&page=%v", "france")
	fmt.Println(url)
	resp, err := http.Get(url)
	var data resWikipedia
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	_ = decoder.Decode(&data)
	// fmt.Println(data.Parse.Text.Content)
	result := strings.Count(data.Parse.Text.Content, "Pizza")
	fmt.Println(result)
	return nil
}

func testLogic() {
	fmt.Println("testlogic")
	a := []int{1, 2, 3, 4}
	a[3]++
	if 4 >= len(a) {
		a = append(a, 10)
	}

	b := make([]int, 10)
	b[8]++
	fmt.Println(b)
}

func arrayManipulationHack(queries []string) int64 {
	start := time.Now()
	var (
		horizontal, vertical int
		queryNumber          []int
	)
	for _, val := range queries {
		splitter := strings.Split(val, " ")
		verVal, _ := strconv.Atoi(splitter[0])
		horVal, _ := strconv.Atoi(splitter[1])
		if vertical < verVal {
			vertical = verVal
		}
		if horizontal < horVal {
			horizontal = horVal
		}
		queryNumber = append(queryNumber, verVal, horVal)
	}
	arraySize := horizontal * vertical

	// fmt.Println(horizontal, vertical, queryNumber)
	arrResult := make([]int64, arraySize)
	var maxValue int64
	startQuery := 0
	lengthLoop := len(queryNumber) / 2
	for i := 1; i <= lengthLoop; i++ {
		// fmt.Println(queryNumber[startQuery], queryNumber[startQuery+1])
		for k := 1; k <= queryNumber[startQuery]; k++ {
			startSum := (k*horizontal - horizontal)
			endSum := (k-1)*horizontal + queryNumber[startQuery+1] - 1
			for index := startSum; index <= endSum; index++ {
				arrResult[index]++
			}
		}
		startQuery += 2
	}
	for _, v := range arrResult {
		if maxValue < v {
			maxValue = v
		}
	}
	fmt.Println(arrResult)
	// for i := 0; i < len(queries); i++ {
	// 	start := queryNumber[startQuery] - 1
	// 	valToAdd := 1
	// 	arrResult[start] += valToAdd
	// 	if queries[i][1] < n {
	// 		arrResult[queries[i][1]] -= valToAdd
	// 	}
	// }
	// for j := int32(1); j < n; j++ {
	// 	arrResult[j] += arrResult[j-1]
	// 	if arrResult[j] > maxValue {
	// 		maxValue = arrResult[j]
	// 	}
	// }
	fmt.Println((time.Since(start)).Nanoseconds())
	return int64(maxValue)
}

func testTimer() {
	start := time.Now()
	time.Sleep(5 * time.Second)
	fmt.Println((time.Since(start)).Nanoseconds())
}

func perfectSubstring2(s string, k int32) int32 {
	start := time.Now()
	result := int32(0)
	occurence := int(k)
	for key, _ := range s {

		endIndex := key + occurence

		totalLoopInside := (len(s) - key) / occurence
		for i := 1; i <= totalLoopInside; i++ {
			checkString := s[key:endIndex]
			endIndex += occurence
			isPerfect := true
			tempMap := make(map[string]int32)
			for _, valStr := range checkString {
				tempMap[string(valStr)]++
			}

			for _, v := range tempMap {
				if v != 2 {
					isPerfect = false
				}
			}
			if isPerfect {
				result++
			}
		}
	}
	fmt.Println((time.Since(start)).Nanoseconds())
	return result //result 6
}

func perfectSubstring(s string, k int32) int32 {
	result := int32(0)
	occurence := int(k)
	start := time.Now()
	for key, _ := range s {

		endIndex := key + occurence

		totalLoopInside := (len(s) - key) / occurence
		for i := 1; i <= totalLoopInside; i++ {
			checkString := s[key:endIndex]
			endIndex += occurence
			isPerfect := true
			for _, valStr := range checkString {
				if strings.Count(checkString, string(valStr)) != 2 {
					isPerfect = false
					break
				}
			}
			if isPerfect {
				result++
			}
		}
	}
	fmt.Println((time.Since(start)).Nanoseconds())
	return result //result 6
}

func countingValleys(steps int32, path string) int32 {
	var (
		currentStep, enterValley = 0, 0
	)
	for _, v := range path {
		switch string(v) {
		case "U":
			currentStep++
			if currentStep == 0 {
				enterValley++
			}

		default:
			currentStep--
		}
	}
	return int32(enterValley)
}

func sockMerchant(n int32, ar []int32) int32 {
	countingSocksVariant := make(map[int32]int32)
	numberOfPairs := int32(0)
	for _, val := range ar {
		countingSocksVariant[val]++
	}
	//countingpairs
	for _, val := range countingSocksVariant {
		if val > 1 {
			numberOfPairs += (val / 2)
		}
	}
	return numberOfPairs

}

func arrayManipulation(n int32, queries [][]int32) int64 {
	start := time.Now()
	arrResult := make([]int64, n)
	maxValue := int64(0)
	for i := 0; i < len(queries); i++ {
		start := queries[i][0] - 1
		valToAdd := int64(queries[i][2])
		arrResult[start] += valToAdd
		if queries[i][1] < n {
			arrResult[queries[i][1]] -= valToAdd
		}
	}
	for j := int32(1); j < n; j++ {
		arrResult[j] += arrResult[j-1]
		if arrResult[j] > maxValue {
			maxValue = arrResult[j]
		}
	}
	fmt.Println((time.Since(start)).Nanoseconds())
	return int64(maxValue)
}

func jumpingOnClouds(c []int32) int32 {
	counter := int32(0)
	currentStep := 0
	for currentStep != len(c)-1 {
		counter++
		if currentStep+2 <= len(c)-1 && c[currentStep+2] == 0 {
			currentStep += 2
		} else {
			currentStep++
		}
	}
	return counter
}
