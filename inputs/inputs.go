package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chijiokejoseph/gowayback-beats/consts"
)

// InputDate prompts the user to enter a date format
// and returns the parsed date and an error.
//
// Returns:
//   - time.Time: parsed date
//   - error: returns an error if the parsing fails or nil if
//     otherwise.
func InputDate() (time.Time, error) {
	prompt := "What year do you wish to travel to? Type the input in this format YYYY-MM-DD: "
	// prompt user
	response, err := Input(prompt)
	// if there is an error during user input return
	if err != nil {
		return time.Now(), nil
	}
	return time.Parse(consts.DATE_FORMAT, response)
}

// Input prompts a user with a given `prompt` message and returns the
// user response trimmed and an error for cases where reading
// user input fails
//
// Arguments:
//   - prompt: a message to display to the user
//
// Returns:
//   - string: user response after trimming
//   - error: an error if reading user input fails or nil if otherwise
func Input(prompt string) (string, error) {
	// display prompt
	fmt.Print(prompt)

	// flush stdout to ensure that  the prompt
	// is displayed before reading stdin
	err := bufio.NewWriter(os.Stdout).Flush()
	if err != nil {
		return "", err
	}

	var response string

	// read standard input
	_, err = fmt.Scanln(&response)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(prompt), nil
}
