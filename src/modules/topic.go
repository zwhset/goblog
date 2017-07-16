package modules

type Topic struct {
	Tid     int
	Title   string
	Author  string
	Content string
	Created string
}

func (this *Topic) ViewAll() []*Topic {
	ts := []*Topic{
		{
			Tid:     1,
			Title:   "这是第一篇文章",
			Author:  "zwhset",
			Content: "随便写点",
			Created: "2017-07-15",
		},
		{
			Tid:     2,
			Title:   "这是第二篇文章",
			Author:  "cntea",
			Content: "随便写点",
			Created: "2017-07-14",
		},
	}
	return ts
}
