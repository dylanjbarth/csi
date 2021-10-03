package qe

type PlanNode interface {
	Next() string
}

type ScanNode struct {
	i    int
	l    int
	data []string
}

func NewScanNode(data []string) *ScanNode {
	return &ScanNode{0, len(data), data}
}

func (s *ScanNode) Next() string {
	if s.i < s.l {
		buff := s.data[s.i]
		s.i += 1
		return buff
	}
	return ""
}

type LimitNode struct {
	i     int
	l     int
	input PlanNode
}

func NewLimitNode(limit int, input PlanNode) *LimitNode {
	return &LimitNode{0, limit, input}
}

func (l *LimitNode) Next() string {
	if l.i < l.l {
		l.i += 1
		return l.input.Next()
	}
	return ""
}
