# Penggunaan AI di Golang

## Kenapa memilih Go ?

Beberapa alasan memilih Go:

- Dikembangkan oleh Google untuk pengembangan aplikasi server-side dan berbasis cloud
- Go adalah compiled language shingga lebih cepat dibanding interpreted language (python, javascript)
- Mendukung concurrency melalu fitur goroutines sehingga dapat menjalankan banyak task secara simultan tanpa membebani memori
- Memili standard libraries yang sangat kaya, didukung oleh komunitas yang kuat
- Termasuk libraries untuk mendukung pengembangan AI

## Libraray Go untuk machine learning

GoLearn

- Bersifat open source
- 'battries included' machine learning library
- Simplicity & customizability
- Contoh penerapan:
  https://github.com/sjwhitworth/golearn#getting-started

Gorgonia

- Memudahkan menulis dan mengevaluasi persamaan matematika yang melibatkan array multidimensi
- Low-level library, seperti Theano, namun lebih tinggi dibanding Tensorflow
- Contoh penerapan:
  https://github.com/gorgonia/gorgonia#usage

gonum

- Sebuah set packages yang didesain untuk menulis numerical & algoritma sains menjadi produktif, memiliki performa tinggi, dan scalable
- Mirip seperti numpy dan scipy, library yang dibuat menggunakan python
- Referensi:
  https://www.gonum.org/post/intro_to_gonum

gomind

- Contoh penggunaan
  https://github.com/surenderthakran/gomind#usage

## AI as a Service (AIaaS)

Sebuah tools AI yang dapat segera digunakan (pre-built or off-the-shelf AI tools), berguna bagi bisnis yang ingin menerapkan AI tanpa merekrut ahlinya dan tanpa mengeluarkan biaya yang relatif banyak.

Beberapa perusahaan penyedia AIaaS, contohnya:

- Amazon Web Services (AWS)
- Microsoft Azure
- Google Cloud Platform (GCP)
- IBM
- OpenAI

Tipe-tipe AIaaS

- Bots/Chatbots: create AI-based customer services
  - Amazon Lex (AWS), Azure Bot Service (Microsoft Azure), DialogFlow (GCP)
- APIs: mengintegrasikan AI milik vendor dengan aplikasi sendiri
  - Amazon Recognition, Amazon Comprehend
  - Azure Cognitive Service, Azure Speech Services
  - Google Cloud Vision, Cloud Natural Language
- Machine learning: build & deploy ML model easily
  - Amazon SageMaker
  - Azure Machine Learning
  - Goolge Cloud AI

Keuntungan

- Pengurangan cost
- Low-code
- Proses deployment cepat
- Flexibility
- Usabilitiy
- Scalability
- Customization

Kelemahan

- Cost yang berlebih dalam jangka panjang
- Privasi data
- Keamanan
- Transparansi
- Vendor lock-in

## OpenAI API

Target:

- Mendesain prompt yang tepat (prompt engineering)
- Membuat aplikasi berbasis AI menggunakan API OpenAI

[Dokumentasi OpenAI API](https://platform.openai.com/docs/introduction)

[OpenAI API Reference](https://platform.openai.com/docs/api-reference/introduction)

Opsional: install library [Go OpenAI](https://github.com/sashabaranov/go-openai)

### OpenAI API Pricing

- Start for free, mendapat free credit senilai $5 yang dapat digunakan selama 3 bulan pertama sejak register
- Pay as you go, setiap model memiliki tarif yang berbeda-beda
- Lihat lebih detail di
  https://openai.com/pricing
- Cek seberapa banyak penggunaan di
  https://platform.openai.com/account/usage
- Best practices manajemen billing:
  https://platform.openai.com/docs/guids/production-best-practices/managing-billing-limits

## Prompt Engineering

Prinsip 1: Gunakan prompt yang jelas & spesifik

- Gunakan delimiters (contoh: ```, ' ", <>, <tag></tag>) untuk menandakan bagian mana yang menjadi input
- Tuliskan struktur output yang diinginkan
- Minta model untuk memeriksa apakah input sudah sesuai
- Few-shot prompting

Prinsip 2: Berikan "waktu berpikir" untuk menghindari solusi yang salah

- Menulis langkah-langkah yang dibutuhkan untuk menyelesaikan tugsa dan output yang diinginkan
- Mengintruksikan model untuk menuliskan solusinya sendiri sebelum masuk ke kesimpulan
