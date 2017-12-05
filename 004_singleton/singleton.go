package singleton

import (
	"fmt"
	"sync"
)

//tesla的CEO都只有一个，无论怎么变都是一个
type CEO struct {

}

func (*CEO) name(){
	fmt.Println("Elon Musk")
}

var manger *CEO
var once sync.Once

func GetInstance() *CEO{

	// go很优雅的实现了单例只执行一次。。完美
	once.Do(func() {
		manger = new(CEO)
		fmt.Println("Doing something...")
	})

	//if manger ==nil {
	//	manger = new(CEO)
	//	fmt.Println("Doing something...")
	//}
	return manger
}