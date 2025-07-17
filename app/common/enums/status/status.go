package statusenum

type Status string

var (
	Draft                       Status = "DRAFT"
	Submitted                   Status = "SUBMITTED"
	Approved                    Status = "APPROVED"
	Rejected                    Status = "REJECTED"
	Active                      Status = "ACTIVE"
	ActiveEditSubmitted         Status = "ACTIVE_EDIT_SUBMITTED"
	ActiveRejectSubmitted       Status = "ACTIVE_REJECT_SUBMITTED"
	ActiveInactivationSubmitted Status = "ACTIVE_INACTIVATION_SUBMITTED"
	Inactive                    Status = "INACTIVE"
	Unassigned                  Status = "UNASSIGNED"
	PendingApproval             Status = "PENDING_APPROVAL"
	ActivePendingApproval       Status = "ACTIVE_PENDING_APPROVAL"
)

var enums = map[Status]struct{}{
	Draft:                       {},
	Submitted:                   {},
	Approved:                    {},
	Rejected:                    {},
	Active:                      {},
	ActiveEditSubmitted:         {},
	ActiveRejectSubmitted:       {},
	ActiveInactivationSubmitted: {},
	Inactive:                    {},
	Unassigned:                  {},
	PendingApproval:             {},
	ActivePendingApproval:       {},
}

func IsValidStatus(s string) bool {
	_, exists := enums[Status(s)]
	return exists
}

func (s Status) ToString() string {
	return string(s)
}
