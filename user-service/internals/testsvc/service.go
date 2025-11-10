package testsvc

type User struct {
	ID   int
	Name string
	Age  int
}

type Repo interface {
	Add(user User) User
	All() []User
	Update(id int, name string, age int) bool
}

type inMemoryRepo struct {
	users  []User
	nextID int
}

func NewInMemoryRepo() *inMemoryRepo {
	return &inMemoryRepo{
		users:  []User{},
		nextID: 1,
	}
}

type UserService struct {
	repo Repo //
}

func NewUserService(repo Repo) *UserService {
	return &UserService{repo: repo}
}

func (r *inMemoryRepo) Add(user User) User {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user

}

func (r *inMemoryRepo) All() []User {
	out := make([]User, len(r.users))
	copy(out, r.users)
	return out
}

func (s *UserService) CreateUser(name string, age int) (User, bool) {
	if name == "" {
		return User{}, false
	}
	if age < 0 {
		return User{}, false
	}
	user := s.repo.Add(User{Name: name, Age: age})
	return user, true
}

func (r *inMemoryRepo) Update(id int, name string, age int) bool {
	for i, u := range r.users {
		if u.ID == id {
			r.users[i].Name = name
			r.users[i].Age = age
			return true
		}
	}
	return false
}

func (s *UserService) ListUsers() []User {
	return s.repo.All()
}

func (s *UserService) UpdateUser(id int, name string, age int) bool {
	return s.repo.Update(id, name, age)
}
