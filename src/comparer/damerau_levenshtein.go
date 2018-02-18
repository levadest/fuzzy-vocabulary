package comparer

// DamerauLevenshtein -  is a string metric
// for measuring the difference between two
// sequences. Informally, the Levenshtein
// distance between two words is the minimum
// number of single-character edits
// (insertions, deletions or substitutions)
// required to change one word into the other
func DamerauLevenshtein(source string, targetWords []string) (result int) {
	//start := time.Now()
	//defer func(){
	//	t := time.Now()
	//	elapsed := t.Sub(start)
	//	fmt.Printf("%s took %+v\n", source, elapsed)
	//}()

	sourceWordRuneSlice := []rune(source)
	for iterator, targetWord := range targetWords {
		var iterationResult int

		targetWordRuneSlice := []rune(targetWord)
		inf := len(sourceWordRuneSlice) + len(targetWordRuneSlice)

		if len(sourceWordRuneSlice) == 0 {
			iterationResult = len(targetWordRuneSlice)
		}
		if len(targetWordRuneSlice) == 0 {
			iterationResult = len(sourceWordRuneSlice)
		}

		if len(sourceWordRuneSlice) != 0 && len(targetWordRuneSlice) != 0 {
			baseMatrixArray := make([][]int, len(sourceWordRuneSlice))
			for i := range baseMatrixArray {
				baseMatrixArray[i] = make([]int, len(targetWordRuneSlice))
			}

			seenRunes := make(map[rune]int)

			if sourceWordRuneSlice[0] != targetWordRuneSlice[0] {
				baseMatrixArray[0][0] = 1
			}

			seenRunes[sourceWordRuneSlice[0]] = 0
			for i := 1; i < len(sourceWordRuneSlice); i++ {
				deleteSourceDistance := baseMatrixArray[i-1][0] + 1
				insertSourceDistance := (i+1)*1 + 1
				var matchSourceDistance int
				if sourceWordRuneSlice[i] == targetWordRuneSlice[0] {
					matchSourceDistance = i
				} else {
					matchSourceDistance = i + 1
				}
				baseMatrixArray[i][0] = min(min(deleteSourceDistance, insertSourceDistance), matchSourceDistance)
			}

			for j := 1; j < len(targetWordRuneSlice); j++ {
				deleteTargetDistance := (j + 1) * 2
				insertTargetDistance := baseMatrixArray[0][j-1] + 1
				var matchTargetDist int
				if sourceWordRuneSlice[0] == targetWordRuneSlice[j] {
					matchTargetDist = j
				} else {
					matchTargetDist = j + 1
				}

				baseMatrixArray[0][j] = min(min(deleteTargetDistance, insertTargetDistance), matchTargetDist)
			}

			for i := 1; i < len(sourceWordRuneSlice); i++ {
				var maxSrcMatchIndex int
				if sourceWordRuneSlice[i] == targetWordRuneSlice[0] {
					maxSrcMatchIndex = 0
				} else {
					maxSrcMatchIndex = -1
				}

				for j := 1; j < len(targetWordRuneSlice); j++ {
					swapIndex, ok := seenRunes[targetWordRuneSlice[j]]
					jSwap := maxSrcMatchIndex
					deleteDist := baseMatrixArray[i-1][j] + 1
					insertDist := baseMatrixArray[i][j-1] + 1
					matchDist := baseMatrixArray[i-1][j-1]
					if sourceWordRuneSlice[i] != targetWordRuneSlice[j] {
						matchDist += 1
					} else {
						maxSrcMatchIndex = j
					}

					var swapDist int
					if ok && jSwap != -1 {
						iSwap := swapIndex
						var preSwapCost int
						if iSwap == 0 && jSwap == 0 {
							preSwapCost = 0
						} else {
							preSwapCost = baseMatrixArray[maxInteger(0, iSwap-1)][maxInteger(0, jSwap-1)]
						}
						swapDist = i + j + preSwapCost - iSwap - jSwap - 1
					} else {
						swapDist = inf
					}
					baseMatrixArray[i][j] = min(min(min(deleteDist, insertDist), matchDist), swapDist)
				}
				seenRunes[sourceWordRuneSlice[i]] = i
			}
			iterationResult = baseMatrixArray[len(sourceWordRuneSlice)-1][len(targetWordRuneSlice)-1]
		}

		if iterator == 0 {
			result = iterationResult
		}
		if result > iterationResult {
			result = iterationResult
		}

	}
	return result
}

// min of two integers
func min(a, b int) (res int) {
	if a < b {
		res = a
	} else {
		res = b
	}
	return
}

// max of two integers
func maxInteger(a, b int) (res int) {
	if a < b {
		res = b
	} else {
		res = a
	}
	return
}
