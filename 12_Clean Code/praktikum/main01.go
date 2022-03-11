package main

type user struct {
	id       int
	username int
	password int
}

type userservice struct {
	t []user //variabel t disini tidak jelas
}

func (u userservice) getallusers() []user {
	return u.t //untuk variabel awal t yang tidak jelas maka seluruh variabel t akan susah dimengerti
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}
