package user

type User struct {
	UserId   Int    `json:user_id`
	Username string `json:username`
	Password struct {
		Password string            `password`
		QA       map[string]string `qa`
	} `json:password`
	Role string `json:role`
	Data struct {
		Email string `json:email`
		FirstName string `json:first_name`
		MiddleName string `json:middle_name`
		LastName string `json:last_name`
		Gender string `json:gender`
		UserNote string `json:user_note`
		Address1 string `json:address_1`
		Address2 string `json:address_2`
		City string `json:city`
		Region string `json:region`
		Zip string `json:zip`
		Country string `json:country`
		Active int `json:active`
		CreatedDate string `json:created_date`
		UpdatedDate string `json:updated_date`
	} `json:data`
	LoginAttempts map[string]interface{} `json:login_attempts`
}

func (user *User) Create() (string, error) {

	return user, nil
}

func (user *User) Scan(keyword string) (*User, error) {

	return user, nil
}

func Read(user_id string, username string) (*User, error) {

	return User{}, nil
}

func (user *User) Update() (*User, error) {

	return user, nil
}

func (user *User) Delete() (bool, error) {

	return true, nil
}
