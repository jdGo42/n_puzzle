package algo

func CalculSigmaDistance(solvedArray []int, currentArray []int, size int) (sigmaDistance int, err error) {
	sigmaDistance = 0
	err = nil
	for i := 0; i < len(currentArray) - 1; i++ {
		value := solvedArray[i]
		currentPosition, err := returnIndexOfAValue(value, currentArray)
		if err != nil {
			sigmaDistance += getDistanceBetweenTwoPoints(i, currentPosition, size)
		} else {
			return sigmaDistance, err
		}
	}
	return sigmaDistance, err
}

func returnIndexOfAValue(value int, currentArray []int) (result int, err error) {
	for i := 0; i < len(currentArray) - 1; i++ {
		if currentArray[i] == value {
			result = i
			err = nil
			return result, err
		}
	}
	return -1, errors.New("Value has not been reached into returnIndexOfAValue")
}

func getDistanceBetweenTwoPoints(index, currentPosition, size int) int {
	distanceInLine := absolute((index % size) - (currentPosition % size))
	distanceInColumn := absolute((index / size) - (currentPosition / size))
	totalDistance := distanceInColumn + distanceInLine
	return totalDistance 
}

func absolute(negativeNbr int) int {
	positiveNumber := negativeNbr
	if negativeNbr < 0 {
		positiveNumber = positiveNumber * -1
	}
	return positiveNumber
}