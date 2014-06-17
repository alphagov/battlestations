package github

import (
	"fmt"
)

type MembershipError struct {
	organisation string
}

func (err *MembershipError) Error() string {
	return fmt.Sprintf("You are not a member of %s", err.organisation)
}
