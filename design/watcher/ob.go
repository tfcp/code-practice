package watcher

import "fmt"

type BusinessInterface interface {
	Register(ob Observer)
	Remove(ob Observer)
	Notice(msg string)
}

type Business struct {
	obs []Observer
}

func (this *Business) Register(ob Observer) {
	this.obs = append(this.obs, ob)
}

func (this *Business) Remove(ob Observer) {
	for index, observer := range this.obs {
		if observer == ob {
			// 去除 ob
			this.obs = append(this.obs[:index], this.obs[index+1:]...)
		}
	}
}

func (this *Business) Notice(msg string) {
	for _, ob := range this.obs {
		ob.Handler(msg)
	}
}

type Observer interface {
	Handler(msg string)
}

type Log struct {
}

func (this *Log) Handler(msg string) {
	fmt.Println("this is Log msg: ", msg)
}

type User struct {
}

func (this *User) Handler(msg string) {
	fmt.Println("this is User msg: ", msg)
}

//type SubjectInterface interface {
//	Register(ob Observer)
//	Remove(ob Observer)
//	Notify(msg string)
//}
//
//type Subject struct {
//	obs []Observer
//}
//
//func (this *Subject) Register(ob Observer) {
//	this.obs = append(this.obs, ob)
//}
//
//func (this *Subject) Remove(ob Observer) {
//	for index, observer := range this.obs {
//		if observer == ob {
//			this.obs = append(this.obs[:index], this.obs[index+1:]...)
//		}
//	}
//}
//
//func (this *Subject) Notify(msg string) {
//	for _, ob := range this.obs {
//		ob.Update(msg)
//	}
//}
//
//type Observer interface {
//	Update(msg string)
//}
//
//type ObserverA struct {
//}
//
//func (this *ObserverA) Update(msg string) {
//	fmt.Println("ob A:", msg)
//}
//
//type ObserverB struct {
//}
//
//func (this *ObserverB) Update(msg string) {
//	fmt.Println("ob B:", msg)
//}
