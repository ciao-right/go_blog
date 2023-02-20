package utils

import (
	"go_blog/model/request/riskBaseConfig"
	"testing"
)

var test_struct riskBaseConfig.RiskAreaConfig

func Test_GetStructLabel(t *testing.T) {
	t.Run("test", func(tt *testing.T) {
		GetStructLabel(test_struct)
	})
}

var testStruct2 = []riskBaseConfig.RiskAreaConfig{
	{
		AreaId:                2,
		RiskLevel:             3,
		ResponsibleDepartment: 4,
		ResponsiblePerson:     5,
		Remark:                "test",
	},
}

func Test_MakeExcel(t *testing.T) {
	t.Run("test", func(tt *testing.T) {
		MakeExcel(testStruct2)
	})
}
