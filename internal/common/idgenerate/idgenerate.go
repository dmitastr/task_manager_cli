package idgenerate

import (
	"fmt"
	"sync"

	"github.com/dmastr/task-manager-cli/internal/domain/entity"
)

type IdGenerator struct {
	sync.Mutex // ensures autoInc is goroutine-safe
	id         entity.TaskID
}

func NewIdGenerator(initialValue entity.TaskID) *IdGenerator {
	return &IdGenerator{id: initialValue}
}

func (g *IdGenerator) ID() (id entity.TaskID) {
	g.Lock()
	defer g.Unlock()

	id = g.id
	g.id++
	return
}


func (g *IdGenerator) IDString() string {
	id := g.id
	return fmt.Sprint(id)
}
