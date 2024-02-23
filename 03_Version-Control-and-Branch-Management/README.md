# Version Control and Branch Management (GIT)

### Apa itu Versioning ?
Versioning adalah mengatur versi dari source code program.

### Apa itu Git ?
Git adalah salah satu version control system populer yang digunakan para developer untuk mengembangkan software bersama-sama.

----

### Get Started using Git

#### Setting up
```
# git config
$ git config --global user.name "Muhammad Shahwal"
$ git config --global user.email "sawalrever23@gmail.com"

# start with init
$ git init
$ git remote add origin <remote_name> <remote_repo_url>
$ git push -u origin <remote_name> <remote_repo_url>

# start with existing project, start working on the project
$ git clone https://github.com/sawalreverr/belajar-git.git
$ cd belajar-git

```

#### Staging area

working directory >> staging area >> repository

```
$ git status // cek status working directory

$ git add <directory> // menambah berdasarkan folder
$ git add file.go // menambah berdasarkan file
$ git add . // menambah semua file yang ada di working directory

$ git commit -m "your message"
```

#### Git diff and stash

```
# git diff
# change file
# add staging area
$ git diff --staged // melihat perubahan

# stashing your work
$ git stash // menyimpan sementara perubahan

# re-applying your stashed changes
$ git stash apply

```

#### Git log, checkout

```
# viewing an old version 
$ git log --oneline

# hash1 Continue doing crazy things
# hash2 Do some stuff
# hash3 Re-code our tools

$ git checkout hash2
```

#### Git reset
```
$ git reset hash1 --soft // uncommit changes, changes are left staged (index)

$ git reset hash2 --hard // uncommit + unstage + delete changes, nothing left
```

#### Git push, fetch & pull
```
# remote
$ git remote -v
$ git remote origin https://github.com/example/example-project.git

# fetch & pull
$ git fetch 
$ git pull origin master

# push
$ git push origin master 
$ git push origin feature/login-user // branching
```

#### Git branching 
```
# show all branch
$ git branch --list

# create a new branch called <branch>
$ git branch <name>

# force delete the specified branch
$ git branch -D <name>

# list remote branch 
$ git branch -a
```

#### Git merge 
```
# start a new feature
$ git checkout -b new-feature master

# edit some files
$ git add .
$ git commit -m "adding new feature"

# merge in the new-feature branch
$ git checkout master
$ git merge new-feature
$ git branch -d new-feature

```