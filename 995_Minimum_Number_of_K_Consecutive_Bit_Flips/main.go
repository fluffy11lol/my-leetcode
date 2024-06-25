package main

func main() {

}
func minKBitFlips(nums []int, k int) int {
	flipped := make([]bool, len(nums))
	validFlipsFromPastWindow := 0
	flipCount := 0
	for i := range len(nums) {
		if i >= k {
			if flipped[i-k] {
				validFlipsFromPastWindow--
			}
		}

		if validFlipsFromPastWindow%2 == nums[i] {
			if i+k > len(nums) {
				return -1
			}

			validFlipsFromPastWindow++
			flipped[i] = true
			flipCount++
		}
	}

	return flipCount
}

//            // Check if current bit needs to be flipped
//            if (validFlipsFromPastWindow % 2 == nums[i]) {
//                // If flipping the window extends beyond the array length,
//                // return -1
//                if (i + k > nums.size()) {
//                    return -1;
//                }
//                // Increment the count of valid flips and mark current as
//                // flipped
//                validFlipsFromPastWindow++;
//                flipped[i] = true;
//                flipCount++;
//            }
//        }
//
//        return flipCount;
//    }
//};
