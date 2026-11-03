package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func strSelect(arr []string) string {
	length := len(arr)
	index := rand.Intn(length)
	item := arr[index]
	return item
}
func numGen(min int, max int) int {
	num := rand.Intn(max-min+1) + min
	return num
}

func genMsg(msg string) string {
	msgStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#0a0a0a")).
		Background(lipgloss.Color("#fa6000")).
		Padding(0, 1).
		Bold(true).
		MarginLeft(1).
		Align(lipgloss.Center)
	message := msgStyle.Render(msg)
	return message
}
func main() {
	// Define a stylish header box
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#f5a71f")).
		Padding(1, 3).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#ffa600")).
		Align(lipgloss.Center)

	// Define an alert/status style
	statusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fffffd")).
		Background(lipgloss.Color("#9e5007")).
		Padding(0, 1).
		Bold(true).
		MarginLeft(1).
		Align(lipgloss.Center).
		Italic(true)
	// Render the components
	title := titleStyle.Render("Random Guy/Girl TikTok Trend")

	// Print everything to the terminal
	fmt.Println(title)

	genders := []string{"Male", "Female"}
	jobs := []string{"Mechanic", "Electronic", "Programmer", "Hacker", "Makeup Artist", "Artist", "Actor", "Sugar Daddy", "Sugar Mommy"}
	gender := statusStyle.Render(strSelect(genders))
	job := statusStyle.Render(strSelect(jobs))
	age := statusStyle.Render(strconv.Itoa(numGen(10, 96)))
	height := statusStyle.Render(strconv.Itoa(numGen(90, 250)) + "cm")
	fmt.Println(genMsg("Gender"), gender)
	fmt.Print("\n")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(genMsg("Age"), age)
	fmt.Print("\n")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(genMsg("Height"), height)
	fmt.Print("\n")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(genMsg("Job"), job)

}
