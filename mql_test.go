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
			"active": true,
			"hContractId": 
			{
				"string": "H2226"
			},
			"packageBenefitPlanCode": {
				"string": "001"
			},
			"segmentId": {
				"string": "null"
			},
			"membershipGroupData": {
				"array": [
				  {
					"groupNumber": {
						"string": ""
					}
				  },
				  {
					"groupNumber": {
						"string": "100"
					}
				  }
				]
			},
			"effectiveDate": {
				"string": "2022-12-31"
			}
		  },
		  {
			"active": true,
			"hContractId": 
			{
				"string": "H2226"
			},
			"packageBenefitPlanCode": {
				"string": "002"
			},
			"segmentId": {
				"string": "null"
			},
			"membershipGroupData": {
				"array": [
					{
						"groupNumber": {
							"string": "12345"
						}
					}
				]
			},
			"effectiveDate": {
				"string": "2022-12-31"
			}
		  },
		  {
			"active": true,
			"hContractId": 
			{
				"string": "H2226"
			},
			"packageBenefitPlanCode": {
				"string": "003"
			},
			"segmentId": {
				"string": "null"
			},
			"membershipGroupData": {
				"array": [
				  
				]
			},
			"effectiveDate": {
				"string": "2022-12-31"
			}
		  },
		  {
			"active": true,
			"hContractId": 
			{
				"string": "H2226"
			},
			"packageBenefitPlanCode": {
				"string": "004"
			},
			"segmentId": {
				"string": "null"
			},
			"membershipGroupData": {
				"array": [
				  {
					"groupNumber": {
						"string": "null"
					}
				  }
				]
			},
			"effectiveDate": {
				"string": "2022-12-31"
			}
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
				// expression: `message.security.securityAlt4SourceSystemCode.string matches 'db{1}[ \t]+$'`,

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

				//solutran
				// expression: `message.memberships.array ?? null`,
				// expression: `message.memberships.array[1].hContractId.string ?? null`,
				// expression: `map(filter(message.memberships.array, .hContractId.string == "H2226"), .hContractId.string)`,
				// expression: `map(filter(message.memberships.array, len(.hContractId.string) > 0), .hContractId.string)`,
				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.hContractId.string) > 0),
				// 		.hContractId.string + "," +
				// 		.packageBenefitPlanCode.string + "," +
				// 		.segmentId.string + "," +
				// 		.effectiveDate.string
				// 	)
				// `,
				// expression: `
				// 	map(
				// 		filter(message.memberships.array,
				// 			len(message.memberships.array) > 0
				// 			and .hContractId.string in ["H2226"]
				// 			and .packageBenefitPlanCode.string in ["001","002","003","004"]
				// 			and .segmentId.string == "null"
				// 			and .effectiveDate.string > '2021-12-31'
				// 		),
				// 		.hContractId.string + "," +
				// 		.packageBenefitPlanCode.string + "," +
				// 		.segmentId.string + "," +
				// 		.effectiveDate.string
				// 	)
				// `,

				expression: `
					any(message.memberships.array, 
						len(message.memberships.array) > 0
						and .hContractId.string in ["H2226"] 
						and .packageBenefitPlanCode.string in ["001","002","003","004"]
						and .segmentId.string == "null"
						and .effectiveDate.string > '2021-12-31'
					)
				`,

				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.hContractId.string) > 0),
				// 		.hContractId.string + "," +
				// 		.packageBenefitPlanCode.string + "," +
				// 		.segmentId.string + "," +
				// 		map(
				// 			filter(.membershipGroupData.array, len(.membershipGroupData.array) > 0),
				// 			.groupNumber.string
				// 		) + "," +
				// 		.effectiveDate.string
				// 	)
				// `,

				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
				// 		map(
				// 			filter(.membershipGroupData.array, len(.groupNumber.string) >=0),
				// 			.groupNumber.string
				// 		)
				// 	)
				// `,

				// expression: `
				// // map(
				// 	// len(
				// 		map(
				// 			filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
				// 			map(
				// 				filter(.membershipGroupData.array, len(.groupNumber.string) >=0 ),
				// 				.groupNumber.string
				// 			)
				// 		)
				// 		// ) in ["null", "100", "12345"]
				// 	// )
				// 		// )[1][0]
				// 		// ),
				// 		// ?? null
				// // )
				// `,

				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
				// 		[0][1].groupNumber.string
				// 	)
				// `,

				// expression: `
				// 	any(
				// 		map(
				// 			filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
				// 			.membershipGroupData.array
				// 		),
				// 		{ [0].groupNumber.string matches '/12345/'}
				// 	)
				// `,

				// expression: `message.memberships.array[1].membershipGroupData.array[0].groupNumber.string ?? null`,
				// expression: `message.memberships.array[1].membershipGroupData.array[0].groupNumber.string ?? null`,
				// expression: `message.memberships.array[2].membershipGroupData.array[0].groupNumber.string ?? null`, //throws error array out of index
				// expression: `len(message.memberships.array[2].membershipGroupData.array) == 0`,
				// expression: `
				// 	len(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) >= 0)
				// 	)
				// `,
				// expression: `
				// 	count(message.memberships.array, len(.membershipGroupData.array) > 0)
				// `,
				// expression: `
				// 	count(message.memberships.array, len(.membershipGroupData.array) >= 0)
				// `,
				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) > 0),
				// 		len(.membershipGroupData.array)
				// 	)
				// `,
				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
				// 		len(.membershipGroupData.array)
				// 	)
				// `,
				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) > 0),
				// 		.membershipGroupData.array
				// 	)
				// `,

				// expression: `
				// // map(
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
				// 		.membershipGroupData.array
				// 	)
				// 	// .groupNumber?.string
				// `,

				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) == 0),
				// 		.groupNumber?.string
				// 	)
				// `,
				// expression: `
				// 	map(filter(
				// 			map(filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
				// 				.membershipGroupData.array),
				// 			len(.groupNumber.string) >= 0),
				// 		.groupNumber.string
				// 	)
				// `,

				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) == 0 or len(.membershipGroupData.array) > 0), "null"
				// 	)
				// `,

				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
				// 			map(
				// 				filter(.membershipGroupData.array, len(.membershipGroupData.array) > 0)
				// 				.groupNumber.string
				// 			)
				// 		)
				// 	)
				// `,

				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.hContractId.string) > 0),
				// 			map(
				// 				filter(.membershipGroupData.array, len(.membershipGroupData.array) > 0)
				// 				.groupNumber.string
				// 			)
				// 		)
				// 	)
				// `,
				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) > 0),
				// 		{
				// 			map(
				// 				filter(.membershipGroupData.array, .groupNumber.string != nil),
				// 				.groupNumber.string
				// 			)
				// 		}
				// 	)
				// `,

				// expression: `message.memberships.array[0].membershipGroupData.array ?? null`,
				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) > 0),
				// 		len(.membershipGroupData.array)
				// 	)
				// `,

				// expression: `len(message.memberships.array)`,
				// expression: `all(message.memberships.array, .hContractId.string in ["H2226", "R7444"])`,
				// expression: `all(message.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array, .securityPermissionValue.int in [2, 5, 0, 3, 1, 7])`, //better option

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
			var value []byte
			value, err = jsonutil.EncodePretty(gotResult)
			if err != nil {
				fmt.Printf("unable to decode the json string: %v\n", err)
			}
			dump.V(string(value))

			fmt.Printf("mql() output type(%T)= value(%#v), want %v, err = %v\n", gotResult, gotResult, tt.wantResult, err)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("mql() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
