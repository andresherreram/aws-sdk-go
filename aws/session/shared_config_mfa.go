package session

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type promptTokenCLI struct{}

func (token promptTokenCLI) Token() *string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter MFA token: ")

	getToken, _ := reader.ReadString('\n')
	strToken := strings.TrimSpace(getToken)
	return &strToken

}
