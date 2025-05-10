package feature

import (
	"strconv"
	"testing"
)

// TestIsUserInRolloutPercentage function to test the eligibility of IDs based on rollout percentages
func TestIsUserInRolloutPercentage(t *testing.T) {
	// The 10 unique BB IDs generated
	bbIDs := []string{
		"BB71452011", "BB81999539", "BB35533637", "BB89162168", "BB15885076",
		"BB44135507", "BB16950341", "BB48431031", "BB03748401", "BB85877089",
	}

	// Define the expected number of eligible IDs for each rollout percentage
	expectedResults := map[int64]int{
		20:  3,
		50:  5,
		70:  8,
		100: 10,
	}

	for percentage, expectedCount := range expectedResults {
		t.Run(strconv.FormatInt(percentage, 10), func(t *testing.T) {
			actualCount := 0
			for _, bbID := range bbIDs {
				if IsUserInRolloutPercentage(bbID, percentage) {
					actualCount++
				}
			}

			if actualCount != expectedCount {
				t.Errorf("For %d%% rollout, expected %d eligible IDs, but got %d", percentage, expectedCount, actualCount)
			}
		})
	}
}
