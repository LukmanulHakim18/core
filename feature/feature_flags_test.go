package feature

import (
	"reflect"
	"testing"
)

func TestNewEnabledFeatures(t *testing.T) {
	type args struct {
		features ConfigValue
	}
	tests := []struct {
		name string
		args args
		want EnabledFeature
	}{
		{
			name: "success",
			args: args{
				features: BBCorporate | CC | ECV,
			},
			want: EnabledFeature{
				value: BBCorporate | CC | ECV,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEnabledFeatures(tt.args.features); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEnabledFeatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnabledFeature_MarshalJSON(t *testing.T) {
	type fields struct {
		value ConfigValue
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				value: 100,
			},
			want:    []byte{49, 48, 48},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := EnabledFeature{
				value: tt.fields.value,
			}
			got, err := tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("EnabledFeature.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnabledFeature.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnabledFeature_GetList(t *testing.T) {
	type fields struct {
		value ConfigValue
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]bool
	}{
		{
			name: "success",
			fields: fields{
				value: NewEnabledFeatures(BBCorporate | CC | ECVRules).value,
			},
			want: map[string]bool{
				"TestServiceType":      false,
				"ECVRules":             true,
				"TripVoucher":          false,
				"CC":                   true,
				"Gopay":                false,
				"SilverbirdChartered":  false,
				"ShopeePay":            false,
				"NewPromoEngine":       false,
				"BBCorporate":          true,
				"ECV":                  false,
				"FixedFareDiscount":    false,
				"EstimateFareDiscount": false,
				"PreAuthCC":            false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ef := &EnabledFeature{
				value: tt.fields.value,
			}
			if got := ef.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnabledFeature.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}
