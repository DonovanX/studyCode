package main

import "fmt"

// 将接口转换为其他接口

// 定义飞行动物接口
type Flyer interface {
	Fly()
}
// 定义行走动物接口
type Walker interface {
	Walk()
}

// 定义鸟类
type bird struct {
}

// 实现飞行动物接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// 为鸟添加Walk()方法, 实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// 定义猪
type pig struct {
}

// 为猪添加Walk()方法, 实现行走动物接口
func (p pig) Walk() {
	fmt.Println("pig: walk")
}

func main() {

	// 创建动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird":new(bird),
		"pig":new(pig),
	}

	for name, obj := range animals {

		// 判断对象是否为飞行动物
		f, isFlyer := obj.(Flyer)

		// 判断对象是否为行走动物
		w, isWalker := obj.(Walker)

		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n",
			name,isFlyer,isWalker)

		// 如果是飞行动物则调用飞行动物接口
		if isFlyer {
			f.Fly()
		}

		// 如果是行走动物则调用行走动物接口
		if isWalker {
			w.Walk()
		}
	}

	// 将接口转换为其他类型
	p1 := new(pig)

	var a Walker = p1
	p2 := a.(*pig)

	fmt.Printf("p1=%p p2=%p", p1, p2)
	
}


