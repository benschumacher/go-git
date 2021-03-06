package main

import (
	"os"

	"srcd.works/go-git.v4"
	. "srcd.works/go-git.v4/_examples"
)

func main() {
	CheckArgs("<url>", "<directory>")
	url := os.Args[1]
	directory := os.Args[2]

	// Clone the given repository to the given directory
	Info("git clone %s %s", url, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:   url,
		Depth: 1,

		// as git does, when you make a clone, pull or some other operations the
		// server sends information via the sideband, this information can being
		// collected provinding a io.Writer to the CloneOptions options
		Progress: os.Stdout,
	})

	CheckIfError(err)
}
