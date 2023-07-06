package entities

type Student struct {
	Id              uint   `json:"id"`
	FullName        string `json:"full_name"`
	StudentUniqueId string `json:"student_unique_id"`
	Gender          int8   `json:"gender"`
	BirthPlace      string `json:"birth_place"`
	BirthDay        string `json:"birth_day"`
	Address         string `json:"address"`
	PhoneNumber     string `json:"phone_number"`
}
