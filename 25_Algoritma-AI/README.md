# Algoritma AI

## Tipe-tipe AI

Berdasarkan level kecerdasan:

- Narrow/weak AI
- General/strong AI
- Super AI

Berdasarkan fungsionalitas:

- Reactive machine
- Limited memory
- Theory of mind
- Self-awareness

Saat ini, tipe AI yang umum digunakan adalah narrow AI atau reactive machine & limited memory. Tipe lainnya masih sebatas teori atau masih tahap pengembangan

### Narrow/Weak AI

- AI untuk melaksanakan tugas yang spesifik
- Kemampuannya sangat dioptimalisasi untuk sebuah tugas dan dapat melampaui kemampuan manusia
- Contoh:
  - Facial Recognition
  - Natural Language Processing (ChatGPT)
  - Autonomous vechicle
  - Voice Assistant (Siri, Alexa)

### Reactive machine & limited memory

Saat ini, reactive machine dan limited memory AI sudah diciptakan

- Reactive machine: tipe AI yang melakukan reaksi terhadap inout tanpa pengetahuan masa lampau dan hanya spesifik melakukan tuas tertentu. Contoh: machine learning, recommendation engine, deep blue (mesin catur IBM)
- Limited memory: pengembangan dari reactive machine, sudah dilengkapi dengan kemampuan mengevaluasi situasi masa lampau. Contoh: deep learning, reinforcement learning, self-driving cars

## What is Algoritma ?

Algoritma adalah sebuah set intruksi yang diberikan kepada komputer untuk memecahkan masalah atau mencapai sebuah tujuan, contoh: menghitung luas segitiga

Dalam konteks AI, algoritma digunakan agar komputer mampu:

- Melakukan kalkulasi & pemrosesan data
- Melatih model machine learning & deep learning
- Membuat prediksi berdasarkan data yang diberikan

## What is Machine Learning ?

- Subset dari AI yang fokus pada pengembangan algoritma dan model
- Memiliki kemampuan untuk belajar, dan mengambil keputusan atau prediksi tanpa perlu diprogram secara eksplisit
- Kualitas dari model machine learning tergantung dari bagus/tidaknya data yang diberikan

## What is Supervised Learning ?

- Model machine learning di training pada data berlabel, dimana nilai input - output sudah diketahui
- Algoritma machine learning mempelajari cara me-mapping sebuah function dari input untuk menghasilkan output, Y = f(X)
- Ada dua tipe supervised learning: regression (output berupa variabel kontinu/angka) dan classification (output berupa kategori)

Algoritma supervised learning classification:

- Logistic regression
- Decision tree
- Support vector machine (SVM)
- K-nearest neighbor (KNN)
- Random forest
- Naive Bayes

Algoritma supervised learning regression:

- Linear regression
- Polynimial regression
- Support vector regression
- Decision tree regression
- Random forest regression

## What is Unsupervised Learning ?

- Model machine learning di training pada data tidak berlabel
- Mencari pola, mengelompokkan data berdasarkan kesamaan atribut pada dataset yang diberikan dan membuat keputusan berdasarkan apa yang ditemukan

Algoritma unsupervised learning

Clustering:

- K-means
- Density-based (DBSCAN)
- Gaussian mixture model
- Balanced iterative reducing and clustering using hierarchies (BIRCH)
- Affinity propagation
- Mean-shift
- Ordering points to identify the clustering structure (OPTICS)
- Agglomerative hierarchy

Association rule learning:

- Apriori algorithm
- Eclat algorithm
- Frequent pattern - growth algorithm

Dimensional reduction:

- Principal component analysis (PCA)
- Singular value decomposition (SVD)
- Linear discrimination analysis (LDA)
- Non-negative matrix factorization (NMF)
- t-distributed stochastic neighbor embedding (t-sne)
- Factor analysis
- Independent component analysis (ICA)

Anomaly detection

- Isolation forest
- One-class SVM
- Local outlier facotr (LOF)

## What is Deep Learning ?

- Subset dari machine learning
- Menggunakan artificial neural network untuk menirus proses pembelajaran manusia
- Dibandingkan machine learning, dapat membangun korelasi antar data dengan lebih kompleks, namun membutuhkan dataset yang relatif lebih besar
- Idealnya menggunakan GPU (graphic processing unit) dalam melatih model deep learning

Algoritma deep learning:

- Multilayer Perceptrons (MLPs)
- Self Organizing Maps (SOMs)
- Deep Belief Networks (DBNs)
- Restricted Boltzmann Machines (RBMs)
- Autoencoders
- Convolutional Neural Networks (CNNs)
- Long Short Term Memory Networks (LSTMs)
- Recurrent Neural Networks (RNNs)
- Generative Adversarial Networks (GANs)
- Radial Basis Function Networks (RBFNs)

## Machine learning vs Deep Learning

Pada Machine Learning:
Input (Dog) -> Feature Extraction -> Classification -> Output (Dog / Not Dog)

Pada Deep Learning:
Input (Dog) -> Feature Extraction + Classification -> Output (Dog / Not Dog)

Deep learning tidak membutuhkan feature extraction secara manual

## Frameworks & libraries untuk AI

Python:

- Pandas
- Tensorflow
- Keras
- NumPy
- Scikit-learn
- PyTorch

Java:

- Deep Java Library (DJL)
- Deeplearning4j
- Apache OpenNLP
- Java Machine Learning Library (Java-ML)
- RapidMiner
- Weka
- Encog Machine Learning Framework

Go:

- golearn
- gorgonia
- eaopt
- gonum
- gomind

Flutter:

- flutter_openai
- edge_detection
- tflite_flutter
- google_ml_kit
