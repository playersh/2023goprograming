package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("input score :")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace(inputScoreString)      // 공란 제거
	inputScore, err := strconv.ParseFloat(inputScoreString, 32) // casting

	if err != nil {
		log.Fatal(err)
	}

	if inputScore >= 90 {
		grade := "A grade!"
	} else {
		grade := "BCDE grade~"
	}

}
