package main

type GitInit struct {
	Quiet bool `q desc:"be quiet"`
}

func (GitInit) FlaglyDesc() string {
	return "Clone a repository into a new directory"
}
