/*
GetIndex - получение номеров  из консоли
Запросы по адресам
Парсинг
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

type xkcd struct {
	Num        int    `json:"num"`
	Transcript string `json:"transcript"`
}

const base = "https://xkcd.com/"
const baseEnd = "/info.0.json"

func getIndex(indexSlice []string) []string {
	sc := bufio.NewScanner(os.Stdin)
	var mask = regexp.MustCompile(`([[:digit:]]{3})`)
	for sc.Scan() {
		tempStr := sc.Text()
		if tempStr == "0" {
			break
		}
		sub := mask.FindString(tempStr)
		tempStr = sub
		indexSlice = append(indexSlice, tempStr)
		//fmt.Println(tempStr)
	}
	return indexSlice
}

/*func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(base + r.URL.Path[1:])
}*/

func main() {
	indexSlice := []string{}

	indexSlice = getIndex(indexSlice)

	//fmt.Println(indexSlice)
	for i := 0; i < len(indexSlice); i++ {
		url := base + indexSlice[i] + baseEnd
		fmt.Println(url)
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}

		defer resp.Body.Close()

		var comic xkcd

		if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}

		fmt.Println(comic.Transcript)
	}
}
