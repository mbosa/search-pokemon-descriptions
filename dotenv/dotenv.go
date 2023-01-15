package dotenv

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

const defaultEnvFilename = ".env"

var envLineRegex = regexp.MustCompile(`^(?P<key>[a-zA-Z0-9_]+)=(?:(?:"(?P<valueInQuotes>(?:[^"\\]|\\.)*)")|(?P<value>[a-zA-Z0-9_\-.:/@+*]+))`)

type envMapT map[string]string

func Load(filename string) error {
	if filename == "" {
		filename = defaultEnvFilename
	}

	file, err := os.Open(defaultEnvFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	envMap, err := parseFile(file)
	if err != nil {
		return err
	}

	if err := setEnvVars(envMap); err != nil {
		return err
	}

	return nil
}

func parseFile(file io.Reader) (envMapT, error) {
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	envMap := make(envMapT)
	for _, line := range lines {
		key, value := parseLine(line)
		if key == "" {
			continue
		}
		envMap[key] = value
	}

	return envMap, nil
}

func parseLine(line string) (key string, value string) {
	subMatches := envLineRegex.FindStringSubmatch(line)

	if len(subMatches) != 4 {
		return "", ""
	}

	/*
		sub matches:
		0: whole match
		1: key
		2: valueInQuotes
		3: value
	*/
	key = subMatches[1]
	if subMatches[2] != "" {
		return key, subMatches[2]
	} else {
		return key, subMatches[3]
	}
}

func setEnvVars(envMap envMapT) error {
	for key, value := range envMap {

		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
