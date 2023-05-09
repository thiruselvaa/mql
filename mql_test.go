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
				"string": "H0001"
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
							"string": "1"
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
				"string": "H0002"
			},
			"packageBenefitPlanCode": {
				"string": "001"
			},
			"segmentId": {
				"string": "2"
			},
			"effectiveDate": {
				"string": "2022-12-31"
			}
		  },
		  {
			"active": true,
			"hContractId": 
			{
				"string": "H0003"
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
)

func Test_mql(t *testing.T) {
	var valueJsonMap map[string]interface{}
	err := jsonutil.DecodeString(valueJsonStr, &valueJsonMap)
	if err != nil {
		fmt.Printf("unable to decode the json string: %v\n", err)
	}
	dump.V(valueJsonMap)

	type Env struct {
		Message map[string]interface{} `expr:"message"`
	}

	configFile := "configs/dsl/solutran/json/solutran-dsl-filter-config.json"
	smfConfig, err := models.NewDSLFilterConfig(configFile)
	if err != nil {
		fmt.Printf("error parsing smf config file: %v", err)
		return
	}
	fmt.Printf("\nsmfConfig: %v\n", smfConfig)

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
			name: "dsl test",
			args: args{
				expression: smfConfig.Filter.Condition.String(),
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
			fmt.Printf("\nResult: type(%T) value(%v)\n", gotResult, gotResult)

			var value []byte
			value, err = jsonutil.EncodePretty(gotResult)
			if err != nil {
				fmt.Printf("unable to decode the json string: %v\n", err)
			}
			dump.V(string(value))

			fmt.Printf("mql() output type(%T), value(%#v), want %v, err = %v\n", gotResult, gotResult, tt.wantResult, err)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("mql() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
