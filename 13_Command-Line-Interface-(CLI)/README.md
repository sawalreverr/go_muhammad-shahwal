# Basic Command Line Interace (CLI)

## Apa itu CLI ?

CLI adalah mekanisme interaksi dengan sistem operasi atau perangkat lunak komputer dengan mengetikkan perintah untuk menjalankan tugas tertentu

## Kenapa harus CLI ?

Kontrol yang leluasa terhadap OS, Pengelolaan sistem menjadi lebih cepat, Memiliki kemampuan untuk melakukan otomasi dengan sebuah script, dll

## Intro to UNIX Shell

```bash
me@linuxbox:~$
```

me -> your username\
linuxbox -> your hostname\
~ -> your current path (home)\
$ -> shell for normal user\
\# -> shell for root user

## Normal User vs Root User

Normal user = Allowed to manipulating **/home/user/** dir only\
Root user = Allowed to manipulating **/** dir\
Normal user + sudo = Allowed to act as root temporary

## Popular command in UNIX

### Directory

- pwd (print working directory) -> mengetahui directory kamu saat ini
- ls (list) -> menampilkan list directory kamu saat ini
- mkdir (make dir) -> membuat folder
- cd (change directory) -> mengganti directory kamu
- rm (remove) -> menghapus file / directory
- cp (copy) -> copy file / folder
- mv (move) -> memindakan file / folder atau bisa juga untuk rename
- ln (symbolic link) -> membuat shortcut

### Files

- create: touch
- view: cat, head, tail, less
- editor: nano, vim
- permission: chmod, chown
- different: diff

## Network

- ping
- ssh
- netstat
- nmap
- ip addr (ifconfig)
- wget
- curl
- telnet
- netcat

## Utility

- man
- env
- echo
- date
- which
- watch
- sudo
- history
- grep
- locate
