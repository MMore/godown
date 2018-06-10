package main

import (
	"fmt"
	sf "github.com/manyminds/gosfml"
	"gopkg.in/alecthomas/kingpin.v1"
	"os"
	"time"
)

type mint int
type FullTime struct {
	hours   mint
	minutes mint
	seconds mint
}

func (i mint) String() string {
	return fmt.Sprintf("%d", i)
}

func playSound(filepath string) {
	data := MustAsset("resources/" + filepath)
	buffer, err := sf.NewSoundBufferFromMemory(data)
	if err != nil {
		panic("soundbuffer not loaded")
	}

	sound := sf.NewSound(buffer)
	sound.Play()
}

func sayTime(time string, lang string, female bool) {
	var mf string = "m"
	if female {
		mf = "f"
	}
	playSound(lang + "/" + mf + "/" + time + ".wav")
}

func separateDuration(duration time.Duration) (int, int, int) {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) - hours*60
	seconds := int(duration.Seconds()) - hours*60*60 - minutes*60

	return hours, minutes, seconds
}

func countDown(time FullTime) FullTime {
	time.seconds -= 1

	if time.seconds == -1 {
		time.seconds = 59
		time.minutes -= 1

		if time.minutes == -1 {
			time.minutes = 59
			time.hours -= 1
		}
	}
	return time
}

func tick(c chan FullTime, initialTime FullTime) {
	var countingTime FullTime = initialTime
	for {
		c <- countingTime
		if countingTime.hours == 0 && countingTime.minutes == 0 && countingTime.seconds == 0 {
			break
		} else {
			countingTime = countDown(countingTime)
			time.Sleep(time.Second)
		}
	}
}

func onTick(c chan FullTime, lang string, female bool) {
	for {
		var counter FullTime = <-c
		fmt.Printf("%dh %sm %ss\n", counter.hours, counter.minutes, counter.seconds)

		if counter.hours == 0 {
			if counter.seconds == 0 {
				// Final
				if counter.minutes == 0 {
					fmt.Println("FINISHED! Enter any key to exit.")
					playSound("finale/reminder.wav")
					break
				} else if counter.minutes == 5 {
					sayTime(counter.minutes.String()+"m", lang, female)
				} else if counter.minutes == 3 {
					sayTime(counter.minutes.String()+"m", lang, female)
				} else if counter.minutes == 2 {
					sayTime(counter.minutes.String()+"m", lang, female)
				} else if counter.minutes == 1 {
					sayTime(counter.minutes.String()+"m", lang, female)
				}
			}

			if counter.minutes == 0 {
				if counter.seconds == 30 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 20 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 10 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 9 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 8 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 7 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 6 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 5 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 4 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 3 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 2 {
					sayTime(counter.seconds.String()+"s", lang, female)
				} else if counter.seconds == 1 {
					sayTime(counter.seconds.String()+"s", lang, female)
				}
			}
		}
	}
}

var (
	app      = kingpin.New("godown", "A talking countdown command-line application.")
	duration = app.Arg("duration", "duration of the countdown (i.e. 1m30s)").Required().Duration()
	female   = app.Flag("female", "male or female voice").Bool()
	language = app.Flag("lang", "english (en) or german (de) language").Default("en").Enum("en", "de")
)

func main() {
	app.Version("1.0.1")
	kingpin.MustParse(app.Parse(os.Args[1:]))

	h, m, s := separateDuration(*duration)
	var fulltime FullTime = FullTime{hours: mint(h), minutes: mint(m), seconds: mint(s)}
	fmt.Println("Starting talking countdown with", *duration)
	fmt.Println("Enter any key to abort")
	fmt.Println()

	var ticker chan FullTime = make(chan FullTime)
	go tick(ticker, fulltime)
	go onTick(ticker, *language, *female)

	var input string
	fmt.Scanln(&input)
}
