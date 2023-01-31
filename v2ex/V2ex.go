package v2ex

import "sp1der/models"

var (
	Title = models.Selector{
		Key:           "标题",
		SelectorQuery: "a.topic-link",
		Attr:          "",
		Text:          nil,
		AttrVal:       nil,
	}
	CommentCount = models.Selector{
		Key:           "评论数",
		SelectorQuery: "a.count_livid",
		Attr:          "",
		AttrVal:       nil,
		Text:          nil,
	}
	Author = models.Selector{
		Key:           "作者",
		SelectorQuery: "span.topic_info > strong:nth-child(3) > a",
		Attr:          "",
		AttrVal:       nil,
		Text:          nil,
	}
	Topic = models.Selector{
		Key:           "话题",
		SelectorQuery: "span.topic_info > a",
		Attr:          "",
		AttrVal:       nil,
		Text:          nil,
	}
)
