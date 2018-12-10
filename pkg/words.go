package pkg

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	//Url = "https://github.com/ylogx/english-words/blob/master/words.txt?raw=true"
	//FileSize = 4862966
	//FileName = "gophrase_words_all.txt"
	Url           = "https://github.com/ylogx/english-words/blob/master/words_alpha.txt?raw=true"
	FileSize      = 4234866
	FileName      = "/tmp/gophrase_words_alpha.txt"
	WordMinLength = 3
	WordMaxLength = 10
)

type Vocabulary interface {
	words() []string
	size() int
}

type english struct {
	filename      string
	wordsData     []string
	WordMinLength int
	WordMaxLength int
}

func NewEnglishVocabulary(filename string) Vocabulary {
	return &english{
		filename:      filename,
		wordsData:     []string{},
		WordMinLength: WordMinLength,
		WordMaxLength: WordMaxLength,
	}
}

func (e *english) words() []string {
	e.ensureWordListExists()
	return e.wordsData
}

func (e *english) size() int {
	e.ensureWordListExists()
	return len(e.wordsData)
}

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
func (e *english) ensureWordListExists() error {
	if len(e.wordsData) > 0 {
		return nil
	}
	fileInfo, err := os.Stat(e.filename)
	if err != nil || fileInfo.Size() != FileSize {
		fmt.Printf("Downloading file\n")
		err = DownloadFile(e.filename, Url)
		if err != nil {
			return err
		}
		fmt.Printf("Finished Downloading Successfully\n")
	}
	e.readWordsFromFile()
	return nil
}

func (e *english) readWordsFromFile() {
	file, err := os.Open(e.filename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
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
		if e.WordMinLength <= len(word) && len(word) <= e.WordMaxLength {
			e.wordsData = append(e.wordsData, word)
		}
	}
	//words, err := ioutil.ReadFile(FileName)
	if err != nil {
		panic(err)
	}
}
