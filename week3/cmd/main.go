package main

import "fmt"

func main() {

	// ch := make(chan int)

	// go func() {
	// 	for i := 0; i < 20; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()

	// for i := range ch {
	// 	fmt.Println(i)
	// }

	// phrase := "These are the times that try men's souls\n"

	// words := strings.Split(phrase, " ") // [These, are, ]

	// ch := make(chan string, len(words))

	// for _, word := range words {
	// 	ch <- word
	// }

	// close(ch)

	// for i := 0; i < len(words); i++ {
	// 	fmt.Print(<-ch + ", ")
	// }

	// ch <- "test"

	in := make(chan int)
	go generate(in)
	for {
		prime := <-in
		fmt.Println(prime)
		out := make(chan int)
		go filter(in, out, prime)
		in = out
	}
}

func generate(in chan int) {
	for i := 2; ; i++ {
		in <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

/* Goroutines
func main() {
	runtime.GOMAXPROCS(2)
	godur, _ := time.ParseDuration("10ms")
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello ")
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutines")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
*/
/**
func main() {
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	err := export(twtr, "twtrdata.txt")
	err = export(fb, "fbdata.txt")

	if err != nil {
		panic(err)
	}

	// twtr := new(twitter.Twitter)
	// lnkdin := new(linkedin.Linkedin)

	// ScrollFeeds(fb, twtr, lnkdin)
}

// ScrollFeeds prints all social media feeds
func ScrollFeeds(platforms ...week3.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("=================================")
	}
}

func export(u week3.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}
*/
