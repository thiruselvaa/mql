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
	"memberships": {
		"array": [
		  {
			"active": true
		  }
		]
	},
	"medicareEntitlement": {
		"array": [
			{
				"effectiveDate": {
					"string": "2021-01-01"
				}
			}
		]
	},
	"security": {
		"com.optum.exts.eligibility.model.common.Security": {
			"securityPermissionInt": {
				"array": [
					{
						"securityPermissionValue": {
							"int": 0
						}
					},
					{
						"securityPermissionValue": {
							"int": 1
						}
					},
					{
						"securityPermissionValue": {
							"int": 2
						}
					}
				]
			},
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
							"string": "2"
						}
					}
				]
			},
			"securityPermissionAny": null,
			"securitySourceSystemCode": {
				"string": "cdb"
			}
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
}`

// 	valueJsonStr = `{
// 	"individualIdentifier": {
// 		"string": "cdb:4:144667964:CO:RAM0507494677209000"
// 	},
// 	"security": {
// 		"com.optum.exts.eligibility.model.common.Security": {
// 			"securityPermission": {
// 				"array": [
// 					{
// 						"securityPermissionValue": {
// 							"string": "0"
// 						}
// 					},
// 					{
// 						"securityPermissionValue": {
// 							"string": "1"
// 						}
// 					},
// 					{
// 						"securityPermissionValue": {
// 							"string": 2 //THIS float64 value MIXED up in ARRAY with string data type with other 2 values above is failing parser
// 						}
// 					}
// 				]
// 			},
// 			"securityPermissionAny": null,
// 			"securitySourceSystemCode": {
// 				"string": "cdb"
// 			},
// 			"securityAlt1SourceSystemCode": {
// 				"string": "CDB"
// 			},
// 			"securityAlt2SourceSystemCode": {
// 				"string": "cdb  "
// 			},
// 			"securityAlt3SourceSystemCode": {
// 				"string": "  cdb"
// 			},
// 			"securityAlt4SourceSystemCode": {
// 				"string": "  cdb  "
// 			}
// 		}
// 	}
// }`

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
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][:] ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][2].securityPermissionValue.string ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][1]["securityPermissionValue"]["string"] ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][-1]["securityPermissionValue"]["string"] ?? "nodata"`,

				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][1]["securityPermissionValue"]["string"] in ["0", "1", "2"]`,
				// expression: `any(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], {#["securityPermissionValue"]["string"] in ["0", "1", "2"]})`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], #["securityPermissionValue"]["string"] in ["0", "1", "2"])`,

				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], {#["securityPermissionValue"]["string"] in [0, 1, 2]})`, //won't work
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], .["securityPermissionValue"]["string"] in ["0", "1", "2"])`, //
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], #["securityPermissionValue"]["string"] in ["0", "1", "2"])`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], {#["securityPermissionValue"]["string"] in ["0", "1", "2"]})`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], {#["securityPermissionValue"]["string"] in ["0", "1", "2"]})`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], {#["securityPermissionValue"]["string"] in [0, 1, 2]})`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], {#["securityPermissionValue"]["string"] in [float("0"), float("1"), float("2")]})`,

				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermissionInt"]["array"], {#["securityPermissionValue"]["int"] in ["0", "1", "2"]})`, //won't work
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermissionInt"]["array"], {#["securityPermissionValue"]["int"] in list[0, 1, 2]})`,   //won't work
				// expression: `all(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int in [2, 5, 0, 3, 1, 7])`, //better option
				// expression: `all(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int in 0..3)`, //better option
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermissionInt"]["array"], .securityPermissionValue.int in 0..3)`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermissionInt"]["array"], #["securityPermissionValue"]["int"] in 0..3)`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermissionInt"]["array"], {#["securityPermissionValue"]["int"] in 0..3})`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermissionInt"]["array"], {#["securityPermissionValue"]["int"] in 0..10})`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermissionInt"]["array"], {#["securityPermissionValue"]["int"] in [0, 1, 2]})`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermissionInt"]["array"], {#["securityPermissionValue"]["int"] in [int("0"), int("1"), int("2")]})`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermissionInt"]["array"], {#["securityPermissionValue"]["int"] in [float("0"), float("1"), float("2")]})`,

				// expression: `filter(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int > 0)`, //better option
				// expression: `map(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int > 0)`, //better option
				// expression: `map(filter(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int > 0), int(.securityPermissionValue.int))`,
				// expression: `map(filter(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int > 0), int(.securityPermissionValue.int)) == [1,2]`, //returns true

				// expression: `map(filter(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int > 0), int(.securityPermissionValue.int)) in [1,2]`, //returns false
				// expression: `map(filter(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int > 0), float(.securityPermissionValue.int)) == [1,2]`, //returns false
				// expression: `map(filter(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int > 0), .securityPermissionValue.int) == [1,2]`, //returns false
				// expression: `map(filter(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int > 0), .securityPermissionValue.int) in [1,2]`, //returns false

				// expression: `message.medicareEntitlement?.array ?? "not-found"`,
				// expression: `all(message.medicareEntitlement.array, .effectiveDate.string > '2020-01-01')`,
				// expression: `all(message.medicareEntitlement.array, .effectiveDate.string > "2020-01-01")`,
				// expression: `any(message.medicareEntitlement.array, .effectiveDate.string > "2020-01-01")`,
				// expression: `any(message.medicareEntitlement.array, .effectiveDate.string >= "2021-01-01")`,
				// expression: `any(message.medicareEntitlement.array, .effectiveDate.string <= "2021-01-01")`,
				// expression: `any(message.medicareEntitlement.array, .effectiveDate.string < "2022-01-01")`,
				// expression: `any(message.medicareEntitlement.array, .effectiveDate.string < "2021-02-01")`,
				// expression: `any(message.medicareEntitlement.array, .effectiveDate.string != "2021-02-01")`,
				// expression: `any(message.medicareEntitlement.array, .effectiveDate.string != "2021-01-01")`, //returns false

				// expression: `message.memberships?.array ?? "not-found"`,
				// expression: `any(message.memberships.array, .active == true)`,
				// expression: `all(message.memberships.array, .active == true)`,
				// expression: `one(message.memberships.array, .active == true)`,

				// expression: `message.security.securityAlt1SourceSystemCode.string == 'cdb'`, //returns false
				// expression: `message.security.securityAlt1SourceSystemCode.string == 'CDB'`,
				// expression: `message.security.securityAlt1SourceSystemCode.string matches '(?i)cdb'`,
				// expression: `message.security.securityAlt1SourceSystemCode.string matches '(?i)cDB'`,

				// expression: `message.security.securityAlt2SourceSystemCode.string matches '(?i)cdb'`,
				// expression: `message.security.securityAlt3SourceSystemCode.string matches '(?i)cdb'`,
				// expression: `message.security.securityAlt4SourceSystemCode.string matches '(?i)cdb'`,

				// expression: `message.security.securityAlt4SourceSystemCode.string contains 'db'`,
				// expression: `!(message.security.securityAlt4SourceSystemCode.string contains 'cb')`,
				// expression: `!(message.security.securityAlt4SourceSystemCode.string contains 'db')`, //returns false
				// expression: `not(message.security.securityAlt4SourceSystemCode.string contains 'db')`, //returns false
				// expression: `!message.security.securityAlt4SourceSystemCode.string contains 'db'`, //won't work - unless parenthis are used
				// expression: `message.security.securityAlt4SourceSystemCode.string !contains 'db'`, //won't work

				// expression: `message.security.securityAlt1SourceSystemCode.string startsWith 'CD'`,
				// expression: `message.security.securityAlt1SourceSystemCode.string endsWith 'DB'`,
				// expression: `message.security.securityAlt2SourceSystemCode.string matches '^[ \t]+cd{1}'`, //returns false
				// expression: `message.security.securityAlt2SourceSystemCode.string matches 'cd{1}'`,
				// expression: `message.security.securityAlt2SourceSystemCode.string matches '^cd{1}'`,
				// expression: `message.security.securityAlt2SourceSystemCode.string matches '^cd{1}'`,
				// expression: `message.security.securityAlt3SourceSystemCode.string matches '^[ \t]+cd{1}'`,
				// expression: `message.security.securityAlt4SourceSystemCode.string matches '^[ \t]+cd{1}'`,

				// expression: `message.security.securityAlt1SourceSystemCode.string matches '(?i)db{1}$'`,
				// expression: `message.security.securityAlt2SourceSystemCode.string matches 'db{1}[ \t]+$'`,
				// expression: `message.security.securityAlt3SourceSystemCode.string matches 'db{1}$'`,
				// expression: `message.security.securityAlt4SourceSystemCode.string matches 'db{1}'`,
				expression: `message.security.securityAlt4SourceSystemCode.string matches 'db{1}[ \t]+$'`,

				// expression: `message.security.securityAlt3SourceSystemCode.string startsWith 'cd'`, //won't work since it has space padding at beginning of the word
				// expression: `map(filter(message.security.securityAlt3SourceSystemCode.string != nil), message.security.securityAlt3SourceSystemCode.string`, //never works, totally incorrect syntax

				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][]["securityPermissionValue"]["string"] ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"]#["securityPermissionValue"]["string"] ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"].#.["securityPermissionValue"]["string"] ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"].{#["securityPermissionValue"]["string"]} ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][:]["securityPermissionValue"]["string"] ?? "nodata"`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][:3]["securityPermissionValue"]["string"] ?? "nodata"`,

				// expression: `any(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], {#["securityPermissionValue"] in [0, 1]})`,
				// expression: `any(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], {#["securityPermissionValue"]["string"] in [0, 1, 2]})`,
				// expression: `message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"][1]["securityPermissionValue"]["string"] in [0, 1, 2]`,
				// expression: `all(float(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"]), {#["securityPermissionValue"]["string"] in ["0", "1", "2"]})`,
				// expression: `all(message["security"]["com.optum.exts.eligibility.model.common.Security"]["securityPermission"]["array"], {#["securityPermissionValue"]["string"] in [float(0), float(1), float(2)]})`,

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
			fmt.Printf("mql() output type(%T)= value(%v), want %v, err = %v\n", gotResult, gotResult, tt.wantResult, err)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("mql() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
