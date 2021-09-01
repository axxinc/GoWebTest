package main

type FakePost struct {
	Id      int
	Content string
	Author  string
}

func (fpost *FakePost) create() (err error) {
	return
}
func (fpost *FakePost) delete() (err error) {
	return
}
func (fpost *FakePost) update() (err error) {
	return
}
func (fpost *FakePost) fetch(id int) (err error) {
	fpost.Id = id
	return
}
