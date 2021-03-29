package models

import (
	"github.com/google/uuid"
)

type Contract struct {
	Id                             uuid.UUID `json:"id,omitempty"`
	Title                          string    `json:"title"`
	SerialValue                    string    `json:"serial_value,omitempty"`
	IntValue                       int64     `json:"int_value,omitempty"`
	GovernmentContract             bool      `json:"government_contract"`
	PrivateContract                bool      `json:private_contract`
	Complex                        bool      `josn:"complex"`
	StagedExecution                bool      `json:"staged_execution"`
	SpecialLegalRequirements       bool      `json:"special_legal_requirements"`
	FixedPricePaymentType          bool      `json:"payment_type"`
	CostPlusWagesPaymentType       bool      `json:"cost_plus_wages"`
	DesignWithContactor            bool      `json:"desing_with_constractor"`
	IntegratedContract             bool      `json:"integrated_contract"`
	RiskManagementDone             bool      `json:"risk_management_done"`
	CounselorManagerNeeded         bool      `json:"conselor_manager_needed"`
	OwnersMainFactorIsTime         bool      `json:"owners_main_factor_is_time"`
	OwnersMainFactorInCost         bool      `json:"owners_main_factor_is_cost"`
	ContractEvidencesAvailable     bool      `json:"cotract_evidence_available"`
	InsuranceOnContractor          bool      `json:"insurance_on_contractor"`
	DesignManagerNeeded            bool      `json:"design_manager_needed"`
	SimpleDeliverType              bool      `json:"simple_deliver_type"`
	MaintainingDeliverType         bool      `json:"maintaining_deliver_type"`
	DeliverToContractorBeforDesign bool      `josn:"deliver_to_contractor_befor_design"`
	DeliverToContractorAfterDesign bool      `json:deliver_to_contractor_after_design`
	ProcurementOnContractor        bool      `json:"procurement_on_contractor"`
}

type ContractSerialDataTemplate struct {
	GovernmentContract             bool `json:"government_contract"`
	PrivateContract                bool `json:private_contract`
	Complex                        bool `josn:"complex"`
	StagedExecution                bool `json:"staged_execution"`
	SpecialLegalRequirements       bool `json:"special_legal_requirements"`
	FixedPricePaymentType          bool `json:"payment_type"`
	CostPlusWagesPaymentType       bool `json:"cost_plus_wages"`
	DesignWithContactor            bool `json:"desing_with_constractor"`
	IntegratedContract             bool `json:"integrated_contract"`
	RiskManagementDone             bool `json:"risk_management_done"`
	CounselorManagerNeeded         bool `json:"conselor_manager_needed"`
	OwnersMainFactorIsTime         bool `json:"owners_main_factor_is_time"`
	OwnersMainFactorInCost         bool `json:"owners_main_factor_is_cost"`
	ContractEvidencesAvailable     bool `json:"cotract_evidence_available"`
	InsuranceOnContractor          bool `json:"insurance_on_contractor"`
	DesignManagerNeeded            bool `json:"design_manager_needed"`
	SimpleDeliverType              bool `json:"simple_deliver_type"`
	MaintainingDeliverType         bool `json:"maintaining_deliver_type"`
	DeliverToContractorBeforDesign bool `josn:"deliver_to_contractor_befor_design"`
	DeliverToContractorAfterDesign bool `json:deliver_to_contractor_after_design`
	ProcurementOnContractor        bool `json:"procurement_on_contractor"`
}
