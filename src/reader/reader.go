package reader

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"unicode"
)

func ReadUsersFromFile() ([]string, error) {
	OpenFile, err := os.Open("users")
	if err != nil {
		return nil, errors.New("Error when Reading Current Valid Users")
	}

	defer OpenFile.Close()
	scanner := bufio.NewScanner(OpenFile)

	var rawUsers []string
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		rawUsers = append(rawUsers, line)
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Error when Reading Current Valid Users")
	}

	users := formantUsers(rawUsers)
	return users, nil
}

func formantUsers(lineArray []string) []string {
	var newUsers []string

	i := 0

	for _, lines := range lineArray {

		if lines == "" {
			continue
		}

		lines = strings.TrimSpace(lines)
		usersPerLine := strings.Fields(lines)

		for _, user := range usersPerLine {
			if isValidUser(user) {
				newUsers = append(newUsers, user)
				i++
			}
		}

	}
	return newUsers
}

func isValidUser(user string) bool {
	if len(user) != 8 {
		return false
	}

	for i := 0; i < 4; i++ {
		if !unicode.IsLower(rune(user[i])) {
			return false
		}
	}
	for i := 4; i < 8; i++ {
		if !unicode.IsDigit(rune(user[i])) {
			return false
		}
	}
	return true
}
