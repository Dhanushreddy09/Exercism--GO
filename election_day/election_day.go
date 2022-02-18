package election_day

import "fmt"

type ElectionResult struct {
	Name  string
	Votes int
}

func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}
func VoteCount(counter *int) int {
	if counter != nil {
		return *counter
	}
	return 0
}
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var candidateResult *ElectionResult
	candidateResult = &ElectionResult{Name: candidateName, Votes: votes}
	return candidateResult
}
func DisplayResult(result *ElectionResult) string {
	announcement := fmt.Sprintf("%v (%v)", result.Name, result.Votes)
	return announcement
}
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}
