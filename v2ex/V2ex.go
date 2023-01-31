package v2ex

import "sp1der/models"

var (
	Title = models.Selector{
		Key:           "标题",
		SelectorQuery: "a.topic-link",
	}
	CommentCount = models.Selector{
		Key:           "评论数",
		SelectorQuery: "a.count_livid",
	}
	Author = models.Selector{
		Key:           "作者",
		SelectorQuery: "span.topic_info > strong:nth-child(3) > a",
	}
	Topic = models.Selector{
		Key:           "话题",
		SelectorQuery: "span.topic_info > a",
	}
	Link = models.Selector{
		Key:           "链接",
		SelectorQuery: "a.topic-link",
		Attr:          "href",
		AttrPrefix:    "https://www.v2ex.com",
	}
)
