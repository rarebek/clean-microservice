package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type BaseUseCase struct{}

func (u *BaseUseCase) Error(msg string, err error) error {
	if len(strings.TrimSpace(msg)) != 0 {
		return fmt.Errorf("%v: %w", msg, err)
	}
	return err
}

func (u *BaseUseCase) beforeRequest(guid *string, createdAt *string, updatedAt *string) {
	if guid != nil {
		*guid = uuid.New().String()
	}

	if *createdAt != "" {
		*createdAt = time.Now().UTC().String()
	}

	if *updatedAt != "" {
		*updatedAt = time.Now().UTC().String()
	}
}
