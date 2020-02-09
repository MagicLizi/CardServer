package fight

import "math/rand"

func LotteryCards(from []FCard, count int) []FCard {
	var result []FCard

	if len(from) >= count {
		for i := 0; i < count; i++ {
			lotteryIndex := rand.Intn(len(from))
			lotteryCard := from[lotteryIndex]
			from = append(from[:lotteryIndex], from[lotteryIndex+1:]...)
			result = append(result, lotteryCard)
		}
		return result
	} else {
		return from
	}
}

func CardMapToIds(cardMap map[string]FCard) []string {
	var ids []string
	for k, _ := range cardMap {
		ids = append(ids, k)
	}
	return ids
}
