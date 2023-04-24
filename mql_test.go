package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/jsonutil"
	"github.com/thiruselvaa/mql/models"
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
				"string": "001"
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
				"string": "002"
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

	// configFile := "configs/native-filter-query.json"
	// configFile := "configs/native/native-filter-query.yaml"
	// configFile := "configs/native/test-filter-query.yaml"
	// smfConfig, err := models.NewSMFConfig(configFile)

	configFile := "configs/dsl/solutran/json/solutran-dsl-filter-config.json"
	// smfConfig, err := models.NewDSLFilterConfig(configFile)
	_, err = models.NewDSLFilterConfig(configFile)
	if err != nil {
		fmt.Printf("error parsing smf config file: %v", err)
		return
	}

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
			name: "isSourceCode_equals_to_cdb",
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

				// expression: `
				// 	message.value.security.securityAlt1SourceSystemCode.string == 'CDB'
				// 	and
				// 	all(
				// 		message.value.security["com.optum.exts.eligibility.model.common.Security"].securityPermissionInt.array,
				// 		.securityPermissionValue.int in [2, 5, 0, 3, 1, 7]
				// 	)
				// `,

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

				// expression: `
				// 	any(message.memberships.array,
				// 			len(message.memberships.array) > 0
				// 		and .hContractId.string in ['H2226', 'R7444']
				// 		and .packageBenefitPlanCode.string in ["001", "002", "003", "004"]
				// 		and .segmentId.string in ["null", "001", "002", "003", "004"]
				// 		and (
				// 			   .effectiveDate.string > '2020-12-31'
				// 			or .effectiveDate.string > '2021-12-31'
				// 			or .effectiveDate.string > '2022-12-31'
				// 		)
				// 	)
				// `,

				// expression: `
				// 	any(message.memberships.array,
				// 			len(message.memberships.array) > 0
				// 		and .hContractId.string in ['R7444', 'R6801', 'R5342', 'R5329', 'R3444', 'R3175', 'R2604', 'R1548', 'R0759', 'H8768', 'H8748', 'H8211', 'H8125', 'H7778', 'H7464', 'H7445', 'H7404', 'H6706', 'H6595', 'H6528', 'H6526', 'H5652', 'H5435', 'H5420', 'H5322', 'H5253', 'H5008', 'H4829', 'H4604', 'H4590', 'H4527', 'H4514', 'H4094', 'H3805', 'H3794', 'H3749', 'H3464', 'H3442', 'H3418', 'H3387', 'H3379', 'H3307', 'H3256', 'H3113', 'H2802', 'H2582', 'H2577', 'H2509', 'H2406', 'H2292', 'H2272', 'H2247', 'H2228', 'H2226', 'H2196', 'H2001', 'H1944', 'H1889', 'H1821', 'H1659', 'H1375', 'H1360', 'H1278', 'H1111', 'H1045', 'H0764', 'H0755', 'H0710', 'H0624', 'H0609', 'H0543', 'H0432', 'H0321', 'H0294', 'H0271', 'H0251', 'H0169']
				// 		and .packageBenefitPlanCode.string in ["001", "002", "003", "004", "005", "006", "007", "008", "009", "010", "011", "012", "013", "014", "015", "016", "017", "018", "019", "020", "021", "022", "023", "024", "025", "026", "027", "028", "029", "030", "031", "032", "033", "034", "035", "036", "037", "038", "039", "040", "041", "042", "043", "044", "045", "046", "047", "048", "049", "050", "051", "052", "053", "054", "055", "056", "057", "058", "059", "060", "061", "062", "063", "064", "065", "066", "067", "068", "069", "070", "071", "072", "073", "074", "075", "076", "077", "079", "080", "081", "082", "083", "084", "085", "086", "087", "088", "089", "091", "092", "093", "094", "095", "096", "097", "098", "099", "100", "101", "102", "103", "104", "105", "107", "108", "109", "110", "111", "112", "113", "115", "116", "117", "118", "119", "120", "121", "122", "124", "125", "126", "127", "128", "129", "130", "131", "140", "145", "146", "147", "148", "151", "152", "153", "163", "164", "165", "166", "167", "168", "169", "170", "172", "173", "176", "177", "178", "179", "183", "185", "188", "189", "191", "193", "195", "196", "202", "204", "205", "209", "210", "214", "215", "216", "217", "218", "219", "220", "221", "222", "223", "224", "225", "226", "227", "228", "229", "230", "231", "232", "234", "235", "236", "237", "238", "239", "240", "241", "242", "243", "817", "859", "868"]
				// 		and .segmentId.string in ["null", "001", "002", "003", "004"]
				// 		and (
				// 			   .effectiveDate.string > '2020-12-31'
				// 			or .effectiveDate.string > '2021-12-31'
				// 			or .effectiveDate.string > '2022-12-31'
				// 		)
				// 	)
				// `,

				// expression: `message.individualIdentifier`,
				// expression: `len(message.individualIdentifier)`,
				// expression: `message.individualIdentifier.string`,
				// expression: `len(message.individualIdentifier.string)`,
				// expression: "len(message.individualIdentifier.string)",

				// expression: `len(message.memberships.array) > 0`,
				// expression: `len(message.memberships.array) > 1.0`,
				// expression: `len(message.memberships.array) > int("1")`,
				// expression: `len(message.memberships.array) > float("1")`,
				// expression: `len(message.memberships.array) > float("1.0")`,

				// expression: `len(message.memberships.array) > "1"`,        //don't work
				// expression: `len(message.memberships.array) > int("1.0")`, //don't work

				//
				// expression: `
				// 	len(message.memberships.array) > 0
				// 	and
				// 		any(message.memberships.array,
				// 				.hContractId.string in ['R7444', 'R6801', 'R5342', 'R5329', 'R3444', 'R3175', 'R2604', 'R1548', 'R0759', 'H8768', 'H8748', 'H8211', 'H8125', 'H7778', 'H7464', 'H7445', 'H7404', 'H6706', 'H6595', 'H6528', 'H6526', 'H5652', 'H5435', 'H5420', 'H5322', 'H5253', 'H5008', 'H4829', 'H4604', 'H4590', 'H4527', 'H4514', 'H4094', 'H3805', 'H3794', 'H3749', 'H3464', 'H3442', 'H3418', 'H3387', 'H3379', 'H3307', 'H3256', 'H3113', 'H2802', 'H2582', 'H2577', 'H2509', 'H2406', 'H2292', 'H2272', 'H2247', 'H2228', 'H2226', 'H2196', 'H2001', 'H1944', 'H1889', 'H1821', 'H1659', 'H1375', 'H1360', 'H1278', 'H1111', 'H1045', 'H0764', 'H0755', 'H0710', 'H0624', 'H0609', 'H0543', 'H0432', 'H0321', 'H0294', 'H0271', 'H0251', 'H0169']
				// 			and .packageBenefitPlanCode.string in ["001", "002", "003", "004", "005", "006", "007", "008", "009", "010", "011", "012", "013", "014", "015", "016", "017", "018", "019", "020", "021", "022", "023", "024", "025", "026", "027", "028", "029", "030", "031", "032", "033", "034", "035", "036", "037", "038", "039", "040", "041", "042", "043", "044", "045", "046", "047", "048", "049", "050", "051", "052", "053", "054", "055", "056", "057", "058", "059", "060", "061", "062", "063", "064", "065", "066", "067", "068", "069", "070", "071", "072", "073", "074", "075", "076", "077", "079", "080", "081", "082", "083", "084", "085", "086", "087", "088", "089", "091", "092", "093", "094", "095", "096", "097", "098", "099", "100", "101", "102", "103", "104", "105", "107", "108", "109", "110", "111", "112", "113", "115", "116", "117", "118", "119", "120", "121", "122", "124", "125", "126", "127", "128", "129", "130", "131", "140", "145", "146", "147", "148", "151", "152", "153", "163", "164", "165", "166", "167", "168", "169", "170", "172", "173", "176", "177", "178", "179", "183", "185", "188", "189", "191", "193", "195", "196", "202", "204", "205", "209", "210", "214", "215", "216", "217", "218", "219", "220", "221", "222", "223", "224", "225", "226", "227", "228", "229", "230", "231", "232", "234", "235", "236", "237", "238", "239", "240", "241", "242", "243", "817", "859", "868"]
				// 			and .segmentId.string in ["null", "001", "002", "003", "004"]
				// 			and
				// 				any(.membershipGroupData.array,
				// 					.groupNumber.string in ["", "12345", "100", "", '97008', '97007', '97006', '97005', '97004', '97003', '12830', 'null', '*']
				// 				)
				// 			and (
				// 				.effectiveDate.string >= '2020-12-31' and .effectiveDate.string <= '2022-12-31'
				// 			)
				// 		)
				// `,
				//

				// expression: `
				// message.security["com.optum.exts.eligibility.model.common.Security"].securitySourceSystemCode.string == "cdb"
				// and
				// len(message.memberships.array) > 0
				// and
				// 	any(message.memberships.array[:],
				// 			#.hContractId.string in ['R7444', 'R6801', 'R5342', 'R5329', 'R3444', 'R3175', 'R2604', 'R1548', 'R0759', 'H8768', 'H8748', 'H8211', 'H8125', 'H7778', 'H7464', 'H7445', 'H7404', 'H6706', 'H6595', 'H6528', 'H6526', 'H5652', 'H5435', 'H5420', 'H5322', 'H5253', 'H5008', 'H4829', 'H4604', 'H4590', 'H4527', 'H4514', 'H4094', 'H3805', 'H3794', 'H3749', 'H3464', 'H3442', 'H3418', 'H3387', 'H3379', 'H3307', 'H3256', 'H3113', 'H2802', 'H2582', 'H2577', 'H2509', 'H2406', 'H2292', 'H2272', 'H2247', 'H2228', 'H2226', 'H2196', 'H2001', 'H1944', 'H1889', 'H1821', 'H1659', 'H1375', 'H1360', 'H1278', 'H1111', 'H1045', 'H0764', 'H0755', 'H0710', 'H0624', 'H0609', 'H0543', 'H0432', 'H0321', 'H0294', 'H0271', 'H0251', 'H0169']
				// 		and #.packageBenefitPlanCode.string in ["001", "002", "003", "004", "005", "006", "007", "008", "009", "010", "011", "012", "013", "014", "015", "016", "017", "018", "019", "020", "021", "022", "023", "024", "025", "026", "027", "028", "029", "030", "031", "032", "033", "034", "035", "036", "037", "038", "039", "040", "041", "042", "043", "044", "045", "046", "047", "048", "049", "050", "051", "052", "053", "054", "055", "056", "057", "058", "059", "060", "061", "062", "063", "064", "065", "066", "067", "068", "069", "070", "071", "072", "073", "074", "075", "076", "077", "079", "080", "081", "082", "083", "084", "085", "086", "087", "088", "089", "091", "092", "093", "094", "095", "096", "097", "098", "099", "100", "101", "102", "103", "104", "105", "107", "108", "109", "110", "111", "112", "113", "115", "116", "117", "118", "119", "120", "121", "122", "124", "125", "126", "127", "128", "129", "130", "131", "140", "145", "146", "147", "148", "151", "152", "153", "163", "164", "165", "166", "167", "168", "169", "170", "172", "173", "176", "177", "178", "179", "183", "185", "188", "189", "191", "193", "195", "196", "202", "204", "205", "209", "210", "214", "215", "216", "217", "218", "219", "220", "221", "222", "223", "224", "225", "226", "227", "228", "229", "230", "231", "232", "234", "235", "236", "237", "238", "239", "240", "241", "242", "243", "817", "859", "868"]
				// 		and #.segmentId.string in ["null", "001", "002", "003", "004"]
				// 		and
				// 			any(#.membershipGroupData.array[:],
				// 				#.groupNumber.string in ["", "12345", "100", "", '97008', '97007', '97006', '97005', '97004', '97003', '12830', 'null', '*']
				// 			)
				// 		and (
				// 			#.effectiveDate.string >= '2020-12-31' and #.effectiveDate.string <= '2022-12-31'
				// 		)
				// 	)
				// `,

				// expression: `
				// len(message.memberships.array) > 0
				// and
				// 	any(message.memberships.array[:],
				// 			#.hContractId.string in ['R7444', 'R6801', 'R5342', 'R5329', 'R3444', 'R3175', 'R2604', 'R1548', 'R0759', 'H8768', 'H8748', 'H8211', 'H8125', 'H7778', 'H7464', 'H7445', 'H7404', 'H6706', 'H6595', 'H6528', 'H6526', 'H5652', 'H5435', 'H5420', 'H5322', 'H5253', 'H5008', 'H4829', 'H4604', 'H4590', 'H4527', 'H4514', 'H4094', 'H3805', 'H3794', 'H3749', 'H3464', 'H3442', 'H3418', 'H3387', 'H3379', 'H3307', 'H3256', 'H3113', 'H2802', 'H2582', 'H2577', 'H2509', 'H2406', 'H2292', 'H2272', 'H2247', 'H2228', 'H2226', 'H2196', 'H2001', 'H1944', 'H1889', 'H1821', 'H1659', 'H1375', 'H1360', 'H1278', 'H1111', 'H1045', 'H0764', 'H0755', 'H0710', 'H0624', 'H0609', 'H0543', 'H0432', 'H0321', 'H0294', 'H0271', 'H0251', 'H0169']
				// 		and #.packageBenefitPlanCode.string in ["001", "002", "003", "004", "005", "006", "007", "008", "009", "010", "011", "012", "013", "014", "015", "016", "017", "018", "019", "020", "021", "022", "023", "024", "025", "026", "027", "028", "029", "030", "031", "032", "033", "034", "035", "036", "037", "038", "039", "040", "041", "042", "043", "044", "045", "046", "047", "048", "049", "050", "051", "052", "053", "054", "055", "056", "057", "058", "059", "060", "061", "062", "063", "064", "065", "066", "067", "068", "069", "070", "071", "072", "073", "074", "075", "076", "077", "079", "080", "081", "082", "083", "084", "085", "086", "087", "088", "089", "091", "092", "093", "094", "095", "096", "097", "098", "099", "100", "101", "102", "103", "104", "105", "107", "108", "109", "110", "111", "112", "113", "115", "116", "117", "118", "119", "120", "121", "122", "124", "125", "126", "127", "128", "129", "130", "131", "140", "145", "146", "147", "148", "151", "152", "153", "163", "164", "165", "166", "167", "168", "169", "170", "172", "173", "176", "177", "178", "179", "183", "185", "188", "189", "191", "193", "195", "196", "202", "204", "205", "209", "210", "214", "215", "216", "217", "218", "219", "220", "221", "222", "223", "224", "225", "226", "227", "228", "229", "230", "231", "232", "234", "235", "236", "237", "238", "239", "240", "241", "242", "243", "817", "859", "868"]
				// 		and #.segmentId.string in ["null", "001", "002", "003", "004"]
				// 		and
				// 			any(#.membershipGroupData.array[:],
				// 				#.groupNumber.string in ["", "12345", "100", "", '97008', '97007', '97006', '97005', '97004', '97003', '12830', 'null', '*']
				// 			)
				// 		and (
				// 			#.effectiveDate.string >= '2020-12-31' and #.effectiveDate.string <= '2022-12-31'
				// 		)
				// 	)
				// `,

				// expression: `
				// (((len(message.memberships.array) > 0)
				// and (
				// 	any(message.memberships.array[:],
				// 		((((#.hContractId.string in ['H2226','R7444'])))
				// 		and((((#.packageBenefitPlanCode.string in ['001','002','003','004'])
				// 		and (#.segmentId.string in ['null','001','002','003','004']))
				// 		and (((#.effectiveDate.string > '2020-12-31')) or ((#.effectiveDate.string > '2021-12-31')) or ((#.effectiveDate.string > '2022-12-31')))
				// 		and (
				// 			any(#.membershipGroupData.array[:],
				// 				(((((#.groupNumber.string in ['','12345','100','97008','97007','97006','97005','97004','97003','12830','null','*']))))))))))))))
				// `,
				// // // expression: smfConfig.Filter.Condition.String(),

				// expression: whereString(smfConfig),
				// expression: "true",
				// expression: "true or false",
				// expression: "len(message.memberships.array) > 0",
				// expression: "(len(message.memberships.array) > 0)and len(message.memberships.array) > 0",
				// expression: "(len(message.memberships.array) > 0)and(len(message.memberships.array) > 0)",
				// expression: "(len(message.memberships.array) > 0) or (((len(message.memberships.array) > 0)))",
				// expression: "true and (true)",
				// expression: "1 == 1 ",
				// expression: "1.23 == 1.23",
				// expression: "1.23==1.23",
				// expression: "-1 >= -2",
				// expression: "-1 >= -2.0",
				// expression: "-1.0 >= -2",
				// expression: "-1 <= -2",

				// expression: `
				//   (
				// 	message.security["com.optum.exts.eligibility.model.common.Security"].securitySourceSystemCode.string == 'CDB'
				//   )
				//   or
				//   (
				// 	message.security["com.optum.exts.eligibility.model.common.Security"].securitySourceSystemCode.string == 'cdb'
				//   )
				//   or
				//   (
				// 	message.security["com.optum.exts.eligibility.model.common.Security"].securitySourceSystemCode.string == 'cDB'
				//   )
				// `,

				//
				// expression: `
				// 	len(message.memberships.array) > 0
				// 	and
				// 		any(message.memberships.array,
				// 				.hContractId.string in ['R7444', 'R6801', 'R5342', 'R5329', 'R3444', 'R3175', 'R2604', 'R1548', 'R0759', 'H8768', 'H8748', 'H8211', 'H8125', 'H7778', 'H7464', 'H7445', 'H7404', 'H6706', 'H6595', 'H6528', 'H6526', 'H5652', 'H5435', 'H5420', 'H5322', 'H5253', 'H5008', 'H4829', 'H4604', 'H4590', 'H4527', 'H4514', 'H4094', 'H3805', 'H3794', 'H3749', 'H3464', 'H3442', 'H3418', 'H3387', 'H3379', 'H3307', 'H3256', 'H3113', 'H2802', 'H2582', 'H2577', 'H2509', 'H2406', 'H2292', 'H2272', 'H2247', 'H2228', 'H2226', 'H2196', 'H2001', 'H1944', 'H1889', 'H1821', 'H1659', 'H1375', 'H1360', 'H1278', 'H1111', 'H1045', 'H0764', 'H0755', 'H0710', 'H0624', 'H0609', 'H0543', 'H0432', 'H0321', 'H0294', 'H0271', 'H0251', 'H0169']
				// 			and .packageBenefitPlanCode.string in ["001", "002", "003", "004", "005", "006", "007", "008", "009", "010", "011", "012", "013", "014", "015", "016", "017", "018", "019", "020", "021", "022", "023", "024", "025", "026", "027", "028", "029", "030", "031", "032", "033", "034", "035", "036", "037", "038", "039", "040", "041", "042", "043", "044", "045", "046", "047", "048", "049", "050", "051", "052", "053", "054", "055", "056", "057", "058", "059", "060", "061", "062", "063", "064", "065", "066", "067", "068", "069", "070", "071", "072", "073", "074", "075", "076", "077", "079", "080", "081", "082", "083", "084", "085", "086", "087", "088", "089", "091", "092", "093", "094", "095", "096", "097", "098", "099", "100", "101", "102", "103", "104", "105", "107", "108", "109", "110", "111", "112", "113", "115", "116", "117", "118", "119", "120", "121", "122", "124", "125", "126", "127", "128", "129", "130", "131", "140", "145", "146", "147", "148", "151", "152", "153", "163", "164", "165", "166", "167", "168", "169", "170", "172", "173", "176", "177", "178", "179", "183", "185", "188", "189", "191", "193", "195", "196", "202", "204", "205", "209", "210", "214", "215", "216", "217", "218", "219", "220", "221", "222", "223", "224", "225", "226", "227", "228", "229", "230", "231", "232", "234", "235", "236", "237", "238", "239", "240", "241", "242", "243", "817", "859", "868"]
				// 			and .segmentId.string in ["null", "001", "002", "003", "004"]
				// 			and
				// 				any(.membershipGroupData.array,
				// 					.groupNumber.string in ["", "12345", "100", "", '97008', '97007', '97006', '97005', '97004', '97003', '12830', 'null', '*']
				// 				)
				// 			and (
				// 				.effectiveDate.string > '2020-12-31'
				// 				or .effectiveDate.string > '2021-12-31'
				// 				or .effectiveDate.string > '2022-12-31'
				// 			)
				// 		)
				// `,
				//

				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) > 0),
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

				//
				// expression: `
				// 	map(
				// 		filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
				// 		map(
				// 			filter(.membershipGroupData.array, len(.groupNumber.string) >=0),
				// 			.groupNumber.string
				// 		)
				// 	)
				// `,
				//

				expression: `
					map(
						filter(message.memberships.array, len(.membershipGroupData.array) >= 0),
						{
							.hContractId.string + "," +
							.packageBenefitPlanCode.string + "," +
							.segmentId.string + "," +
							.effectiveDate.string
						}
					)
				`,
				// expression: `map(message.memberships.array,  map(filter(.membershipGroupData.array, .groupNumber.string matches '.*'), .groupNumber.string))`,

				// expression: `len(message.memberships.array) ?? null`,
				// expression: `message.memberships.array[:] ?? nodata`,
				// expression: `message.memberships.array[0:][0]?.membershipGroupData.array[0:] ?? null`,
				// expression: `message.memberships.array[0:][0].membershipGroupData.array[0:] ?? null`,
				// expression: `message.memberships.array[0:][0:len(message.memberships.array)] ?? null`,
				// expression: `message.memberships.array[:][:][0].membershipGroupData ?? null`,
				// expression: `message.memberships.array[:][:][2].membershipGroupData.groupNumber ?? null`,
				// expression: `message.memberships.array[:][:][2].membershipGroupData.groupNumber ? null`,
				// expression: `message.memberships.array[:][1].membershipGroupData.array[0] ?? null`,
				// expression: `message.memberships.array[:][1].membershipGroupData.array[0].groupNumber.string ?? null`, //working

				// expression: `message.memberships.array[0:len(message.memberships.array)][0].membershipGroupData.array[0].groupNumber.string ?? null`,
				// expression: `message.memberships.array[0:len(message.memberships.array)]?? null`,

				// expression: `map(message.memberships.array[:], .membershipGroupData.array) ?? nodata`,
				// expression: `map(map(message.memberships.array[:], .membershipGroupData.array)[0], .groupNumber.string) ?? nodata`,
				// expression: `map(map(message.memberships.array[:], .membershipGroupData.array)[1], .groupNumber.string) ?? nodata`,
				// expression: `map(map(message.memberships.array[:], .membershipGroupData.array)[0], .groupNumber.string)[1] ?? nodata`,
				// expression: `map(message.memberships.array[:], len(.membershipGroupData.array)) ?? nodata`,
				// expression: `map(message.memberships.array[:], map(.membershipGroupData.array, len(.groupNumber.string))) ?? nodata`,
				// expression: `map(message.memberships.array, map(.membershipGroupData.array, .groupNumber.string))`,
				// expression: `map(message.memberships.array, map(filter(.membershipGroupData.array, .groupNumber.string in ["", "12345"]), .groupNumber.string))`,
				// expression: `map(message.memberships.array, map(filter(.membershipGroupData.array, .groupNumber.string in ["", "12345", "100", "null", "", "*", "12830"]), .groupNumber.string))`,

				// expression: `map(message.memberships.array, map(filter(.membershipGroupData.array, .groupNumber.string in ["", "12345", "100", "", '97008', '97007', '97006', '97005', '97004', '97003', '12830', 'null', '*']), .groupNumber.string))`,
				// expression: `any(message.memberships.array, any(.membershipGroupData.array, .groupNumber.string in ["", "12345", "100", "", '97008', '97007', '97006', '97005', '97004', '97003', '12830', 'null', '*']))`,

				// expression: `map(message.memberships.array[:], .membershipGroupData.array)[0:len(message.memberships.array)] ?? nodata`,
				// expression: `map(map(message.memberships.array[:], .membershipGroupData.array), #[0].groupNumber) ?? nodata`,

				// expression: `message.memberships.array[:][:][:][1].membershipGroupData ?? null`, //
				// expression: `message.memberships.array[0:].membershipGroupData ?? null`,
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
			// gotResult, _ := mql(tt.args.expression, tt.args.env)
			// t.Logf("result: %v", gotResult)
			gotResult, err := mql(tt.args.expression, tt.args.env)
			if err != nil && err != tt.wantErr {
				t.Errorf("mql() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// fmt.Printf("SMF Config: %v", smfConfig)

			var value []byte
			value, err = jsonutil.EncodePretty(gotResult)
			if err != nil {
				fmt.Printf("unable to decode the json string: %v\n", err)
			}
			dump.V(string(value))

			// fmt.Printf("gotResult: %#v\n\n", arrutil.AnyToString(arrutil.SliceToStrings(gotResult.([]interface{}))))

			fmt.Printf("mql() output type(%T)= value(%#v), want %v, err = %v\n", gotResult, gotResult, tt.wantResult, err)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("mql() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// func whereString(smfConfig *models.SMFConfig) string {
// 	if where, ok := smfConfig.Query.Where.(string); ok {
// 		return where
// 	}
// 	return ""
// }
