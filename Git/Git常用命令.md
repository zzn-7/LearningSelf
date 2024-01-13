# Git常用命令

### 基础

- git help \<command>：获取git某个命令的使用帮助
- git init：新建一个git仓库，数据存储在.git目录中
- git status：告诉你发生了什么事
- git add \<filename>：将文件添加到暂存区
- git commit：创建一个新的提交
- git log：显示历史日志信息
- git log --all --graph --decorate：将历史可视化为DAG图
- git diff \<filename>：显示某文件相对于暂存区所作的更改
- git diff \<revision> \<filename>：显示快照之间某文件的差异
- git checkout \<revision>：更新HEAD和当前的分支



### 分支与合并

- git branch：显示分支信息
- git branch \<name>：创建一个分支
- git checkout -b \<name>：创建一个分支并切换到它；相当于git branch \<name>然后 git checkout \<name>
- git merge \<revision>：合并到当前分支
- git mergetool：使用一个工具来解决合并时的冲突问题



### 远程

- git remote：列出远程仓库
- git remote add \<name> \<url>：添加一个远程仓库
- git push \<remote>  \<local branch> : \<remote branch>：将对象递送到远程仓库，并更新远程的引用
- git branch --set-upstream-to=\<remote>/ \<remote branch>：设置本地和远程分支之间的对应关系
- git fetch：从远程库抓取到本地仓库，不进行任何的合并操作
- git pull：从远程库拉到本地仓库，自动进行合并，然后放到工作区，相当于fetch+merge
- git clone：从远程仓库下载到本地仓库



### 撤销

- git reset HEAD \<file>：取消暂存文件
- git checkout -- \<file>：放弃更改



### 相关资源

- [Pro Git]([Git - Book (git-scm.com)](https://git-scm.com/book/en/v2))：建议阅读第1-5章。
- [Learn Git Branching](https://learngitbranching.js.org/)：是一个在浏览器中运行的游戏，以游戏的方式教你Git。

ps：这个游戏网站挺有趣的哈哈哈。