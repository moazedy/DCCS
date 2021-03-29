package utils

import (
	"DCCS/domain/models"
	"reflect"
	"strconv"
)

func convertContractToContractSerialDataTemplate(con models.Contract) models.ContractSerialDataTemplate {
	return models.ContractSerialDataTemplate{
		GovernmentContract:             con.GovernmentContract,
		PrivateContract:                con.PrivateContract,
		Complex:                        con.Complex,
		StagedExecution:                con.StagedExecution,
		SpecialLegalRequirements:       con.SpecialLegalRequirements,
		FixedPricePaymentType:          con.FixedPricePaymentType,
		CostPlusWagesPaymentType:       con.CostPlusWagesPaymentType,
		DesignWithContactor:            con.DesignWithContactor,
		IntegratedContract:             con.IntegratedContract,
		RiskManagementDone:             con.RiskManagementDone,
		CounselorManagerNeeded:         con.CounselorManagerNeeded,
		OwnersMainFactorIsTime:         con.OwnersMainFactorIsTime,
		OwnersMainFactorInCost:         con.OwnersMainFactorInCost,
		ContractEvidencesAvailable:     con.ContractEvidencesAvailable,
		InsuranceOnContractor:          con.InsuranceOnContractor,
		DesignManagerNeeded:            con.DesignManagerNeeded,
		SimpleDeliverType:              con.SimpleDeliverType,
		MaintainingDeliverType:         con.MaintainingDeliverType,
		DeliverToContractorBeforDesign: con.DeliverToContractorBeforDesign,
		DeliverToContractorAfterDesign: con.DeliverToContractorAfterDesign,
		ProcurementOnContractor:        con.ProcurementOnContractor,
	}
}

func ContractToSerialCode(c models.Contract) string {
	serialDataTemplate := convertContractToContractSerialDataTemplate(c)
	return TemplateToSerialCode(serialDataTemplate)
}

func TemplateToSerialCode(serialTemplate models.ContractSerialDataTemplate) string {
	var codedStruct string
	values := reflect.ValueOf(serialTemplate)

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface().(bool) {
			codedStruct = codedStruct + "1"
		} else {
			codedStruct = codedStruct + "0"
		}
	}

	return codedStruct

}

func SerialCodeToIntValue(serial string) (*int64, error) {
	res, err := strconv.ParseInt(serial, 2, 64)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
