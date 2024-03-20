#!/bin/bash

# membuat tiga folder 
mkdir foods drinks snacks

# membuat menu.txt untuk setiap folder
touch foods/menu.txt
touch drinks/menu.txt
touch snacks/menu.txt

# mengisi data pada menu.txt untuk setiap folder
echo "nasi goreng, ayam goreng, bubur ayam" >> foods/menu.txt
echo "kopi susu, susu oat, es teh" >> drinks/menu.txt
wget https://gist.githubusercontent.com/nadirbslmh/c84ee3527272e52a8e6257d46e627f91/raw/74593cc3fe61ca4ff36e11ed8e3fffcfb1d0c602/snacks.json -O snacks/menu.txt