package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

const csvFileName = "posts.csv"

func main() {
	start := time.Now()
	file, err := os.Open(csvFileName)
	if err != nil {
		allPosts := [4]Post{
			{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
			{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
			{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
			{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
		}

		// CSVファイル作成フェーズ
		csvFile, err := os.Create(csvFileName)
		if err != nil {
			panic(err)
		}
		defer csvFile.Close()

		// CSVファイル書き込みフェーズ
		writer := csv.NewWriter(csvFile)
		for _, post := range allPosts {
			line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
			err := writer.Write(line)
			if err != nil {
				panic(err)
			}
		}
		writer.Flush()
	}
	defer file.Close()

	// CSVファイル読み込みフェーズ
	file, err = os.Open(csvFileName)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // 1レコード当たりの予想フィールド数。-1の場合可変長、0の場合は最初のレコード
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// CSV内データ読み込みフェーズ
	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	// 読み込みデータ出力フェーズ
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
