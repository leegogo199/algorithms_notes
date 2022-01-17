package algor_notes
/*
介绍
解决问题问题的算法
以及存储信息的数据结构

datatypes  stack queue bag union-find,priority,queue
sorting    quicksort mergesort headsort radix sorts
searching  bst red-black bst ,hash table
graphs     bfs dfs prim kruskal dijkstra
strings    kmp regular expressions tst huffman lzw
advanced   b tree suffix array maxflow
算法在所有领域都有使用。
历史悠久，机缘深厚。
书籍配套网站。

 先修要求
编程，java，高中数学。
编程环境
练习
写一个java程序。

*/

/*
sec 2
如何开发一个有用的算法
对问题建模
找到一个方法解决它
时间空间复杂度是否满足要求
如果不满足
找到一个方法解决它
迭代到合适为止。

Given a set of n objects
union command  connect two objects
Find/connected query is there a path connecting the two objects
 连接性的特性
自反
对称
传递
 联通的对象组成的集合叫连同分量。
两个方法
1 查找两个元素是否在一个集合
2 将两个元素放到一个集合。

目标 针对并查集设计一个有效的数据结构
元素n，操作m都很大。
查找和合并操作量很大。
*/
type UF struct {
	n int
}
func union( p,q int){

}
func connected(p,q)bool{

}
func find(int)int{

}
func count()int{

}
// quick find
/*
数据结构
容量为n的整数数组
当 p和q的id 一致时，我们认为两者是相连的。
方法 一 查找
较为简单，判断两个的id是否相同。
方法 二 合并
较为困难，需要将和项1id相同的项id均改为项2的id
 */
type quickfindUF struct{
	var id []int
	var n int
}
func (q quickfindUF)qUF_init()[]int{
	q.id=make([]int,q.n)
	for i,_:=range q.id{
		q.id[i]=i
	}
	return q.id
}
func (Q quickfindUF)qUF_Connected(p,q int)bool{
	return Q.id[p]==Q.id[q]
}
func (Q quickfindUF)qUF_union(p,q int){
	pid=Q.id[p]
	qid=Q.id[q]
	for i:=0;i<Q.n;i++{
		if Q.id[i]==pid{
			Q.id[i]=qid
		}
	}
}
// 评估
/*
  初始化 n，合并 n,查找 n
查找费时间，它的缺点。
N的平方级时间。
非常慢
*/
/*
quick union 改进
避免计算
数据结构
n的整数数组
id[i]是i的父辈
i的根是id【id【id【id【i】
// 查找
检查p和q的根是否相同
//合并
将p的root改成q的root。
*/

type Quuf struct {
	id []int
	n int
}
func (q Quuf)Qu_init(){
	q.id=make([]int,q.n)
	for i:=0;i<n;i++{
		q.id[i]=i
	}
}
func (q Quuf)root( i int) int{
	for i!=id[i]{
		i=id[i]
	}
	return i
}
func (Q Quuf)connected(p,q int)bool{
	return Q.root(p)==Q.root(q)
}
func (Q Quuf)union(p,q int){
	i:=Q.root(p)
	j:=Q.root(q)
	Q.id[i]=j
}
// 评价
/*
trees can get tall
find too expensice


*/
//quick union improvemnet.
/*
避免生成一个很高的树，尽量生成平衡树。
增加一个数据结构，存储树的数量，
大树做根，小树去连接大树的根。
 */
/*
data structure
same as quick-union
maintain extra array sz[i]

*/
type imuf struct{
	Quuf
	sz []int
}
func (q imuf)Qu_init(){
	q.id=make([]int,q.n)
	for i:=0;i<n;i++ {
		q.id[i] = i

		q.sz[i] = 1
	}
}
func (Q imuf)quick_union(p,q int){
	i:=Q.root(p)
	j:=Q.root(q)
	if (i==j) {
		return
	}
	if (Q.sz[i]<Q.sz[j]){
		Q.id[i]=j
		Q.sz[j]+=Q.sz[i]
	}else{
		Q.id[j]=i
		Q.sz[i]+=Q.sz[j]
	}
}
//评估
/*
find 正比于p和q的时间
union 常数时间
节点n的深度最多为lg2n。
*/
// 改进 方案，每次 合并后，d
//每次都指向 更节点。
// in practivce we can just do it.

//1.6 application
标记图像中的最小连同集。
