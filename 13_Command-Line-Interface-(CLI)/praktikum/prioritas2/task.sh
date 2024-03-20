#!bin/bash

echo "Go Project Generator"
read -p "Nama project tanpa spasi: " namaproject
read -p "simple/api: " jenisproject

if [ $jenisproject = "simple" ]; then
    mkdir $namaproject
    cd $namaproject || exit
    go mod init $namaproject
    touch main.go

    echo "  Sukses!!"
elif [ $jenisproject = "api" ]; then
    mkdir $namaproject
    cd $namaproject || exit
    go mod init $namaproject
    mkdir routes models repositories services configs controllers
    touch main.go

    echo "Sukses!!"
else
    echo "hanya simple/api"
fi
