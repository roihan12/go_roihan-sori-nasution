package main

type user struct {
	id       int
	username int
	password int
}

/*
Penamaan variabel jika lebih dari dua kata dapat menggunakan gaya penulisan camelCase agar mudah di cari dan dieja contoh pada userservice dapat diubah menjadi userService
*/
type userservice struct {
	// Penamaan tidak mudah dipahami dan tidak mendeskripsikan konteks
	t []user
}

/*
Penamaan function atau method jika lebih dari dua kata dapat menggunakan gaya penulisan camelCase agar mudah di cari dan dieja contoh pada method "getallusers" dapat diubah menjadi "getAllUsers" dan "getuserbyid" dapat diubah menjadi "getUserById"
*/
func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
			// variable "r" tidak mudah dipahami dan tidak mendeskripsikan konteks
		}
	}

	return user{}
}

//  Tidak ada komentar pada kode atau function yang menjelaskan function tersebut