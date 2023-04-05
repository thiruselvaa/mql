package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/jsonutil"
)

var (
	// canonicalDocName = "coordination-of-benefits.v4"

	// keyJsonStr = `{
	// 	"individualIdentifier": {
	// 		"string": "cdb:4:144667964:CO:RAM0507494677209000"
	// 	}
	// }`

	valueJsonStr = `{
	"individualIdentifier": {
		"string": "cdb:4:144667964:CO:RAM0507494677209000"
	},
	"security": {
		"com.optum.exts.eligibility.model.common.Security": {
			"securityPermission": {
				"array": [
					{
						"securityPermissionValue": {
							"string": "0"
						}
					},
					{
						"securityPermissionValue": {
							"string": "1"
						}
					},
					{
						"securityPermissionValue": {
							"string": 2
						}
					}
				]
			},
			"securityPermissionAny": null,
			"securitySourceSystemCode": {
				"string": "cdb"
			},
			"securityAlt1SourceSystemCode": {
				"string": "CDB"
			},
			"securityAlt2SourceSystemCode": {
				"string": "cdb  "
			},
			"securityAlt3SourceSystemCode": {
				"string": "  cdb"
			},
			"securityAlt4SourceSystemCode": {
				"string": "  cdb  "
			}
		}
	}
}`

// // testJsonStr = `{
// // 	"a" : {
// // 		"string" : "1"
// // 	}
// // }`
)

func Test_mql(t *testing.T) {
	var valueJsonMap map[string]interface{}

	// err := jsonutil.DecodeString(testJsonStr, &valueJsonMap)
	// err := jsonutil.DecodeString(keyJsonStr, &valueJsonMap)
	err := jsonutil.DecodeString(valueJsonStr, &valueJsonMap)
	if err != nil {
		fmt.Printf("unable to decode the json string: %v\n", err)
	}
	dump.V(valueJsonMap)

	type Env struct {
		Message map[string]interface{} `expr:"message"`
	}

	// type Tweet struct {
	// 	Len int
	// }

	// type Env struct {
	// 	Tweets []Tweet
	// }

	type args struct {
		expression string
		env        Env
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
		wantErr    interface{}
	}{
		// TODO: Add test cases.
		{
			name: "isSourceCode equals to cdb",
			args: args{
				// expression: `any(Tweets, {.Len in [0, 1, 2, 3]})`,
				// env: Env{
				// 	Tweets: []Tweet{{1}, {10}, {11}},
				// },

				//working
				// expression: `message.individualIdentifier ?? "nodata"`,
				// expression: `message.individualIdentifier.string ?? "nodata"`,
				// expression: `message.security['com.optum.exts.eligibility.model.common.Security'].securitySourceSystemCode.string ?? "nodata"`,
				// expression: `message.security["com.optum.exts.eligibility.model.common.Security"].securitySourceSystemCode.string ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securitySourceSystemCode"]["string"] ?? "nodata"`,
				// expression: `message.security["com.optum.exts.eligibility.model.common.Security"].securitySourceSystemCode.string == "cdb"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"] ?? "nodata"`,
				expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][:] ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][:]["securityPermissionValue"] ?? "nodata"`,
				// /0/securityPermissionValue/string"

				//not-working
				// expression: `message.security."com.optum.exts.eligibility.model.common.Security".securitySourceSystemCode.string ?? "nodata"`,
				// expression: `"security.com.optum.exts.eligibility.model.common.Security.securitySourceSystemCode.string" == "cdb"`,
				// expression: `"security.com\.optum\.exts\.eligibility\.model\.common\.Security.securitySourceSystemCode.string" == "cdb"`,
				// expression: `security["com.optum.exts.eligibility.model.common.Security"]["securitySourceSystemCode"]["string"] == "cdb"`,
				env: Env{
					Message: valueJsonMap,
				},
			},
			wantResult: true,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := mql(tt.args.expression, tt.args.env)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("mql() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			fmt.Printf("mql() = %v, want %v, err = %v\n", gotResult, tt.wantResult, err)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("mql() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
