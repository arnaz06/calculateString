package calculateString

// Calculate is function to calculate the valid formatted string.
// the format must be like this.
// e.g
// "4 + 5 - 1", "542 + 4 - 12", etc...
// this function only support 2 math operator: + & -.
// NOTE: in this function i used ASCCI manipulation.
func Calculate(input string) (int, string) {
	nums, ops := split(input)
	if len(nums) == 0 {
		return 0, ops[0]
	}

	temp := nums[0]
	indexOps := 0
	for i := 1; i < len(nums); i++ {
		switch ops[indexOps] {
		case "+":
			temp += nums[i]
		case "-":
			temp -= nums[i]
		}

		if indexOps == len(ops)-1 {
			break
		}
		indexOps++
	}
	return temp, ""
}

func split(input string) ([]int, []string) {
	// add ! char to simplify the logic
	r := []rune(input + "!")
	mergeInt := []int{}
	nums := []int{}
	ops := []string{}
	for i := 0; i < len(r); i++ {

		//  skip if ASCCI string
		if r[i] == 32 {
			continue
		}

		// transform from rune to int
		if r[i] <= 57 && r[i] >= 48 {
			mergeInt = append(mergeInt, int(r[i]-48))
		} else {
			// merging the int
			if len(mergeInt) == 1 {
				nums = append(nums, mergeInt[0])
			} else {
				zeroes := len(mergeInt) - 1
				temp := 0
				for j := 0; j < len(mergeInt); j++ {
					if j+1 == len(mergeInt) {
						temp += mergeInt[j]
						break
					}
					temp += mergeInt[j] * pow(10, zeroes)
					zeroes--
				}
				nums = append(nums, temp)
			}

			// resolve the math operator
			mergeInt = []int{}
			switch r[i] {
			case 43:
				ops = append(ops, string(r[i]))
			case 45:
				ops = append(ops, string(r[i]))
			case 33:
			default:
				return []int{}, []string{"validation error"}
			}
		}

		if r[i] == 33 {
			break
		}

	}
	return nums, ops
}

func pow(x, y int) int {
	temp := x
	for i := 1; i < y; i++ {
		x *= temp
	}
	return x
}
