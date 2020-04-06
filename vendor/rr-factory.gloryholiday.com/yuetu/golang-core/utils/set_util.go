package utils

type StringSet interface {
	Add(item string) bool
	AddAll(items... string) bool
	Contain(item string) bool
	ContainAll(items... string) bool
	Remove(item string) bool
	RemoveAll(items ...string) bool
	GetValues() []string
	GetOrgValue() map[string]bool
	Size() int64
	ClearAll()
}

type stringSet struct {
	m map[string]bool
}

func NewStringSet() StringSet {
	return &stringSet{
		m: map[string]bool{},
	}
}

func (s *stringSet) Add(item string) bool {
	if len(item) > 0 {
		s.m[item] = true
		return true
	}
	return false
}

func (s *stringSet) AddAll(items... string) bool {
	for _, item := range items {
		r := s.Add(item)
		if !r {
			return false
		}
	}
	return true
}

func (s *stringSet) Contain(item string) bool {
	_, ok := s.m[item]
	return ok
}

func (s *stringSet) ContainAll(items... string) bool {
	for _, item := range items {
		if !s.Contain(item) {
			return false
		}
	}
	return true
}

func (s *stringSet) Remove(item string) bool {
	if len(item) == 0 || !s.Contain(item) {
		return false
	}
	delete(s.m, item)
	return true
}

func (s *stringSet) RemoveAll(items... string) bool {
	if 0 == len(items) {
		return false
	}
	for _, item := range items {
		s.Remove(item)
	}
	return true
}

func (s *stringSet) GetValues() []string {
	var r []string
	for key := range s.m {
		r = append(r, key)
	}
	return r
}

func (s *stringSet) GetOrgValue() map[string]bool {
	return s.m
}

func (s *stringSet) Size() int64 {
	return int64(len(s.m))
}


func (s *stringSet) ClearAll() {
	s.m = map[string]bool{}
}


