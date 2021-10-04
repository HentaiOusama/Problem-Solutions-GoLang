package Maximizing_Mission_Points_HackerRank

import (
	"Problem_Solutions/CustomTools"
	"fmt"
	"strconv"
	"strings"
)

const MAX = int32(200000)
const MinInt32 = (int32(^uint32(0)/2) * -1) - 1

var fullCityList [MAX]*City
var toolSet *ToolSet

type CoOrdinate struct {
	latitude  int32
	longitude int32
}

type City struct {
	CoOrdinate
	height      int32
	points      int32
	totalPoints int32
}

type ToolSet struct {
	CoOrdinate
	blockMapping             map[int32]*[]*City
	getSignificantCoordinate func(City) int32
	getBlockNumber           func(*City) int32
	isCityInRange            func(city1 *City, city2 *City) (bool, int32)
}

func chooseMax(val1 int32, val2 int32) int32 {
	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}

func getAbsVal(val int32) int32 {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

func generateToolSet(dLat int32, dLong int32) *ToolSet {
	toolSet := new(ToolSet)
	toolSet.latitude = dLat
	toolSet.longitude = dLong
	toolSet.blockMapping = make(map[int32]*[]*City)

	if dLat < dLong {
		toolSet.getSignificantCoordinate = func(city City) int32 {
			return city.latitude
		}

		toolSet.getBlockNumber = func(city *City) int32 {
			return city.latitude / dLat
		}
	} else {
		toolSet.getSignificantCoordinate = func(city City) int32 {
			return city.longitude
		}

		toolSet.getBlockNumber = func(city *City) int32 {
			return city.latitude / dLong
		}
	}

	toolSet.isCityInRange = func(city1 *City, city2 *City) (bool, int32) {
		if (getAbsVal(city1.longitude-city2.longitude) <= dLong) && (getAbsVal(city1.latitude-city2.latitude) <= dLat) {
			return true, chooseMax(chooseMax(city1.totalPoints, city1.totalPoints), chooseMax(city2.points, city2.totalPoints))
		} else {
			return false, MinInt32
		}
	}

	return toolSet
}

func getBlock(blockNumber int32) *[]*City {
	if blockNumber < 0 {
		return nil
	} else {
		return toolSet.blockMapping[blockNumber]
	}
}

func findMaxPointsInBlock(city *City, block *[]*City) int32 {
	currentBlock := *block

	for i := len(currentBlock) - 1; i >= 0; i-- {
		condition, points := toolSet.isCityInRange(city, currentBlock[i])

		if condition {
			return points
		}
	}

	return MinInt32
}

func addCityToBlock(city *City, blockIndex int32) {

	if blockIndex < 0 {
		fmt.Println("Algorithm Error...")
		return
	}
	currentBlock := getBlock(blockIndex)

	if currentBlock == nil {
		toolSet.blockMapping[blockIndex] = &([]*City{
			city,
		})
	} else {
		maxIndex := len(*(currentBlock)) - 1
		insertIndex := maxIndex

		for ; (*currentBlock)[insertIndex].totalPoints > city.totalPoints; insertIndex-- {
		}

		var tempVar []*City

		if insertIndex == maxIndex {
			tempVar = append(*(currentBlock), city)
		} else if insertIndex == -1 {
			tempVar = append([]*City{city}, *currentBlock...)
		} else if insertIndex > -1 && insertIndex < maxIndex {
			tempVar = make([]*City, maxIndex+2)
			i := 0
			i += copy(tempVar[i:], (*currentBlock)[:(insertIndex+1)])
			i += copy(tempVar[i:], []*City{city})
			copy(tempVar[i:], (*currentBlock)[(insertIndex+1):])
		} else {
			fmt.Println("Algorithm Error...")
		}

		toolSet.blockMapping[blockIndex] = &tempVar
	}
}

func performMission() {
	for i := int32(0); i < MAX; i++ {
		if fullCityList[i] != nil {
			processCity(i)
		}
	}
}

func processCity(cityIndex int32) {
	currentCity := fullCityList[cityIndex]
	cityBlockIndex := toolSet.getBlockNumber(currentCity)
	maxSubPoints := MinInt32
	prevBlock, currBlock, nextBlock := getBlock(cityBlockIndex-1), getBlock(cityBlockIndex), getBlock(cityBlockIndex+1)
	if prevBlock != nil {
		maxSubPoints = chooseMax(maxSubPoints, findMaxPointsInBlock(currentCity, prevBlock))
	}
	if currBlock != nil {
		maxSubPoints = chooseMax(maxSubPoints, findMaxPointsInBlock(currentCity, currBlock))
	}
	if nextBlock != nil {
		maxSubPoints = chooseMax(maxSubPoints, findMaxPointsInBlock(currentCity, nextBlock))
	}

	tempPoints := maxSubPoints + currentCity.points
	if tempPoints > currentCity.points {
		currentCity.totalPoints = tempPoints
	} else {
		currentCity.totalPoints = currentCity.points
	}

	addCityToBlock(currentCity, cityBlockIndex)
}

func getMaxPoints() int32 {
	retVal := MinInt32

	for _, city := range fullCityList {
		if city != nil && city.totalPoints > retVal {
			retVal = city.totalPoints
		}
	}

	return retVal
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func RunProgram() {
	reader := CustomTools.GetInputReader(16 * 1024 * 1024)

	firstMultipleInput := strings.Split(CustomTools.ReadLine(reader), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dLatTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	dLat := int32(dLatTemp)

	dLongTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	dLong := int32(dLongTemp)

	for nItr := 0; nItr < int(n); nItr++ {
		secondMultipleInput := strings.Split(CustomTools.ReadLine(reader), " ")

		latitudeTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
		checkError(err)
		latitude := int32(latitudeTemp)

		longitudeTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
		checkError(err)
		longitude := int32(longitudeTemp)

		heightTemp, err := strconv.ParseInt(secondMultipleInput[2], 10, 64)
		checkError(err)
		height := int32(heightTemp)

		pointsTemp, err := strconv.ParseInt(secondMultipleInput[3], 10, 64)
		checkError(err)
		points := int32(pointsTemp)

		// Write your code here
		fullCityList[height-1] = new(City)
		fullCityList[height-1].latitude = latitude
		fullCityList[height-1].longitude = longitude
		fullCityList[height-1].height = height
		fullCityList[height-1].points = points
		fullCityList[height-1].totalPoints = MinInt32
	}

	toolSet = generateToolSet(dLat, dLong)
	performMission()
	fmt.Println(getMaxPoints())
}
