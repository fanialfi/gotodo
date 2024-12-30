package lib

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/fanialfi/gotodo/internal/task"
)

var headersTable = []string{
	lipgloss.NewStyle().Bold(true).Render("ID"),
	lipgloss.NewStyle().Bold(true).Render("STATUS"),
	lipgloss.NewStyle().Bold(true).Render("DESCRIPTION"),
	lipgloss.NewStyle().Bold(true).Render("CREATED"),
}

var tableLayout = table.New().
	Border(lipgloss.HiddenBorder()).
	// BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("128"))).
	StyleFunc(func(row, col int) lipgloss.Style {
		return lipgloss.NewStyle().
			Margin(0, 0, 1, 0).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("128")).
			Align(lipgloss.Center)
	}).
	Headers(headersTable...)

var headersStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF000")).Bold(true)

func PrintingOutput(tasks *[]task.Task, header string) {
	if len(*tasks) != 0 {
		fmt.Println(headersStyle.Render(header))
		for _, task := range *tasks {
			taskIdString := strconv.Itoa(int(task.ID))
			taskstatus := string(task.Status)
			taskDescription := string(task.Description)
			creationTimeString := time.UnixMilli(task.CreatedAt).Format("Monday, 02-Jan-2006 15:00:4")
			tableLayout.Rows([]string{taskIdString, taskstatus, taskDescription, creationTimeString})
		}
		fmt.Println(tableLayout)
	} else {
		fmt.Println(headersStyle.Render("Task Not Found"))
	}
}
