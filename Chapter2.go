package algor_notes
//算法分析
/*
观察
数学模型
分类
算法理论
记忆
每个角色 不同
程序员解决
客户 高效解决
理论 为啥work
避免性能错误。
举例 快速傅里叶变换 算法。
我的程序在大输入的时候，还能正常工作吗？
slow，run out of memory。
观察，假设，预测，验证，矫正。
原则：
实验可重复
假设可证伪。

 */
// 观察
/*
just test and get answer
估计算法的运行时间。
对性能做出预测，但不能帮助我们理解算法做了什么。

 */
//数学模型
/*
total running time
sum of cost*frequency for all operations
need to analyze program to determine set of operations
cost depends on machine,compiler
frequency depends on algorithm,input data.
 */
/*
基本操作的时间。
统计各项时间比较。
精确模型，留给理论专家
我们使用近似模型。

 */

//增长数量级的分类
//常数大于对数大于线性大于平方大于立方大于指数级。
//目标
//二分搜索例子
// 时间阶数 理论
 // 运行上界以及运行下界，平均运行时间。
 //实际数据输入和输入模型不匹配。
 // 需要了解如何处理它
 //为最坏和随机情况设计。
 //目标
 设计问题的困难性
 开发最优化算法
 //只考虑最坏情况。
 最优算法
 上界，下界，平均。
 目标：
 对问题的难度进行估计
 开发最优算法
 上界，o(n)
 下界 omega n
 如何判断最优算法。
 上界/下界等于常数，这是个最优算法。
// p12 内存
 bit 0/1
 Byte 8 bits
 MegaByte (MB) 1million or 2^20 bytes
 GigaByte (GB) 1billion or 2^30 bytes
8 byte 指针,64位寻址。
 char[] 2N+24 长度，容量。
 // objected overhead 16bytes
 // reference 8 bytes
 // padding
 //Each object uses a multiple of 8bytes
//经验分析
//数学分析
//科学方法。