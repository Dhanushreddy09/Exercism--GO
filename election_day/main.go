package election_day

import "fmt"

func Election_day() {
	var initialVotes int

	fmt.Println("Enter initial no of votes")
	fmt.Scan(&initialVotes)

	counter := NewVoteCounter(initialVotes)
	noOfVotes := VoteCount(counter)
	fmt.Println(noOfVotes)
}
