package feature

import (
	"strconv"
)

type ConfigValue int

const (
	TestServiceType      ConfigValue = 1 << iota // 1
	ECVRules                                     // 2
	TripVoucher                                  // 4
	CC                                           // 8
	Gopay                                        // 16
	SilverbirdChartered                          // 32
	ShopeePay                                    // 64
	NewPromoEngine                               // 128
	BBCorporate                                  // 256
	ECV                                          // 512
	FixedFareDiscount                            // 1024
	EstimateFareDiscount                         // 2048
	PreAuthCC                                    // 4096
)

var (
	AllDisabled = EnabledFeature{}
	AllEnabled  = NewEnabledFeatures(TestServiceType | ECVRules | TripVoucher | CC | Gopay | SilverbirdChartered | ShopeePay | NewPromoEngine | BBCorporate | ECV | FixedFareDiscount | EstimateFareDiscount | PreAuthCC)
	ListConfig  = map[ConfigValue]string{
		TestServiceType:      "TestServiceType",
		ECVRules:             "ECVRules",
		TripVoucher:          "TripVoucher",
		CC:                   "CC",
		Gopay:                "Gopay",
		SilverbirdChartered:  "SilverbirdChartered",
		ShopeePay:            "ShopeePay",
		NewPromoEngine:       "NewPromoEngine",
		BBCorporate:          "BBCorporate",
		ECV:                  "ECV",
		FixedFareDiscount:    "FixedFareDiscount",
		EstimateFareDiscount: "EstimateFareDiscount",
		PreAuthCC:            "PreAuthCC",
	}
)

type EnabledFeature struct {
	value ConfigValue
}

func (ef *EnabledFeature) Enable(feature ConfigValue) {
	ef.value = ef.value | feature
}

func (ef EnabledFeature) IsEnabled(feature ConfigValue) bool {
	return ef.value&feature != 0
}

func (ef EnabledFeature) ToInt() int {
	return int(ef.value)
}

func (ef EnabledFeature) MarshalJSON() ([]byte, error) {
	timeStr := strconv.Itoa(ef.ToInt())
	return []byte(timeStr), nil
}

func (ef *EnabledFeature) UnmarshalJSON(b []byte) error {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	enabledFeature := NewEnabledFeatures(ConfigValue(i))
	*ef = enabledFeature
	return nil
}

func (ef EnabledFeature) GetList() map[string]bool {
	var list = map[string]bool{}
	for k, c := range ListConfig {
		list[c] = ef.IsEnabled(k)
	}
	return list
}

func NewEnabledFeatures(features ConfigValue) EnabledFeature {
	feature := EnabledFeature{
		value: features,
	}
	return feature
}
