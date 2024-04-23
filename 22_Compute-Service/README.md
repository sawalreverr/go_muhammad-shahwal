# Deployment

## What is Deployment ?

Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. Penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web & api ke server sedangkan aplikasi mobile ke Playstore/Appstore.

## Strategy Deployment

- Big-Bang/Replace Deployment

  Big-Bang strategy, atau sering juga disebut seperti Replace deployment, atau Recreate, dsb. Yang dimana sifatnya menimpa dan mengganti (mereplace) aplikasi yang aktif secara langsung.

  Kelebihan:

  - Mudah di-implementasikan. Cara klasik, tinggal replace.
  - Perubahan kepada sistem langsung 100% secara instan.

  Kekurangan:

  - Terlalu beresiko, rata-rata downtime cukup lama.

- Rollout Deployment

  Dengan metode ini, kita melakukan deployment secara bertahap per-server yang hidup. Dan jika satu server saja langsung error, kita dapat langsung rollback tanpa melanjut deploy kesemua server.

  Kelebihan:

  - Lebih aman dan less downtime dari versi sebelumnya

  Kekurangan:

  - Akan ada 2 versi aplikasi berjalan secara barengan sampai semua server terdeploy, dan bisa membuat bingung. Seperti kita ketahui, untuk management versioning itu sedikit susah dan cukup buat pusing kepala
  - Karena sifatnya perlahan satu persatu, untuk deployment dan rollback lebih lama dari yang Bigbang, karena prosesnya perlahan-lahan sampai semua server terkena efeknya.
  - Tidak ada kontrol request. Server yang baru ke-deploy dengan aplikasi versi baru, langsung mendapat request yang sama banyaknya dengan server yang lain. Sehingga jika terjadi error, juga dapat menyebabkan kerugian besar.

- Blue/Green Deployment

  Pada strategi ini konsepnya cukup sederhana, pertama-tama, kita akan membuat satu environment yang serupa dengan yang sedang aktif/live, kemudian kita pun melakukan switching request ke environment baru tersebut.

  Kelebihan:

  - Perubahan sangat cepat, sekali switch service langsung berubah 100%.
  - Tidak ada issue beda versi pada service seperti yang terjadi pada Rollout Deployment.

  Kekurangan:

  - Resource yang dibutuhkan lebih banyak. Karena untuk setiap deployment kita harus menyediakan service yang serupa environmentnya dengan yang sedang berjalan di production.
  - Testing harus benar-benar sangat diprioritaskan sebelum di switch, aplikasi harus kita pastikan aman dari request yang tiba-tiba banyak.

- Canary Deployment

  Strategi ini lebih advance dari semua metode release tersebut diatas. Prinsip kerjanya mirip seperti Rollout Deployment, tetapi bedanya, jika pada Rollout Deployment, ketika aplikasi di deploy pada satu server, maka server tersebut akan langsung kebagian request dari user sama rata dengan server lainnya.

  Nah pada canary, aplikasi juga akan di deploy pada satu server, tetapi bedanya adalah, request yang masuk kedalamnya akan dikontrol dan ditahan berdasarkan kebutuhan.

  Kelebihan:

  - Cukup aman
  - Mudah untuk rollback jika terjadi error/bug, tanpa berimbas kesemua user

  Kekurangan:

  - Untuk mencapai 100% cukup lama dibanding dengan Blue/Green deployment. Dengan Blue/Green deployment, aplikasi langsung 100% terdeploy keseluruh user.
