package daybookr

type TagSet struct {
	list map[string]struct{} // empty structs occupy 0 memory
}

func (set *TagSet) Has(t string) bool {
	_, ok := set.list[t]
	return ok
}

func (set *TagSet) Add(t string) {
	set.list[t] = struct{}{}
}

func (set *TagSet) Remove(t string) {
	delete(set.list, t)
}

func (set *TagSet) Clear() {
	set.list = make(map[string]struct{})
}

func (set *TagSet) Size() int {
	return len(set.list)
}

func (set *TagSet) Iterate() (tags []string) {
	for t := range set.list {
		tags = append(tags, t)
	}
	return tags
}

func NewTagSet() *TagSet {
	set := &TagSet{}
	set.list = make(map[string]struct{})
	return set
}
