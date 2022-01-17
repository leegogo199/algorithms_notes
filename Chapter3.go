package algor_notes
// 1.3 bags,queues,stacks.
// Fundamental data types
/*
值 对象集合
操作，插入，删除，迭代，空检验
插入比较简单
删除哪一个元素值得考虑。
经典结构
stack.
queue.
区分实现和接口
优点
客户端不知道实现的细节
实现不知道客户的需求
设计 创造模块，重用 库
性能 使用最优的有效算法。
 */
// stack api
// warmup api
// 存放字符串类型的栈
// 公共类 字符串类型的栈
/*
以下几个方法
初始化，push，pop，判空，大小。

 */
type Node struct {
	Val string
	Prev *Node
}
type Stack struct{
	Head *Node
	Tail *Node
}
func (s *Stack) isEmpty()bool{
	f:=(s.Head==nil)
	return f
}
func (s *Stack)Push(item string){
	temp:=&Node{item,nil}
	temp.Prev=s.Tail
	s.Tail=temp
}
func (s *Stack)Pop() string{
	item:=s.Tail.Val
	s.Tail=s.Tail.Prev
	return item
}
//容量调整
// 满的时候，容量翻倍。
// 1/4的时候，容量减半。
// 可调数组和链表实现各有各的优缺点。
// 针对具体情况，抉择使用。

//Q3 queue.
/*
队列需要的函数 ，字符队列
入队列，出队列，队列判空，队列尺寸。


*/
//q4 generics
//go 最近才实现泛型，所以，这里不考虑了。
//可以用接口来实现。
// 包裹类型
type Stack2 struct{
	value []interface{}
}


//q5 迭代器 优化
/*
可以返回迭代器的方法的一种类。
Iterable
iterator
*/
type Iterator interface {
	HasNext() bool
	next() interface{}
}
func(s *Stack)HasNext()bool{
	return s.IsEmpty
}
//q 6应用
//使用前了解api。
// 编译器如何实现一个功能
/*
功能调用，压栈 并且返回地址
返回， 谈栈，返回当地地址
迭代函数，调用自己
笔记 可以永远使用一个栈来替代递归调用。
两个栈，解决 算术表达式求值的方法。
*/
type stack3 interface {
	value []interface{}
}
func (s *stack3)Push( val interface{}){
	s=append(s,val)
}
func(s *stack3)Pop()interface{}{
	item:=s.value[(len(s.value)-1)]
	s.value=s.value[0:len(s.value)-1]
	return item
}
func test3(s string)int{
	var s1 stack3
	var s2 stack3
	var ans int
	for i,_:=range s{
		if s[i]=="+"||s[i]=="-"{
			s1.Push(s[i])
		}else if s[i]=='('{
		}else if s[i]==')'{
			if s1.Pop()=='+'{
				ans=ans+s2.Pop()+s2.Pop()
			}else if s1.Pop()=='*'{
				ans=ans+s2.Pop()*s2.Pop()
			}
		}else if int(s[i])>-1||int(s[i])<10{
			s2.Push(s[i])
		}
	}
	return ans
}
