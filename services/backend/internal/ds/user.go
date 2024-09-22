package ds

import "sync"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserDatastore interface {
	Create(User)
	GetAll() []User
	GetOne(string) (User, bool)
}

type userDatastore struct {
	m map[string]User
	*sync.RWMutex
}

func (ds *userDatastore) Create(u User) {
	ds.RLock()
	defer ds.RUnlock()
	ds.m[u.ID] = u
}

func (ds *userDatastore) GetOne(id string) (User, bool) {
	ds.RLock()
	defer ds.RUnlock()
	u, ok := ds.m[id]
	return u, ok
}

func (ds *userDatastore) GetAll() []User {
	ds.RLock()
	defer ds.RUnlock()
	users := make([]User, 0, len(ds.m))
	for _, v := range ds.m {
		users = append(users, v)
	}
	return users
}

func CreateUserDatastore(initUsers []User) UserDatastore {
	ds := &userDatastore{
		m:       map[string]User{},
		RWMutex: &sync.RWMutex{},
	}
	for _, v := range initUsers {
		ds.m[v.ID] = v
	}
	return ds
}
