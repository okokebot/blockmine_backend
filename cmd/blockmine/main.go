package main

import (
	"fmt"

	"github.com/okokebot/blockmine/pkg/redmine"
)

func main() {
	c := redmine.NewClient()
	p := c.GetIssue(65527)
	c.GetChildrenInfo(p)
	s := p.createReleaseBlock(*c)
	fmt.Println(s)
}
