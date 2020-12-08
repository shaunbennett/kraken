package kraken

import "testing"

func TestAssetOptions_QueryString(t *testing.T) {
	type fields struct {
		Info       string
		AssetClass string
		Assets     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty options returns empty string", fields{}, ""},
		{"info is passed as a param", fields{Info: "boop"}, "info=boop"},
		{"asset class is passed as a param", fields{AssetClass: "doop"}, "aclass=doop"},
		{"asset is passed as a param", fields{Assets: []string{"woop"}}, "asset=woop"},
		{"params can be joined", fields{"boop", "doop", []string{"woop", "hoop"}}, "aclass=doop&asset=woop%2Choop&info=boop"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AssetOptions{
				Info:       tt.fields.Info,
				AssetClass: tt.fields.AssetClass,
				Assets:     tt.fields.Assets,
			}
			if got := a.QueryString(); got != tt.want {
				t.Errorf("QueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
