package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type (
	cards struct {
		reader         *bufio.Reader
		TermToDef      map[string]string `json:"termToDef"`
		defToTerm      map[string]string
		errors         map[string]int
		highestCounter int
		log            []string
	}
)

func main() {
	c := initCardsStruct()
	expPath := c.handleFlags()
	c.menu(expPath)
}

func initCardsStruct() *cards {
	return &cards{
		reader:    bufio.NewReader(os.Stdin),
		TermToDef: make(map[string]string),
		defToTerm: make(map[string]string),
		errors:    make(map[string]int),
	}
}

func (c *cards) handleFlags() string {
	var imp, exp string

	flag.StringVar(&imp, "import_from", "", "")
	flag.StringVar(&exp, "export_to", "", "")
	flag.Parse()

	if imp != "" {
		c.importCards(imp)
	}

	return exp
}

func (c *cards) menu(expPath string) {
	for {
		msg := "Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):"
		c.writeLogEntry(msg, true)

		line, _ := c.reader.ReadString('\n')
		line = strings.ToLower(strings.TrimSpace(line))
		c.writeLogEntry(line, false)

		switch line {
		case "add":
			c.addCard()
		case "remove":
			c.removeCard()
		case "import":
			c.importCards("")
		case "export":
			c.exportCards("")
		case "ask":
			c.ask()
		case "log":
			c.saveLog()
		case "hardest card":
			c.printHardestCards()
		case "reset stats":
			c.resetStats()
		case "exit":
			msg = "Bye bye!"
			if expPath != "" {
				c.exportCards(expPath)
			}
			c.writeLogEntry(msg, true)

			return
		default:
		}
	}
}

func (c *cards) addCard() {
	msg := "Input the number of cards:"
	c.writeLogEntry(msg, true)

	term := c.getTerm()
	def := c.getDefinition()

	c.TermToDef[term] = def
	c.defToTerm[def] = term

	msg = fmt.Sprintf("The pair (\"%s\":\"%s\") has been added.", term, def)
	c.writeLogEntry(msg, true)
}

func (c *cards) getTerm() string {
	msg := "The card:"
	c.writeLogEntry(msg, true)

	for {
		line, _ := c.reader.ReadString('\n')
		line = strings.TrimSpace(line)
		c.writeLogEntry(line, false)

		if _, ok := c.TermToDef[line]; !ok {
			return line
		} else {
			msg = fmt.Sprintf("The term \"%s\" already exists. Try again:", line)
			c.writeLogEntry(msg, true)
		}
	}
}

func (c *cards) getDefinition() string {
	msg := "The definition of the card:"
	c.writeLogEntry(msg, true)

	for {
		line, _ := c.reader.ReadString('\n')
		line = strings.TrimSpace(line)
		c.writeLogEntry(line, false)

		if _, ok := c.defToTerm[line]; !ok {
			return line
		} else {
			msg = fmt.Sprintf("The definition \"%s\" already exists. Try again:", line)
			c.writeLogEntry(msg, true)
		}
	}
}

func (c *cards) removeCard() {
	msg := "Which card?"
	c.writeLogEntry(msg, true)

	line, _ := c.reader.ReadString('\n')
	line = strings.TrimSpace(line)
	c.writeLogEntry(line, false)

	def, ok := c.TermToDef[line]
	if ok {
		delete(c.TermToDef, line)
		delete(c.defToTerm, def)

		msg = "The card has been removed."
		c.writeLogEntry(msg, true)
	} else {
		msg = fmt.Sprintf("Can't remove \"%s\": there is no such card.", line)
		c.writeLogEntry(msg, true)
	}

}

func (c *cards) importCards(path string) {
	var line, msg string

	if path != "" {
		line = path
	} else {
		msg = "File name:"
		c.writeLogEntry(msg, true)

		line, _ = c.reader.ReadString('\n')
		line = strings.TrimSpace(line)
		c.writeLogEntry(line, false)
	}

	file, err := os.Open(line)
	if err != nil {
		msg = "File not found."
		c.writeLogEntry(msg, true)

		return
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	fileCards := make(map[string]string)

	err = json.Unmarshal(byteValue, &fileCards)
	if err != nil {
		msg = fmt.Sprintf("Error at Unmarshal: %s", err.Error())
		c.writeLogEntry(msg, true)

		return
	}

	for term, def := range fileCards {
		c.TermToDef[term] = def
		c.defToTerm[def] = term
	}

	msg = fmt.Sprintf("%d cards have been loaded.", len(fileCards))
	c.writeLogEntry(msg, true)
}

func (c *cards) exportCards(expPath string) {
	var line, msg string
	if expPath != "" {
		line = expPath
	} else {
		msg = "File name:"
		c.writeLogEntry(msg, true)

		line, _ = c.reader.ReadString('\n')
		line = strings.TrimSpace(line)
		c.writeLogEntry(line, false)
	}

	cardsJSON, err := json.MarshalIndent(c.TermToDef, "", "  ")
	if err != nil {
		msg = fmt.Sprintf("Error at MarshalIndent: %s", err.Error())
		c.writeLogEntry(msg, true)

		return
	}

	err = os.WriteFile(line, cardsJSON, 0644)
	if err != nil {
		msg = fmt.Sprintf("Error at WriteFile: %s", err.Error())
		c.writeLogEntry(msg, true)

		return
	}

	msg = fmt.Sprintf("%d cards have been saved.", len(c.TermToDef))
	c.writeLogEntry(msg, true)
}

func (c *cards) ask() {
	msg := "How many times to ask?"
	c.writeLogEntry(msg, true)

	line, _ := c.reader.ReadString('\n')
	line = strings.TrimSpace(line)
	c.writeLogEntry(line, false)

	num, _ := strconv.Atoi(line)

	for i := 0; i < num; i++ {
		term, def := pickRandomEntry(c.TermToDef)

		userDef := c.getAnswer(term)
		if userDef == def {
			msg = "Correct!"
			c.writeLogEntry(msg, true)
		} else {
			c.errors[term] += 1
			if c.errors[term] > c.highestCounter {
				c.highestCounter = c.errors[term]
			}

			cDef, ok := c.defToTerm[userDef]
			if ok {
				msg = fmt.Sprintf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".", def, cDef)
				c.writeLogEntry(msg, true)
			} else {
				msg = fmt.Sprintf("Wrong. The right answer is \"%s\".", def)
				c.writeLogEntry(msg, true)
			}
		}
	}
}

func pickRandomEntry(m map[string]string) (string, string) {
	i := rand.Intn(len(m))
	for term, def := range m {
		if i == 0 {
			return term, def
		}
		i--
	}
	return "", ""
}

func (c *cards) getAnswer(term string) string {
	msg := fmt.Sprintf("Print the definition of \"%s\":", term)
	c.writeLogEntry(msg, true)

	line, _ := c.reader.ReadString('\n')
	line = strings.TrimSpace(line)
	c.writeLogEntry(line, false)

	return line
}

func (c *cards) writeLogEntry(str string, doPrint bool) {
	if doPrint {
		fmt.Println(str)
	}

	c.log = append(c.log, str)
}

func (c *cards) saveLog() {
	msg := "File name:"
	c.writeLogEntry(msg, true)

	var log string
	for _, entry := range c.log {
		log += entry + "\n"
	}

	line, _ := c.reader.ReadString('\n')
	line = strings.TrimSpace(line)
	c.writeLogEntry(line, false)

	err := os.WriteFile(line, []byte(log), 0644)
	if err != nil {
		msg = fmt.Sprintf("Error at WriteFile: %s", err.Error())
		c.writeLogEntry(msg, true)

		return
	}

	msg = "The log has been saved."
	c.writeLogEntry(msg, true)
}

func (c *cards) getHardestCards() map[int][]string {
	cardsMap := make(map[int][]string)
	for term, counter := range c.errors {
		cardsMap[counter] = append(cardsMap[counter], term)
	}

	return cardsMap
}

func (c *cards) printHardestCards() {
	if len(c.errors) == 0 {
		msg := "There are no cards with errors."
		c.writeLogEntry(msg, true)
	} else {
		cardsMap := c.getHardestCards()

		msg := fmt.Sprintf("The hardest card %s \"%s\". You have %d errors answering it.",
			func() string {
				if len(cardsMap[c.highestCounter]) > 1 {
					return "are"
				} else {
					return "is"
				}
			}(),
			strings.Join(cardsMap[c.highestCounter], ", "), c.highestCounter)

		c.writeLogEntry(msg, true)
	}
}

func (c *cards) resetStats() {
	c.errors = make(map[string]int)
	c.highestCounter = 0

	msg := "Card statistics have been reset."
	c.writeLogEntry(msg, true)
}
