## Go语言结构

==（ps：需要具有一定的c语言基础）==

Go的基础组成有以下几个部分：

- 包声明
- 引入包
- 函数
- 变量
- 语句&表达式（程序主体）
- 注释

下面是一个简单的go程序：

```go
package main

import("fmt")

func main(){    /*在Go中，左括号‘{’不能单独占一行，否则会报错*/
    fmt.Println("Hello,World!")
}
```

第一行代码`package main`定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，每个 Go 应用程序都包含一个名为 main 的包。

第三行的代码表示需要用到fmt包，它里面实现了标准输入输出的函数，后面的Println中的ln是指line，所以打印后会自动换行；如果需要格式化输出，使用的是Printf，占位符类比c语言即可。

当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。

关于包：

- ==文件名==与包名没有直接关系，不一定要将文件名与包名定成同一个
- ==文件夹名==与包名没有直接关系，并非需要一致。
- 同一个文件夹下的文件只能有一个包名，否则编译报错



## Go语言基础语法

#### 行分隔符

在Go中，一行代表一个语句结束，不需要像c语言以；结尾，但如果你想要将多个语句写在同一行，必须使用；人为区分。

#### 注释

```go
//单行注释
/*
多行注释
*/
```

#### 标识符

标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

#### 字符串连接

Go 语言的字符串连接可以通过 **+** 实现：

```go
package main
import "fmt"
func main() {
    fmt.Println("Google" + "Runoob")
}
```

输出：

```tex
GoogleRunoob
```

#### Go语言的空格

Go中的空格，通常用于分隔标识符、关键字、运算符和表达式，以提高代码的可读性。但Go 语言中变量的声明必须使用空格隔开，否则变量类型和标识则组成了一个新的标识。

无空格：

```
fruit=apples+oranges;
```

有空格：

```
fruit = apples + oranges; 
```

对比下来，有空格更舒服一点。

#### 格式化字符串

Go 语言中使用 **fmt.Sprintf** 或 **fmt.Printf** 格式化字符串并赋值给新串：

- **Sprintf** 根据格式化参数生成格式化的字符串==并返回==该字符串。
- **Printf** 根据格式化参数生成格式化的字符串并==写入标准输出==。



## Go语言数据类型

在 Go 编程语言中，数据类型用于声明函数和变量。

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。

除了最基础的几种类型外，Go为了方便编程，派生了一些新的数据类型：

- Channel类型
- 切片类型
- 接口类型
- Map类型

后续会对这几种作介绍。

## Go语言变量

声明变量的一般形式是使用 var 关键字，基本格式如下：

`var identifier type`

变量的类型是写在后面的，与c语言刚好相反，刚接触会有点别扭。

#### 变量声明

##### 第一种，指定了变量类型，如果未初始化，默认为零值：

- 数值类型为0
- 布尔类型为false
- 字符串为“”（空串）
- 以下几种类型为nil（后续介绍）：

```go
var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口
```

##### 第二种，根据值自行判定变量类型（类Python）

##### 第三种，**如果变量已经使用 var 声明过了，再使用 `:=` 声明变量，就产生编译错误，格式：**

`v_name := value`

例如：

```go
var intVal int 
intVal := 1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
```

`intVal := 1`相当于：

```go
var intVal int
intVal = 1
```



#### 多变量声明

直接看一个实例即可：

```go
package main
import "fmt"

var x, y int
var(	// 这种因式分解关键字的写法一般用于声明全局变量
	a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现（类Python）
//g, h := 123, "hello"

func main(){
    g, h = 123, "hello"
    fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
```

执行结果为：

`0 0 0 false 1 2 123 hello 123 hello`



#### 使用 `:=` 赋值操作符

这是使用变量的首选形式，但是它==只能被用在函数体内==，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。

##### 注：

如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误**a declared but not used**。但是全局变量是允许声明但不使用的。

空白标识符 _ 也被用于抛弃值，如值 5 在 `_, b = 5, 7` 中被抛弃。\_实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。



## Go语言常量

常量需使用`const`进行修饰，定义格式如下：

`const identifier type = value`

```go
const WIDTH int = 7
```

常量还可用作枚举：

```
const(
	Unknown = 0
	Female = 1
	Male = 2
)
```

#### iota

iota是一个特殊常量，或认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次

iota也可被用作枚举：

```go
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "go"   //独立值，iota继续递增1
            e          //"ha"   iota += 1
            f = 77    //iota +=1
            g          //77  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a, b, c, d, e, f, g, h, i)
}
```

输出结果：

`0 1 2 go go 77 77 7 8`



## Go语言条件语句

以后需要补充

#### switch语句

switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 **fallthrough** ，使用fallthrough将不会判断下一条case的表达式是否为true。具体语法与c语言类似。

#### type switch

根据某个interface变量中实际存储的变量类型来做出不同反应（有点多态的意思）

#### select语句

类似于switch语句，但select语句只能用于通道操作，要么发送要么接收。select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。



## Go语言循环语句

#### for循环

和C语言的for一样：

`for init ; condition ; post {}` 或 `for {}`

同C的while：

`for condition {}`

for的range格式可以对slice、map、数组、字符串等进行迭代循环：

```go
for key, value := range oldMap{
    newMap[key] = value
}
```

这里贴一个乘法表的代码：

```go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v x %v = %v ", i, j, i*j)
		}
		fmt.Print("\n")
	}
}
```

#### 循环控制语句

| break 语句    | 经常用于==中断==当前 for 循环或跳出 switch 语句      |
| ------------- | ---------------------------------------------------- |
| continue 语句 | ==跳过==当前循环的剩余语句，然后继续进行下一轮循环。 |
| goto 语句     | 将控制转移到被标记的语句。                           |



## Go语言函数

Go中函数和c语言很类似，只是把函数返回值类型放在了函数名称后面，另外可以返回多个值，使用如下：

```go
func function_name([parameter_list]) (type1,type2){
	//code
}
```

#### 函数用法

- 闭包（匿名函数）：匿名函数是一个“内联”语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量而不必声明。匿名函数是一种==没有函数名==的函数，通常用于==在函数内部定义函数==，或者作为函数参数进行传递。

```go
package main

import "fmt"

func main() {
    // 定义一个匿名函数并将其赋值给变量add
    add := func(a, b int) int {
        return a + b
    }

    // 调用匿名函数
    result := add(3, 5)
    fmt.Println("3 + 5 =", result)

    // 在函数内部使用匿名函数
    multiply := func(x, y int) int {
        return x * y
    }

    product := multiply(4, 6)
    fmt.Println("4 * 6 =", product)

    // 将匿名函数作为参数传递给其他函数
    calculate := func(operation func(int, int) int, x, y int) int {
        return operation(x, y)
    }

    sum := calculate(add, 2, 8)
    fmt.Println("2 + 8 =", sum)

    // 也可以直接在函数调用中定义匿名函数
    difference := calculate(func(a, b int) int {
        return a - b
    }, 10, 4)
    fmt.Println("10 - 4 =", difference)
}
```

以上代码中，关于将匿名函数作为参数传递可能会有点难理解。在这个代码中，它的作用就是将数学计算中的加减乘除全部抽象为一种计算（calculate），在需要计算时，只需要传递操作名（如add，multiply）和操作数即可得到想要的结果。



## Go语言函数方法

类比C++中的类，在类里面的函数我们也叫做方法。Go语言中只需在定义函数时，在函数名前面加上该方法的接受者（可以理解为所属者）。接受者可以是命名类型或结构体类型的一个值或指针，具体格式如下：

```go
func (variable_name variable_data_type) function_name() [return_type]{
   /* 函数体*/
}
```

例子：求面积

```go
package main

import "fmt")

/* 定义结构体 */
type Circle struct {
  radius float64
}

func main() {
  var c1 Circle
  c1.radius = 10.00
  fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}
```



## Go语言变量作用域

作用域是指已声明标识符的常量、类型、变量、函数的作用范围。一旦超出该作用范围，编译器则无法识别到它们。与C语言类似：

- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

#### 局部变量

它们的作用域只在函数体内参数和返回值也是局部变量。

#### 全局变量

在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。

Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。

#### 形式参数

形式参数会作为函数的局部变量来使用。



## Go语言数组

数组的具体含义参考C语言，这里主要记录一下在Go中数组的定义等。

#### 声明数组

`var arrayName [size]dataType`

#### 初始化数组

未初始化默认初始值为0

`var nums [5]int{1, 2, 3, 4, 5}`

需要注意的是，在Go中，数组的大小是类型的一部分，因此不同大小的数组是不兼容的，或者说是不同的类型。

如果数组长度不确定，可以使用`...`来代替数组的长度，编译器会自行推断数组长度。

有意思的是，如果设置了数组长度，可以通过指定下标来初始化元素：

```go
array1 := [5]int{1:2, 2:5}
```

这个特性，让数组和字典（python中的叫法）有了共同点：数组可以理解为key是连续增长的数字的字典。



## Go语言指针

指针总的概念和用法与C语言差不多，唯一需要注意的就是指针声明时的区别：

```go
var var_name *var_type
```

指针声明中的那个*，是在变量类型的前面。

#### 

## Go语言结构体

结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。在Go中定义结构体与C语言仍然有着一些区别：

```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```

一旦定义了结构体之后，它就能用来当做变量类型来使用（由此可见结构体和类在本质上是差不多的）可以理解为自己定义了某种变量类型的关键词。

在Go中，不管是结构体指针的变量还是结构体变量的实例，在访问成员时都是使用.来进行访问，与C语言的->有所不同。



## Go语言切片(Slice)

数组长度不可变，一些特定场景不太适用。而切片的长度不是固定的，可以追加元素，且追加时可能使切片的容量增大。与python中的类似。

#### 定义切片

```go
var identifier []type
```

实际上就是声明一个未指定大小的数组

或使用make()函数来创建切片：

```go
var identifier []type = make([]type, len)
//或简写成
slice := make([]type, len)
```

len是数组的长度，也是切片的原始长度。

```go
s := array[startIndex:endIndex]
```

表示从array数组中下标startIndex到endIndex-1的元素创建为一个新的切片s。

#### len()和cap()函数

len()得出的是切片中有几个元素，而cap()得到的是切片的容量（capacity)

#### append() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

append()可以向切片中追加元素，copy()函数作用就是复制。

```go
package main

import "fmt"

func main(){
	var nums []int
    nums = append(nums, 0, 1, 2, 3)//添加元素
    nums1 := make([]int, len(nums))
    copy(nums1, nums)//复制
}
```



## Go语言范围(Range)

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中==返回 key-value 对==。

```go
package main
import "fmt"
func main() {
    //这是我们使用 range 去求一个 slice 的和。使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    //在数组上使用 range 将传入索引和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    //range 也可以用在 map 的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    //range也可以用来枚举 Unicode 字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```



## Go语言集合(Map)

Map 是一种无序的键值对的集合。它可以通过key来快速检索数据，key类似于索引，指向数据的值。

在获取 Map 的值时，如果键不存在，返回该类型的零值，例如 int 类型的零值是 0，string 类型的零值是 ""。

Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量。

#### 定义Map

使用函数make或使用map关键字来定义

```go
map_var := make(map[kType]vType, initialCapacity)
```

当 Map 中的键值对数量达到容量时，Map 会==自动==扩容。如果不指定 initialCapacity，Go 语言会根据实际情况选择一个合适的值。

#### delete()函数

delete()用于删除集合的元素，参数为map和其对应的key。用法：

```go
delete(map, "key")
```



## Go语言类型转换

#### 字符串类型转换

需要import"strconv"，使用其中的函数。

- strconv.Itoa(int):将int整型转换为字符串
- strconv.Atoi(string):将字符串转换为整型

以上函数会返回两个值，第一个是转换后的值，第二个是可能发生的错误。

#### 接口类型转换



## Go语言接口

这里可以结合C++的纯虚函数和抽象类来理解。它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。换句话说，某个类如果实现了该接口的所有方法，也就是实现了这个接口。

Go 语言中的接口是隐式实现的，也就是说，如果一个类型实现了一个接口定义的所有方法，那么它就自动地实现了该接口。因此，我们可以通过将接口作为参数来实现对不同类型的调用，从而实现多态。

```go
/* 定义接口 */
type interface_name interface{
    method1_name [return_type]
    ...
    methodn_name [return_type]
}

type struct_name struct{
    //variables
}

func (s_name_var s_name) method1_name() [return_type]{
    //code
}
...
func (s_name_var s_name) method_name() [return_type]{
    //code
}
```

接口类型变量可以存储任何实现了该接口的类型的值。简单理解就是接口从属于实现了它的类型，实现了接口的类型，一定包含了接口。



## Go并发

Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。

goroutine 语法格式：

```go
go 函数名( 参数列表 )
```

#### 通道(channel)

通道是用来传递数据的一个数据结构。通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 `<-` 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

```go
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v
```

声明通道使用关键词chan：

```go
ch := make(chan int)
```

默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

#### 通道缓冲区

通过make函数的第二个参数设置缓冲区大小：

```go
ch := make(chan int, 100)
```

带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

**注意**：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。

#### Go遍历通道与关闭通道

Go 通过 range 关键字来实现遍历读取到的数据。格式如下：

```go
v, ok := <-ch
```

如果通道接收不到数据后 ok 就为 false，这时通道就可以使用 **close()** 函数来关闭。
