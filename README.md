# go
learn go

# git

## git 使用规范

[Git 使用规范流程](https://ruanyifeng.com/blog/2015/08/git-use-process.html)

### 第一步：新建分支
- 首次开发新功能，都应该新建一个分支
```
#获取主干分支代码
git checkout master
git pull

#新建一个开发分支loyzhang
git checkout -b loyzhang
```

### 第二步：提交分支commit
- 分支修改后，就可以提交commit了

```
git add .
git status
git commit -- verbose
```
git add . 命令 表示保存所有变化
git status命令 用来查看发生变动的文件
git commit 命令的verbose参数，会列出diff的结果。

### 第三步：撰写提交信息
- 提交commit时，必须给出完整扼要的提交信息，下面是一个范本
[规范GIT代码提交信息](https://aotu.io/notes/2020/09/10/git-commit-control/index.html)


### 第四步：与主干同步
- 分支的开发过程中，要经常与主干保持同步
```
 git fetch origin
 git rebase origin/master
```
### 第五步：合并commit
分支开发完成后，很可能有一堆commit，但是合并到主干的时候，往往希望只有一个（或最多两三个）commit，这样不仅清晰，也容易管理。
那么，怎样才能将多个commit合并呢？这就要用到 git rebase 命令。
```
git rebase -i [分支hash]
```
git rebase命令的i参数表示互动（interactive），这时git会打开一个互动界面，进行下一步操作
当前我们只要知道 pick 和 squash 这两个命令即可。

pick 的意思是要会执行这个 commit
squash 的意思是这个 commit 会被合并到前一个commit
我们将 c4e858b5 这个 commit 前方的命令改成 squash 或 s，然后输入:wq以保存并退出

Pony Foo提出另外一种合并commit的简便方法，就是先撤销过去5个commit，然后再建一个新的。
```
git reset HEAD~5
$ git add .
$ git commit -am "Here's the bug fix that closes #28"
```













































## git 提交规范
#### 提交规范
Conventional Commits(约定式提交规范)，是目前使用最广泛的提交信息规范，其主要受AngularJS规范的启发,下面是一个规范信息的结构：
```
<type>[optional scope]: <subject>
//空一行
[optional body]
//空一行
[optional footer(s)]
```
#### 规范说明
- type 提交的类别，必须是以下类型中的一个
```
feat：增加一个新功能
fix：修复bug
docs：只修改了文档
style：做了不影响代码含义的修改，空格、格式化、缺少分号等等
refactor：代码重构，既不是修复bug，也不是新功能的修改
perf：改进性能的代码
test：增加测试或更新已有的测试
chore：构建或辅助工具或依赖库的更新
```
scope 可选，表示影响的范围、功能、模块

subject
必填，简单说明，不超过50个字

body
选填，用于填写更详细的描述

footer
选填，用于填关联issus,或BREAK CHANGE

BREAKING CHANGE

必须是大写，表示引入了破坏性 API 变更，通常是一个大版本的改动，BREAKING CHANGE: 之后必须提供描述，下面一个包含破坏性变更的提交示例
```
feat: allow provided config object to extend other configs

BREAKING CHANGE: `extends` key in config file is now used for extending other config files

```

#### 如何约束规范
- 怎么确保每个提交都能符合规范呢，最好的方式就是通过工具来生成和校验，commitizen是一个nodejs命令行工具，通过交互的方式
```
# 全局安装
npm install -g commitizen 
# 或者本地安装
$ npm install --save-dev commitizen
# 安装适配器
npm install cz-conventional-changelog

```
- 安装完成后可以使用git cz 来代替git commit,然后根据提示一步步输入即可
