package hw10programoptimization

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := domainStat(r, domain)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return u, nil
}

func domainStat(r io.Reader, domain string) (DomainStat, error) {
	dec := json.NewDecoder(r)
	result := make(DomainStat)
	for {
		var user User
		if err := dec.Decode(&user); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("decode error: %w", err)
		}

		if strings.HasSuffix(user.Email, "."+domain) {
			lower := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[lower] = result[lower] + 1
		}
	}

	return result, nil
}

func GetDomainStatOld(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		var user User
		if err = json.Unmarshal([]byte(line), &user); err != nil {
			return
		}
		result[i] = user
	}
	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		matched, err := regexp.Match("\\."+domain, []byte(user.Email))
		if err != nil {
			return nil, err
		}

		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
