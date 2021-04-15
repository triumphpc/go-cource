package main

// storage for private params Context class
type privateValues struct {
	// current active concrete strategy
	actionStrategy Strategy
	// current user filters
	userFilters map[string]int
}

// Initialization default strategy
func initStrategy(s Strategy) *privateValues {
	// init user filters
	f := make(map[string]int)

	return &privateValues{
		s,
		f,
	}
}

// Set new active strategy
func (v *privateValues) setStrategy(s Strategy) {
	v.actionStrategy = s
}

// get data by filters
func getData(v *privateValues) {
	v.actionStrategy.doSearch(v.userFilters)
}
