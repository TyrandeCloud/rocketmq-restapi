package rocketmq_restapi

type Request struct {
	c         *Client
	namespace string
	tag       string
}

func NewRequest(c *Client, namespace string) *Request {
	return &Request{
		c:         c,
		namespace: namespace,
	}
}

func (r *Request) Tag(tag string) *Request {
	r.tag = tag
	return r
}
