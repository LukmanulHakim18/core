package feature

import "hash/fnv"

func IsUserInRolloutPercentage(input string, rolloutPercentage int64) bool {
	if rolloutPercentage == 0 {
		return false
	}

	// Use FNV-1a 32-bit hash
	h := fnv.New32a()
	h.Write([]byte(input))
	hashValue := h.Sum32()

	// Calculate the user percentage based on hash value
	userPercentage := hashValue % 100
	return int(userPercentage) < int(rolloutPercentage)
}
