package algor_notes
// Goal sort any type of data
//ex1 sort random real numbers in ascending oeder.
//可比较的接口。
//compareto 泛化方法。
//total order 全序关系
// 反对称性
// 传递性
// 任意性
//两个 有用的函数
//less

//compare


//Q2 selection sort
//选择排序
//迭代i.找到剩下项中最小的。
//依次选出最小的。交换a[i]和min。冒泡排序
//保证 算法不变量
// Propositon
// selection sort uses N2/2 compares
// and N exchanges。

//Q3 insertion sort
//basic method
//每次 把 a[i] 和左边比它大的进行交换。
// 1/4 N2/2 compare 1/4 N2/2 exchanges.
//定义逆序对。
// 一个数组中逆序对数量小于一个值，即可认为它是部分有序。
// 插入排序，线性时间。

//Q4 shell sort
//希尔排序
//h-sorted array.
//一个大的增量，那么是一个小的子数组，
//小的子数组，近似有序。
//
//
//
//
 type shell struct{
 	h []int
 }
 func (s shell)Sort(){
 	n:=len(s.h)
 	h:=1
 	for h<n/3{
 		h=3*h+1
 		for h>1{
 			for i:=h;i<n;i++{
 				for j:=i;j>=h&&less(a[j],a[j-h]);j-=h{
 					exch(a,j,j-h)
				}
			}
			h=h/3
		}
	}
 }
 func less(i,j int){

 }
 func exch(a []int,i,j int){
 	exch
 }
 //the performance is O(n3/2).


 //q5 XiPai
 //generate a random real number for each
 // array entry
 //sort the array
 //数组排序
 // 均匀的随机数列。
 //生成随机数非常重要.

 //q6 计算几何领域。 排序的应用
 // 凸包，包住所有点的最小图多边形。
 // 必须要意识到这类问题。边和顶点。
 // 隔离恒 扫描法，
 // 建立一个坐标系。按照急性角，对点进行排序。
 // 抛弃掉不合适的点。
 // 判断三个点之间是顺序还是逆序的。
 // ccw Given three points
 //a ,b,c is a->b->c a counterclockwise tuen.
 // ccw 可以通过计算点积来确定。

