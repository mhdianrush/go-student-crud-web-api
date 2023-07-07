package entities

type Student struct {
	Id              uint   `json:"id"`
	FullName        string `validate:"required" label:"Full Name" json:"full_name"`
	StudentUniqueId string `validate:"required" label:"Student Unique Id" json:"student_unique_id"`
	Gender          string `validate:"required" label:"Gender" json:"gender"`
	BirthPlace      string `validate:"required" label:"Birth Place" json:"birth_place"`
	BirthDay        string `validate:"required" label:"Birth Day" json:"birth_day"`
	Address         string `validate:"required" label:"Address" json:"address"`
	PhoneNumber     string `validate:"required" label:"Phone Number" json:"phone_number"`
}
