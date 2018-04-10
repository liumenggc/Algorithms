package main

const arraySize = 10

func BinarySearch(array [arraySize]int, number int) int {
	minIndex := 0
	maxIndex := len(array) -1
	for minIndex <=maxIndex  {
		midIndex := int((minIndex + maxIndex)/2)
		midItem := array[midIndex]

		if midItem == number {
			return midIndex
		}

		if midItem < number {
			minIndex = midItem + 1
		}else if midItem > number{
			maxIndex = midItem - 1
		}
	}
	return -1
}