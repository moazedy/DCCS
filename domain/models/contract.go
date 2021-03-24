package models

import (
	"github.com/google/uuid"
)

type Contract struct {
	Id                         uuid.UUID     `json:"id"`
	Title                      string        `json:"title"`
	GovernmentContract         bool          `json:"government_contract"`
	Complex                    bool          `josn:"complex"`
	TimeRange                  TimeRangeType `json:"time_range"`
	StagedExecution            bool          `json:"staged_execution"`
	SpecialLegalRequirements   bool          `json:"special_legal_requirements"`
	PaymentType                string        `json:"payment_type"`
	DesignWithContactor        bool          `json:"desing_with_constractor"`
	IntegratedContract         bool          `json:"integrated_contract"`
	RiskManagementDone         bool          `json:"risk_management_done"`
	CounselorManagerNeeded     bool          `json:"conselor_manager_needed"`
	OwnersMainFactor           string        `json:"owners_main_factor"`
	ContractEvidencesAvailable bool          `json:"cotract_evidence_available"`
	InsuranceOnContractor      bool          `json:"insurance_on_contractor"`
	DesignManagerNeeded        bool          `json:"design_manager_needed"`
	DeliverType                string        `json:"deliver_type"`
	DeliverToContractorStage   uint          `josn:"deliver_to_contractor_stage"`
	DeliverToOwnerStage        uint          `json:"deliver_to_owner_stage"`
	ProcurementOnContractor    bool          `json:"procurement_on_contractor"`
}

type TimeRangeType struct {
	Lower uint `json:"lower"`
	Upper uint `json:"upper"`
}
