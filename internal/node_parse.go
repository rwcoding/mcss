package internal

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
)

func Parse(text []byte) ([]*Node, error) {
	tokenizer := html.NewTokenizer(bytes.NewReader(text))
	queue := newQueue()

	for {
		tt := tokenizer.Next()

		if tt == html.CommentToken {
			continue
		}

		if tt == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return queue.nodes(), nil
			}
			return nil, tokenizer.Err()
		}

		token := tokenizer.Token()
		isVoidTag := isVoidTag(token.Data)

		if tt == html.EndTagToken {
			if !isVoidTag {
				queue.trace(token.Data)
			}
			continue
		}

		node := Node{
			NeedClose:  !isVoidTag,
			Attributes: map[string]string{},
		}
		queue.add(&node)

		if tt == html.TextToken {
			node.Type = TextType
			node.Children = nil
			node.Content = html.UnescapeString(token.String())
			continue
		}

		for _, v := range token.Attr {
			node.Attributes[v.Key] = v.Val
		}

		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			node.Type = TagType
			node.Tag = token.Data
		} else if tt == html.DoctypeToken {
			node.Type = DocType
			node.NeedClose = false
			node.Content = token.String()
		}
	}

	return queue.nodes(), nil
}
