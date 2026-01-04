package main

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MembershipCharLimit {
		return message, true
	} else {
		return "", false
	}
}

type User struct{
	Name string
	Membership
}

type Membership struct {
	Type string
	MembershipCharLimit int
}

func newUser(name string, membershipType string) User{
	membership := Membership{Type: membershipType}
	if membershipType == "premium"{
		membership.MembershipCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MembershipCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}
