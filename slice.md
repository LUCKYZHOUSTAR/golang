####slice是可变长的
定义完一个slice变量之后，不需要为它的容量而担心，你随时可以往slice里面加数据

比如：

v:=[]string{}

v=append(v, "hello")

这里附带说一下，slice和array的写法很容易混

v:=[2]string{"str1", "str2"} //这个是array

m:=[]string{"str1","str2"} //这个是slice

写法上默念：array有长度slice没长度，array有长度slice没长度…
#### slice是一个指针而不是值。
指针比值可就小多了，因此，我们将slice作为函数参数传递比将array作为函数参数传递会更有性能。

slice是一个指针，它指向的是一个array机构，它有两个基本函数len和cap。

看下面的图示：

clip_image003

slice是一个带有point（指向数组的指针），Len（数组中实际有值的个数），Cap（数组的容量）

比如上面的这个slice，它指向的数组是[3]int,其中的前两个有值，第三个为空

那么

len(slic) = 2

cap(slic) = 3

 append函数就理解为往slice中加入一个值，如果未达到容量（len<cap）那么就直接往数组中加值，如果达到容量（len = cap）那么就增加一个新的元素空间，将值放在里面

个人理解：

如果真的想用数组，应该用slice
