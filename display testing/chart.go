package main

type Chart struct {
	songName        string
	songArtist      string
	songYear        int
	songTrackGuitar string
}

type Note struct {
	// 0 is open
	// 1 and up are normal (yes this goes on forever, defined by the chart)
	value  int
	length int
	forced bool
	tap    bool
}

func main() {

}
