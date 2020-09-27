package esquery

type NestedQuery struct {
	path           string
	query          Mappable
	scoreMode      string
	ignoreUnmapped bool
}

//Nested
func Nested(path string, query Mappable) *NestedQuery {
	return &NestedQuery{
		path:           path,
		query:          query,
		ignoreUnmapped: false,
	}
}

//ScoreMode
func (n *NestedQuery) ScoreMode(val string) *NestedQuery {
	n.scoreMode = val
	return n
}

//IgnoreUnmapped
func (n *NestedQuery) IgnoreUnmapped(val bool) *NestedQuery {
	n.ignoreUnmapped = val
	return n
}

func (n *NestedQuery) Map() map[string]interface{} {
	var query = make(map[string]interface{})
	var nq = make(map[string]interface{})
	query["nested"] = nq
	src := n.query.Map()
	nq["query"] = src
	nq["path"] = n.path
	if n.scoreMode != "" {
		nq["score_mode"] = n.scoreMode
	}

	if n.ignoreUnmapped != false {
		nq["ignore_unmapped"] = n.ignoreUnmapped
	}

	return query
}
