package memory

func (uc *userMemoryRepo) SetUser(username, token string) {
	uc.userData[username] = token
}

func (uc *userMemoryRepo) GetUser(username string) (data string, isFound bool) {
	data, isFound = uc.userData[username]
	return
}
