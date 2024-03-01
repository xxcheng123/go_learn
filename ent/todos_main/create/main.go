package main

import (
	"context"
	"encoding/json"
	"ent-demo/db"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := db.Get()
	for i := 0; i < 10; i++ {
		s, err := req()
		if err != nil {
			continue
		}
		client.Todo.Create()
		result, err := client.Todo.Create().
			SetTitle(s.Title).
			SetText(fmt.Sprintf("%s_%s", s.From, s.Author)).
			Save(context.Background())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}

}

type Sentence struct {
	Title  string `json:"hitokoto"`
	From   string `json:"from"`
	Author string `json:"from_who"`
}

func req() (*Sentence, error) {
	resp, err := http.Get("https://v1.hitokoto.cn/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var s *Sentence
	if err = json.Unmarshal(bs, &s); err != nil {
		return nil, err
	}
	return s, nil
}
