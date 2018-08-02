package function

import (
	"fmt"
	"github.com/s8sg/faaschain"
)

// Handle a serverless request to chian
func Define(chain *faaschain.Fchain) (err error) {
	chain.ApplyModifier(func(data []byte) ([]byte, error) {
		return []byte(fmt.Sprintf("%s", string(data))), nil
	}).ApplyModifier(func(data []byte) ([]byte, error) {
		return []byte(fmt.Sprintf("you said %s", string(data))), nil
	})
	return
}
