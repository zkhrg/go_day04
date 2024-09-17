package candy

var (
	candyPriceList = map[string]int{
		"CE": 10,
		"AA": 15,
		"NT": 17,
		"DE": 21,
		"YR": 23,
	}
)

func IsInCandyList(candyType string) bool {
	_, res := candyPriceList[candyType]
	return res
}

func CalculateOrderPrice(candyType string, candyCount int) int {
	return candyPriceList[candyType] * candyCount
}
