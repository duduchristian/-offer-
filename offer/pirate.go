package offer

func GetBestPlans(N, coins int) (plans [][]int) {
	if N < 3 {
		return
	}
	plans = append(plans, []int{coins, 0, 0})
	for numPirates := 4; numPirates <= N; numPirates++ {
		expectations := getExpectations(plans)
		for get := coins; get >= 0; get-- {
			left := coins - get
			suitableResults := getSuitableResults(numPirates, get, left, expectations)
			if len(suitableResults) == 0 {
				continue
			}
			plans = suitableResults
			break
		}
	}

	return
}

func getSuitableResults(N, get, left int, expectations []int) [][]int {
	var (
		resultChan = make(chan []int, N*2)
		plan       = make([]int, N-1)
	)
	go func() {
		allocateCoins(plan, left, N-1, resultChan)
		close(resultChan)
	}()
	var suitableResults [][]int
	for result := range resultChan {
		if len(result) == 0 {
			break
		}
		if isSuitablePlan(expectations, result) {
			suitableResult := make([]int, N)
			suitableResult[0] = get
			for i := range result {
				suitableResult[i+1] = result[i]
			}
			suitableResults = append(suitableResults, suitableResult)
		}
	}
	return suitableResults
}

// O(N)
func getExpectations(plans [][]int) []int {
	expectations := make([]int, len(plans[0]))
	for _, plan := range plans {
		for i := 0; i < len(plan); i++ {
			expectations[i] += plan[i]
		}
	}
	for i := 0; i < len(expectations); i++ {
		expectations[i] /= len(plans)
	}
	return expectations
}

// O(N)
func isSuitablePlan(expectations []int, plan []int) bool {
	N := len(expectations)
	agree := 0
	for i := 0; i < N; i++ {
		if expectations[i] < plan[i] {
			agree++
		}
	}
	return agree >= (N+1)/2
}

// O(N^C)
func allocateCoins(plan []int, coins, N int, resultChan chan<- []int) {
	if coins == 0 {
		result := make([]int, len(plan))
		copy(result, plan)
		resultChan <- result
		return
	}
	if N == 1 {
		result := make([]int, len(plan))
		copy(result, plan)
		result[len(result)-1] = coins
		resultChan <- result
		return
	}

	for i := coins; i >= 0; i-- {
		plan[len(plan)-N] = i
		allocateCoins(plan, coins-i, N-1, resultChan)
	}
}
