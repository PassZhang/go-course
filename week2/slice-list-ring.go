// slice list ring 
/* 
1. slice 在初始化的时候会分配一块连续的内存空间

2. list 和 ring 是非连续的内存空间（上一个块指向下一个块的位置），可以延迟初始化的（初始化操作延后，仅在实际需要的时候进行，可以分散初始化操作带来的计算量和存储空间消耗）

3. 如果需要频繁的增加或删除元素，可以用链表来操作

4. list可以作为队列和stack的基础数据结构；ring可以用来报错固定数量的元素，例如保存最近100条日志，用户最近10次操作；heap可以用来排序，游戏编程中是一种高效的定时器实现方案
list的len方法时间复杂度O(1),Ring的Len方法时间复杂度为O(N)

*/ 

package main


func (l *List) Len() int { return l.len }

func (r *Ring) Len() int {
   n := 0
   if r != nil {
      n = 1
      for p := r.Next(); p != r; p = p.next {
         n++
      }
   }
   return n
}