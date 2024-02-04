package ID

import (
	"fmt"
	"github.com/google/uuid"
)

func ReturnID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
	return u.String()
}
