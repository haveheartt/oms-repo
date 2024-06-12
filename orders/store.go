package main

import "context"

type store struct {
    //add here our mongoDB instance
}

func NewStore() *store {
    return &store{}
}

func (s *store) Create(context.Context) error {
    return nil
}
