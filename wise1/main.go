package main

import (
	"fmt"
	"sort"
)

type LogActor interface {
	Receive(interval []int32) bool
	AddInterval(time int32)
	Sort()
}
type logActor struct {
	times []int32
}

func (l *logActor) AddInterval(time int32) {

	l.times = append(l.times, time)
}

func (l *logActor) Sort() {
	sort.Slice(l.times, func(i, j int) bool {
		return l.times[i] < l.times[j]
	})
}

func (l *logActor) Receive(interval []int32) bool {

	return binarySearch(l.times, interval)
}

func binarySearch(logTimes []int32, interval []int32) bool {
	mid := len(logTimes) / 2
	switch {
	case len(logTimes) == 0:
		return false
	case logTimes[mid] > interval[1]:
		return binarySearch(logTimes[:mid], interval)
	case logTimes[mid] < interval[0]:
		return binarySearch(logTimes[mid+1:], interval)
	default:
		return true
	}
}

func ProvideLogActor() LogActor {

	return &logActor{times: []int32{}}
}

/*
 * Complete the 'getStaleServerCount' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY log_data
 *  3. INTEGER_ARRAY query
 *  4. INTEGER X
 */
func getStaleServerCount(n int32, log_data [][]int32, query []int32, X int32) []int32 {

	actors := make([]LogActor, n)
	for i := 0; i < int(n); i++ {
		actors[i] = ProvideLogActor()
	}

	for _, datum := range log_data {
		actors[datum[0]-1].AddInterval(datum[1])
	}
	for i := 0; i < int(n); i++ {
		actors[i].Sort()
	}
	ans := make([]int32, len(query))

	for i, q := range query {
		count := int32(0)
		for _, actor := range actors {
			//each actor can be a goroutine
			if actor.Receive([]int32{q - X, q}) {
				count++
			}
		}
		ans[i] = int32(len(actors)) - count
	}

	return ans
}

func main() {

	runExample(3, [][]int32{{1, 3}, {2, 6}, {1, 5}}, []int32{10, 11}, 5) //3
}

func runExample(n int32, logData [][]int32, query []int32, X int32) {
	result := getStaleServerCount(n, logData, query, X)
	fmt.Printf("The result of %v is %v\n", n, result)
}
