package census

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{Name: name, Age: age, Address: address}
}
func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && len(r.Address) != 0 && r.Address["street"] != ""
}
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}
func Count(residents []*Resident) int {
	count := 0
	for _, val := range residents {
		if val.HasRequiredInfo() {
			count++
		}
	}
	return count
}
