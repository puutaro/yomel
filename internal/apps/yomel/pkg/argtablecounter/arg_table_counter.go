package argtablecounter

func IncrementStageNo(isStage bool) int {
	if isStage {
		return 1
	}
	return 0
}
