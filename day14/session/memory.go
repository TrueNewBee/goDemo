package session

import (
	"errors"
	"sync"
)

// 内存管理session

type MemorySession struct {
	sessionId string
	// 存 kv
	data map[string]interface{}
	rwlock sync.RWMutex
}

// 构造函数
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data : make(map[string]interface{}, 16),
	}
	return s
}

// Set
func (m *MemorySession)Set(key string , value interface{})(err error)  {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// 设置值
	m.data[key] = value
	return
}

// Get
func (m *MemorySession)Get(key string )( value interface{}, err error)  {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value , ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

// Delete
func (m *MemorySession)Delete(key string )( err error)  {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	 delete(m.data, key)
	return
}

// Save
func (m *MemorySession)Save(key string )( err error) {
	return
}
