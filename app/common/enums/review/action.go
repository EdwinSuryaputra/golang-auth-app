package reviewenum

type Action string

var (
	Approve Action = "APPROVE"
	Reject  Action = "REJECT"
)

var actionEnums = map[Action]struct{}{
	Approve: {},
	Reject:  {},
}

func IsValidAction(s string) bool {
	_, exists := actionEnums[Action(s)]
	return exists
}
