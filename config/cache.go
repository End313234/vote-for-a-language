package config

import "vote-for-a-language/utils"

type CacheUsers map[string]map[string]int64

type cache struct {
	Users CacheUsers
}

func (cu *CacheUsers) Has(userId string) bool {
	for k := range *cu {
		if k == userId {
			return true
		}
	}

	return false
}

func (cu *CacheUsers) Set(userId string, rateLimit map[string]int64) {
	(*cu)[userId] = utils.MergeMaps(rateLimit, map[string]int64{
		"add_language": 0,
		"invite":       0,
		"vote":         0,
	})
}

var Cache = cache{
	Users: make(CacheUsers),
}
