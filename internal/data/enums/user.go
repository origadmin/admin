package enums

// Gender defines the gender of a user.
type Gender string

const (
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
	GenderUnknown Gender = "unknown"
)

// String returns the string representation of the gender.
func (g Gender) String() string {
	return string(g)
}
