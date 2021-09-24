package hds

type queue struct {
	data []*Node
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) add(node *Node) {
	q.data = append(q.data, node)
}

func (q *queue) trace(tag string) {
	length := len(q.data)
	for i := length-1; i >= 0; i-- {
		if q.data[i].Tag == tag {
			if i >= length - 1 {
				break
			}
			for _,v := range q.data[i+1:] {
				q.data[i].Add(v)
			}
			q.data = q.data[:i+1]
			break
		}
	}
}

func (q *queue) nodes() []*Node {
	return q.data
}
