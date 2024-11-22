package tokenset

import (
	"log"
	"sync"

	"github.com/gorilla/sessions"
)

var (
	tokenSet = make(map[interface{}]struct{})
	mu       sync.Mutex
)

func Add(session *sessions.Session) {
	mu.Lock()
	defer mu.Unlock()
	tokenSet[session] = struct{}{}
	log.Printf("Add token, len: ", len(tokenSet))
}

func Delete(session *sessions.Session) {
	mu.Lock()
	defer mu.Unlock()
	delete(tokenSet, session)
	log.Printf("Delete token, len: ", len(tokenSet))
}

func IsTokenRevoked(token interface{}) bool {
	mu.Lock()
	defer mu.Unlock()
	_, ok := tokenSet[token]
	log.Printf("IsTokenRevoked: ", ok, len(tokenSet))
	return ok
}
