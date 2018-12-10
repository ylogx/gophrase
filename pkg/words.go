package pkg

import (
	"net/http"
	"io"
	"os"
	"fmt"
	"bufio"
	"strings"
)

const (
	//Url = "https://github.com/ylogx/english-words/blob/master/words.txt?raw=true"
	//FileSize = 4862966
	//FileName = "gophrase_words_all.txt"
	Url           = "https://github.com/ylogx/english-words/blob/master/words_alpha.txt?raw=true"
	FileSize      = 4234866
	FileName      = "gophrase_words_alpha.txt"
	WordMinLength = 3
	WordMaxLength = 10
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Knows the url and corresponding file size
func ensureWordListExists(filename string) error {
	fileInfo, err := os.Stat(filename)
	if err != nil || fileInfo.Size() != FileSize {
		fmt.Printf("Downloading file\n")
		err = DownloadFile(filename, Url)
		if err != nil {
			return err
		}
		fmt.Printf("Finished Downloading Successfully\n")
	}

	return nil
}

func ReadWords() []string {
	err := ensureWordListExists(FileName)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	var words []string
	for {
		//word, err := reader.ReadString(10) // 0x0A separator = newline
		word, err := reader.ReadString(0x0A)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err) // if you return error
		}
		word = string(word)
		word = strings.TrimSuffix(word, "\r\n")
		if WordMinLength <= len(word) && len(word) <= WordMaxLength {
			words = append(words, word)
		}
	}
	//words, err := ioutil.ReadFile(FileName)
	if err != nil {
		panic(err)
	}
	return words
}
