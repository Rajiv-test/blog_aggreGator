package main

import (
	"context"
	"fmt"
)


func handlerReset(s *state,cmd command) error{
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset the users %w",err)
	}
	return nil
}