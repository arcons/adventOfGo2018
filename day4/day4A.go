package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func day1A() {
	input := input()
	fmt.Print("Day four AOCC\n")
	//create guard map with their number and 60 int array
	//+1 every time they fall asleep at a given minute
	//keep track of largest +1 and guard number
	var inputLength = len(input)
	var maxTime = 1
	var maxGuard = ""

	var m map[time.Time]int
	timeArray := []time.Time{}
	m = make(map[time.Time]int)

	var guardTotalSleep map[string]float64
	guardTotalSleep = make(map[string]float64)
	var guardSleepDuartions map[string][60]int
	guardSleepDuartions = make(map[string][60]int)
	var startSleep = time.Now()
	var currentGuard = ""

	//Magic time form https://programming.guide/go/format-parse-string-time-date-example.html
	const timeForm = "2006-01-02T15:04:05"
	for i := 0; i < inputLength; i++ {
		//split the input at char 19
		s := input[i]
		splitString := strings.Split(s, "]")
		// format the string to get a UTC time including seconds
		inputT := strings.Replace(splitString[0], "[", "", -1)
		inputT = strings.Replace(inputT, " ", "T", -1)
		inputT += ":00"
		timestamp, _ := time.Parse(timeForm, inputT)
		//map the timestamp to the index
		m[timestamp] = i
		timeArray = append(timeArray, timestamp)
	}

	//Go 1.8 sort for time.Time this took entirely too long to figure out
	sort.Slice(timeArray, func(i, j int) bool { return timeArray[i].Before(timeArray[j]) })
	for i := 0; i < inputLength; i++ {
		//timestamps are in order now and you have a map of timestamp to index
		index := m[timeArray[i]]
		//sting of the item
		s := input[index]
		splitString := strings.Split(s, "]")
		// dateTime := strings.Split(strings.Replace(splitString[0], "[", "", -1), " ")

		//use the time
		// date index 0 time 1
		// only need minutes since the datetime is already in order
		// fmt.Print(dateTime[1])
		action := strings.Split(strings.TrimSpace(splitString[1]), " ")
		if action[0] == "Guard" {
			// fmt.Print(" @ Guard starts duty" + action[1])
			currentGuard = action[1]
		} else if action[0] == "falls" {
			// fmt.Print(" @ Sleepy time\n")
			startSleep = timeArray[i]
		} else {
			// fmt.Print(" @ Wakey wakey\n")
			duration := timeArray[i].Sub(startSleep)
			guardTotalSleep[currentGuard] += duration.Minutes()
			// check if the guard time is less than max time
			if int(guardTotalSleep[currentGuard]) > maxTime {
				maxTime = int(guardTotalSleep[currentGuard])
				maxGuard = currentGuard
			}
			//populate array to reload into the map
			oldTimes := guardSleepDuartions[currentGuard]
			//loop through duration they were asleep and increment the minutes that they slept
			startMin := startSleep.Minute()
			durMin := int(duration.Minutes())
			if startMin+durMin > 59 {
				startMin = startMin - durMin
			}
			for min := 0; min < durMin; min++ {
				//need to start at startSleep minute
				oldTimes[min+startMin]++
			}
			guardSleepDuartions[currentGuard] = oldTimes
		}
	}
	var maxMin = 0
	var maxManInd = 0
	for i := 0; i < 60; i++ {
		if guardSleepDuartions[maxGuard][i] > maxMin {
			maxMin = guardSleepDuartions[maxGuard][i]
			maxManInd = i
		}
	}

	fmt.Print("Guard ")
	fmt.Print(maxGuard)
	fmt.Print("\n")
	fmt.Print("Minute ")
	fmt.Print(maxManInd)
	fmt.Print("\n")
}

func testInput() []string {
	input := []string{
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep",
		"[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep",
		"[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep",
		"[1518-11-05 00:55] wakes up",
	}
	return input
}
