package search

// 评分规则通用接口
type SearchScorer interface {
	// 给一个文档评分，文档排序时先用第一个分值比较，如果
	// 分值相同则转移到第二个分值，以此类推。
	// 返回空切片表明该文档应该从最终排序结果中剔除。
	Score(doc IndexedDocument, fields interface{}) []float32
}

// 默认值见engine_init_options.go
type BM25Parameters struct {
	K1 float32
	B  float32
}
