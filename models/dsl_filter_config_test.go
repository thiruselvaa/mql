package models

import (
	"testing"
)

func TestNewDSLFilterConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg *DSLFilterConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				file: "../configs/dsl/solutran/json/solutran-dsl-filter-config.json",
			},
			wantCfg: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = NewDSLFilterConfig(tt.args.file)
			// gotCfg, err := NewDSLFilterConfig(tt.args.file)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("NewDSLFilterConfig() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
			// 	t.Errorf("NewDSLFilterConfig() = %v, want %v", gotCfg, tt.wantCfg)
			// }
		})
	}
}
