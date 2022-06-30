package entities

type Students struct {
	StudentId        string `json:"studentId"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Mobile           string `json:"mobile"`
	Email            string `json:"email"`
	EmergencyName    string `json:"emergencyName"`
	EmergencyContact string `json:"emergencyContact"`
	ReferredBy       string `json:"referredBy"`

	Gender        string `json:"gender"`
	DateOfBirth   string `json:"dateOfBirth"`
	Occupation    string `json:"occupation"`
	Education     string `json:"education"`
	Relationship  string `json:"relationship"`
	DateOfJoining string `json:"dateOfJoining"`
	File          string `json:"file"`

	Height                 string `json:"height"`
	Weight                 string `json:"weight"`
	CurrentLevelOfActivity string `json:"currentLevelOfActivity"`
	CurrentRoutine         string `json:"currentRoutine"`
	MostStress             string `json:"mostStress"`
	PranayamaPractice      string `json:"pranayamaPractice"`
	AnySurgery             string `json:"anySurgery"`
	AnyMedical             string `json:"anyMedical"`
	Smoked                 string `json:"smoked"`

	ReasonForYoga             string `json:"reasonForYoga"`
	Musculoskeletal           string `json:"musculoskeletal"`
	Respiratory               string `json:"respiratory"`
	Cardiovascular            string `json:"cardiovascular"`
	Circulatory               string `json:"circulatory"`
	Neurological              string `json:"neurological"`
	Gastrointestinal          string `json:"gastrointestinal"`
	Endocrinological          string `json:"endocrinological"`
	GynecologicalOrUrological string `json:"gynecologicalOrUrological"`
	OtherMedicalConditions    string `json:"otherMedicalConditions"`
	OtherInformations         string `json:"otherInformations"`

	// UserId       string `json:"userId"`
	// CreatedBy    string `json:"createdBy"`
	// CreatedDate  string `json:"createdDate"`
	// ModifiedBy   string `json:"modifiedBy"`
	// ModifiedDate string `json:"modifiedDate"`
	// ActiveStatus bool   `json:"activeStatus"`
}
