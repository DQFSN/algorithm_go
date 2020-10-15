package main




func minArray(numbers []int) int {
    low, high := 0, len(numbers)-1

    if len(numbers) == 0 {
        return -1
    }
    mid := low + (high - low) / 2

    for low < high {

        if numbers[mid] < numbers[high] {
            high = mid
        }else if numbers[mid] > numbers[high] {
            low = mid + 1
        }else {
            high--
        }
        mid = low + (high - low) / 2


    }

    return numbers[mid]
}