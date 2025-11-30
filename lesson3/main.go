package main

import (
	"fmt"
)

// 火影忍者管理系统
type Naruto struct {
	name     string
	addr     string
	ninjutsu string
}
type Narutosmanger struct {
	nanjas map[string]Naruto
}

func newNarutosmanger() *Narutosmanger {
	return &Narutosmanger{
		nanjas: make(map[string]Naruto, 100), //初始化
	}
}
func (m *Narutosmanger) Add(n Naruto) {
	if _, ok := m.nanjas[n.name]; ok {
		fmt.Printf("%v已存在，请勿重复添加", n.name)
		return
	}
	m.nanjas[n.name] = n
	fmt.Printf("添加用户：%v,来自%v,忍术有%v\n", n.name, n.addr, n.ninjutsu)
}
func (m *Narutosmanger) update(n *Naruto, newaddr, newninjutsu string) {
	v, ok := m.nanjas[n.name]
	if !ok {
		fmt.Printf("%v不存在，无法修改", n.name)
		return
	}
	v.addr = newaddr
	v.ninjutsu = newninjutsu
	fmt.Printf("更新用户:%v,更改数据为：%v,%v\n", n.name, v.addr, v.ninjutsu)
}
func (m *Narutosmanger) delete(n Naruto) {
	_, ok := m.nanjas[n.name]
	if !ok {
		fmt.Printf("%v不存在，无法删除", n.name)
		return
	}
	delete(m.nanjas, n.name) //运用go中自带的delete删除值键对(map，key)
	fmt.Printf("已删除用户：%v\n", n.name)
}
func (m *Narutosmanger) list() {
	if len(m.nanjas) == 0 {
		fmt.Println("当前无忍者信息")
		return
	}
	fmt.Println("==忍者列表==")
	for _, v := range m.nanjas {
		fmt.Printf("忍者：%v 所属：%v 忍术：%v\n", v.name, v.addr, v.ninjutsu)
	}
}

func main() {
	manger := newNarutosmanger()
	naruto1 := Naruto{
		name:     "漩涡鸣人",
		addr:     "木叶",
		ninjutsu: "螺旋丸",
	}
	naruto2 := Naruto{
		name:     "宇智波佐助",
		addr:     "宇智波一族",
		ninjutsu: "千鸟",
	}
	manger.Add(naruto1)
	manger.Add(naruto2)
	manger.update(&naruto1, "木叶村", "影分身")
	manger.delete(naruto1)
	manger.list()
}

/*
package main

import (
"fmt"
"strings"
)

type Alice struct {
	Name  string
	Age   int
	Lover string
}
type Bob struct {
	Name  string
	Age   int
	Intro string
}
type Carol struct {
	Name  string
	Age   int
	Lover string
}
type AliceConf struct {
	Name  string
	Age   int
	Lover string
}
type BobConf struct {
	Name string
	Age  int
	Tags []string
}
type CarolConf struct {
	Name string
}
type DefaultPtr[T usedType, C usedTypeConf] interface { //泛型接口
	*T
	Init(C)
}
type usedType interface {
	Alice | Bob | Carol
}
type usedTypeConf interface {
	*AliceConf | *BobConf | *CarolConf
}

func (a *Alice) Init(cfg *AliceConf) {
	a.Name = cfg.Name
	a.Age = cfg.Age
	a.Lover = cfg.Lover
}
func (b *Bob) Init(cfg *BobConf) {
	b.Name = cfg.Name
	b.Age = cfg.Age
	b.Intro = fmt.Sprintf("I like to%v", strings.Join(cfg.Tags, ",")) //Sprintf返回拼接后的字符串，printf直接打印内容到终端，Join（切片，分隔符）能将切片中的字符串用分隔符连接成新字符串
}
func Default[T usedType, C usedTypeConf, ptr DefaultPtr[T, C]](cfg C) T {
	var t T
	ptr.Init(&t, cfg) //(&t).Init(cfg)
	return t
}
func main() {
	fmt.Println(Default[Alice, *AliceConf](&AliceConf{
		Name:  "Alice",
		Age:   20,
		Lover: "Bob",
	}))
	fmt.Println(Default[Bob, *BobConf](&BobConf{
		Name: "Bob",
		Age:  18,

		Tags: []string{"Sport", "Coding"},
	}))
}
*/
