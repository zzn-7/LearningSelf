# Git简介

Git是一款分布式的代码版本控制工具，由Linus自己开发用来维护Linux版本的工具（这是真大佬）。

自己在最开始使用时，很难理解其版本控制的逻辑，觉得非常难用，有时候甚至出现版本控制给控制没了的情况。但是很多人在掌握Git后，都觉得Git十分优雅好用。

在一些多人合作进行的大型项目中，合理运用好Git版本控制，你可以很轻松的知道：

- 某个模块是谁写的？

- 某个文件的这一行是什么时候编辑的？谁编辑的？为什么要编辑？

## Git的数据模型

### 快照（Snapshots）

Git将某个顶级目录中的文件和文件夹的集合的历史记录建模成为一系列的快照。在Git中，一个文件被称为“blob”，它只是一堆字节。目录被称为“tree”，它将名称映射到Blob或树（因此目录可以包含其他目录）。

### 历史记录的建模与快照之间的联系

在一个简单的模型中，历史记录是线性的，按照时间顺序排列。由于各种原因，Git则没有使用这种简单的线性模型。

在Git中，历史记录是快照的有向无环图（DAG）。Git中每个快照都指向之前的一组快照，因为会有合并两个并行的开发分支的情况，所以每个快照可以指向不同的前驱快照（合并），也会被不同的快照所指（分支）。Git将这些快照称为“提交”，如下：

```
o <-- o <-- o <-- o <---- o
            ^            /
             \          v
              --- o <-- o
```

在上图中，每一个 o 代表着一个提交（即快照），在第三次提交之后，历史记录分支为两个独立的分支，类似于开发两个独立的功能，在之后的开发中，这些分支可能会被合并为一个新的快照。

Git中的提交是不可变的。如果在提交后发现的错误需要修改，看似对提交历史的编辑，实际上是在创建了一个全新的提交。

### 用伪代码来表示数据模型

```
//文件（blob）就是一堆字节
type blob = array<byte>

//目录（tree）包含了文件名和目录
type tree = map<string, tree | blob>

//提交（commit）具有父级，元数据和顶级树
type commit = struct{
	parents: array<commit>
	author: string
	message: string
	snapshot: tree
}
```

以上是用伪代码表示的一个干净简易的历史记录模型。

### 对象和内容寻址

对象可能是文件、目录或提交：

```
type object = blob | tree | commit
```

在Git的数据存储中，所有对象都是由 <u>SHA-1 hash</u> 散列进行内容寻址。

```
objects = map<string, object>

def store(object):
	id = sha1(object)
	objects[id] = object
	
def load(id):
	return objects[id]
```

### 引用（references）

通过 SHA-1 计算出的散列值是一个40位十六进制字符的字符串，它对应着每个快照，并予以区分。但咱们人类想要记住它是很困难的。

Git的解决办法是给SHA-1哈希值提供一个人类可读的名称，称为”引用“。引用是指向提交（commit）的指针。与不可变的对象不同，引用是可变的（通常更新指向最新的提交）。具体如下：

```
references = map<string, string>

def update_reference(name, id):
	references[name] = id

def read_reference(name):
	return references[name]

def load_reference(name_or_id)
	if name in references:
		return load(references[name])
	else:
		return load(id)
```

这样，Git就可以使用人类可读便记的方式来引用历史记录的特定快照，而不是无规律的哈希值。

### 本地仓库（Repository）

在磁盘上，Git存储的都是数据对象和引用,这也是Git数据模型的全部内容。而git的所有命令，则是通过添加对象或更新引用来对提交（commit）这个DAG图进行操作。

### 暂存区（Staging area）

在介绍暂存区前，先简要叙述一下git的工作原理，如下图：

![image-20230904100018420](C:\Users\24518\AppData\Roaming\Typora\typora-user-images\image-20230904100018420.png)

1. clone（克隆）：从远程仓库中克隆代码到本地仓库。
2. checkout（检出）：从本地仓库中选取一个分支到工作区进行修改，若工作区已有该文件，则会覆盖原有的修改，所以有时也会用作丢弃当前修改。
3. add（添加）：在提交前，先将代码提交到==暂存区==。
4. commit（提交）：提交到本地仓库。仓库中维护代码修改的各个历史版本（DAG图）。
5. fetch（抓取）：从远程库抓取到本地仓库，不进行任何的合并操作。
6. pull（拉取）：从远程库拉到本地仓库，自动进行合并，然后放到工作区，相当于fetch+merge。
7. push（推送）：代码完成后，推送到远程仓库，和团队成员共享代码。

为什么需要暂存区呢？其实这是git为了让我们更加精确的对代码实施控制。假设我们在debug时，整个代码中添加了很多打印语句，而我想要提交错误修复后的代码，但同时丢弃打印语句，这时暂存区的作用就显现出来了。
