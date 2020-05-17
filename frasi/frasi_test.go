package frasi

import (
	"reflect"
	"testing"
)

func TestGetPeriodi(t *testing.T) {
	type args struct {
		testo string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"primo", args{testo: "Ciao, come va? Tutto bene."}, []string{"Ciao", "come va", "Tutto bene", ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPeriodi(tt.args.testo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeriodi() = %v, want %v", got, tt.want)
			}
		})
	}
}
