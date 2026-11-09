package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// “准备时间”会根据千层面层数计算出制作所需的时间。
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime 计算制作千层面所花费的时间。该时间包括准备时间和千层面在烤箱中烘烤的时间。
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return (numberOfLayers * 2) + actualMinutesInOven
}
