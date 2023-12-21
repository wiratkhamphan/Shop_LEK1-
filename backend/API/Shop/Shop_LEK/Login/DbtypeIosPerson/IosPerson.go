package dbtypeiosperson

type IOSPerson struct {
	ID         string `json:"id"`
	Iso_p_code string `json:"iso_p_code"`
	P_code     string `json:"p_code"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	P_name     string `json:"p_name"`
}
