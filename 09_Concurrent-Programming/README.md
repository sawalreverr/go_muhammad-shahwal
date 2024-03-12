## Concurrent Programming

### Sequential Program

Didalam konsep sequantial program sebelum memulai task baru, task sebelumnya harus diselesaikan terlebih dahulu.

Sebagai contoh pada tampilan sebuah website yang terdapat video, website, image.. Misal disini ditetapkan untuk video terlebih dahulu yang di tampilkan maka video penampilan untuk video akan dimulai terlebih dahulu kemudian dilanjutkan dengan tampilan website dan terakhir baru image.

video (2 sec) -> website (3 sec) -> image (4 sec)

### Parallel Program

Didalam konsep parallel program semua task yang diperlukan dapat dilakukan secara bersamaan, berbeda seperti sequential yang diharuskan untuk menunggu task sebelumnya hingga selesai disini kita tidak harus menunggu melainkan kita bisa langsung mengerjakannya secara bersamaan. (Multi-core machines dibutuhkan)

### Concurrent Program

Konsep Concurrent itu sama seperti konsep parallel sebelumnya, kita dapat mengerjakan task secara bersamaan dan dapat dikerjakan secara independent (tidak mewajibkan multicore).

#### Concurrent dapat dijalankan secara Parallel

Dengan Go concurrency (goroutines) kita dapat membuat concurrent secara parallel dengan mudah yang dapat memberikan kelebihan pada multiple processors

## Feature Go

- Concurrent execution (Goroutines)
- Synchronization and messaging (Channels)
- Multi-way concurrent control (Select)

### What is Goroutine ?

Goroutine adalah fungsi atau metode yang berjalan secara concurrent (independent) dengan fungsi atau metode lainnya.

Biaya yang diperlukan dalam membuat sebuah goroutine itu kecil dibandingkan dengan sebuah thread. Thread adalah lightweight proccess atau dalam arti lain ialah thread adalah sebuah unit yang mengeksekusi kode dibawah program (user and kernel levels)

### Example goroutine

```go
// Using goroutine and threads

func Hello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go Hello("goroutine") // using goroutine
	Hello("thread")       // using thread
}
```

```go
// Using multiple goroutine
var start time.Time

func init() {
	start = time.Now()
}

func GetChars(sentence string) {
	for _, c := range sentence {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func GetDigits(digits []int) {
	for _, d := range digits {
		fmt.Printf("%v at time %v\n", d, time.Since(start))
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	fmt.Printf("\nmain execution started at time %v\n\n", time.Since(start))

	go GetChars("Hello")
	go GetDigits([]int{1, 2, 3})

	time.Sleep(200 * time.Millisecond)
	fmt.Printf("\nmain execution stopped at time %v\n", time.Since(start))

}
// main execution started at time 0s

// 1 at time 0s
// H at time 0s
// e at time 22.1061ms
// l at time 37.6714ms
// 2 at time 37.6714ms
// l at time 53.214ms
// 3 at time 69.0531ms
// o at time 69.0531ms

// main execution stopped at time 208.13ms
```

### GOMAXPROCS

GOMAXPROCS digunakan untuk mengatur jumlah dari operating systemthreads yang tersedia untuk mengeksekusi goroutine ke program yang berbeda-beda.

```sh
$ time GOMAXPROCS=1 go run main.go
```

### Channel and Select

Channel adalah sebuah alat komunikasi yang digunakan goroutine untuk bisa berkomunikasi antar yang lain

```go
// Simple Example and Declaration
func Greet(c chan string) {
	data := <-c // storing data from channel c
	fmt.Printf("Hello %v!\n", data)
}

func main() {
	fmt.Println("main() started")

	c := make(chan string)
	go Greet(c)
	c <- "Ilham" // push or write data to the channel c

	fmt.Println("main() stopped")
}
```

```go
// Example using channel

func main() {
	theMine := [5]string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)

	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			oreChan <- item // send
		}
	}(theMine)

	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan // receive
			fmt.Printf("Miner: Received %v from finder\n", foundOre)
		}
	}()

	<-time.After(time.Second * 5) // Again, ignore this for now
}
```

### Unbuffered and Buffered Channels

Perbedaan antara Unbuffer dan Buffer pada sebuah Channel adalah bagaimana Channel itu dapat menampung data, pada unbuffered channel data yang diterima dan dikirim hanya dapat sekali sedangkan pada Buffered channel kita dapat mengatur jumlah data yang dapat diterima dan dikirim.

```go
// Example buffered channel
func main() {
	messages := make(chan string, 2) // buffered channel

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
```

### Select

Select digunakan untuk mempermudah dalam mengontrol komunikasi antar data pada channels

```go
// Select Example
func GetAverage(numbers []int, ch chan float64) {
	var sum = 0

	for _, e := range numbers {
		sum += e
	}

	ch <- float64(sum) / float64(len(numbers))
}

func GetMax(numbers []int, ch chan int) {
	var max = numbers[0]

	for _, e := range numbers {
		if max < e {
			max = e
		}
	}

	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)
	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 3, 4, 6, 3}
	fmt.Printf("Numbers\t: %v\n", numbers)

	var ch1 = make(chan float64)
	go GetAverage(numbers, ch1)

	var ch2 = make(chan int)
	go GetMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %v \n", max)
		}
	}
}

// Numbers : [3 4 3 5 6 3 2 2 3 4 6 3]
// Max     : 6
// Avg     : 3.67
```

### Race Condition

Race Condition adalah kondisi ketika 2 thread mengakses memori secara bersamaan, salah satunya adalah writing. Race condition terjadi dikarenakan akses yang tidak tersinkronisasi ke shared memory.

```go
// Example
func GetNumber() int {
	var i int

	go func() {
		i = 5
	}()

	return i
}

func main() {
	fmt.Println(GetNumber()) // 0
}
// run : go run -race main.go
```

### Fixing data race

#### Waitgroups

Cara paling mudah untuk menyelesaikan sebuah data race, adalah dengan ngeblok read access sampai operasi write diselesaikan.

```go
func GetNumber() int {
	var i int

	var wg sync.WaitGroup // initialize waitgroup variable

	wg.Add(1) // Add(1) -> there is 1 task that we need to wait for
	go func() {
		i = 5
		wg.Done() // Done -> indicates that we are done with the task we are waiting for
	}()

	wg.Wait() // Wait -> blocks until Done() is called the same number of times as the amount of task we have Add(1)
	return i
}

func main() {
	fmt.Println(GetNumber()) // 5
}
```

#### Channel Blocking

Ini memungkinkan goroutine kita untuk saling menyinkronkan satu sama lain sementara.

```go
func GetNumber() int {
	var i int

	done := make(chan struct{}) // create a channel to push an empty struct to once we're done

	go func() {
		i = 5
		done <- struct{}{} // push an empty struct once we're done
	}()

	<-done // this statement blocks until something gets pushed into the done channel
	return i
}

func main() {
	fmt.Println(GetNumber()) // 5
}
```

#### Mutex

Kondisi dimana kita tidak peduli dengan urutan reads dan writes, kita hanya mensyaratkan bahwa keduanya tidak terjadi secara bersamaan.
Jika kondisi diatas terdengar sama dengan kasus kamu, maka kita dapat mempertimbangkan untuk menggunakan sebuah mutex.

```go
// First, create a struct that contains the value we want to return, along with a mutex instance
type SafeNumber struct {
	val int
	m   sync.Mutex
}

func (i *SafeNumber) Get() int {
	// the lock method of the mutex blocks if it is already locked
	// if not, then it blocks other calls until the Unlock method is called
	i.m.Lock()

	// defer Unlock until this method returns
	defer i.m.Unlock()

	// return the value
	return i.val
}

func (i *SafeNumber) Set(val int) {
	// Similiar to the Get method, except we Lock until we are done
	// writing to i.val
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

func GetNumber() int {
	// Create an instance of Safenumber
	i := &SafeNumber{}

	// use Set and Get instead of regular assignment and reads
	// we can now be sure that we can read only if the write has completed, or vice versa
	go func() {
		i.Set(5)
	}()

	time.Sleep(time.Second)
	return i.Get()
}

func main() {
	fmt.Println(GetNumber()) // 5
}
```
