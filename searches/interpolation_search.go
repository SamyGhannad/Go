package main


//Interpolation, like binary search works for sorted arrays only.
//It is an improved version of binary search for specific cases
//where the distribution of values in the target sorted array is uniform.
//Instead of starting from the middle element like binary search,
//interpolation search will go to different locations based on the value of the target being searched.
func interpolationSearch(array []int, target int,lowIndex int, highIndex int) int {

	if highIndex < lowIndex {
		return -1
	}
	pos := lowIndex + (target-array[lowIndex])*(highIndex-lowIndex)/(array[highIndex]-array[lowIndex])

	if array[pos] > target {
		return binarySearch(array, target, lowIndex, pos)
	} else if array[pos] < target {
		return binarySearch(array, target, pos+1, highIndex)
	} else {
		return pos
	}
}


func iterativeInterpolationSearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var pos int
	for startIndex < endIndex {
		pos = startIndex + (target-array[startIndex])*(endIndex-startIndex)/(array[endIndex]-array[startIndex])
		if array[pos] > target {
			endIndex = pos-1
		} else if array[pos] < target {
			startIndex = pos+1
		} else {
			return pos
		}
	}
	return -1
}
