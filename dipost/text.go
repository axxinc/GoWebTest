package main

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	delete() (err error)
	update() (err error)
}
