package algo

func CalculSigmaDistance(solvedArray []int, currentArray []int, size int) (int error) {
	sigmaDistance := 0
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

func returnIndexOfAValue(value int, currentArray []int) (int error) {
	for i := 0; i < len(currentArray) - 1; i++ {
		if currentArray[i] == value {
			return i, nil
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