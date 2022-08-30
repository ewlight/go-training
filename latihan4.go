package main
import ("fmt"
		"sync"
)


type User struct {
	nama string
}

type UserService struct {
	db []*User
}

type UserInterface interface {
	Register(u *User) string
	GetUser() []*User

}	

func NewUserService(db []*User) UserInterface {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.nama + " berhasil didaftarkan"
}

func (u *UserService) GetUser() []*User {
	return u.db
}

func main() {
	var db []*User
	userService := NewUserService(db)
	names := []string{"hilmi", "bayu", "muthus", "bintang", "thony"}
	var wg1 sync.WaitGroup
	wg1.Add(len(names))
	for _, n := range names {
		go func(name string) {
			res := userService.Register(&User{nama: name})
			fmt.Println(res)
			wg1.Done()
		}(n)
	}
	wg1.Wait()

	resGet := userService.GetUser()
	fmt.Println("-----------Hasil get user-------------")
	var wg sync.WaitGroup
	wg.Add(len(resGet))
	for _, v := range resGet {
		go cetakNama(&wg, v.nama)
	}
	wg.Wait()

}

func cetakNama(wg *sync.WaitGroup, nama string) {
	fmt.Println(nama)
	wg.Done()
}

