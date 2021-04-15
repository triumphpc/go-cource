package main

func main() {
	// run first strategy
	strategy := &FilterFoFirstAlgorithm{}
	pv := initStrategy(strategy)
	pv.userFilters["role"] = 1
	getData(pv)

	// run second strategy
	strategySecond := &filterFoSecondAlgorithm{}
	pv.userFilters["role"] = 2
	pv.setStrategy(strategySecond)
	getData(pv)

}
