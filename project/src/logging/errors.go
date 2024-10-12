package logging

import (
	"errors"
	"fmt"
	"project/src/flags"
	"runtime"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const reset = "\033[0m"

const (
	bgGreen   = "\033[97;42m"
	bgWhite   = "\033[90;47m"
	bgYellow  = "\033[90;43m"
	bgRed     = "\033[97;41m"
	bgBlue    = "\033[97;44m"
	bgMagenta = "\033[97;45m"
	bgCyan    = "\033[97;46m"
)
const (
	green   = "\033[32m"
	white   = "\033[97m"
	yellow  = "\033[33m"
	red     = "\033[31m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
)

func GetLinePrefix(timestamp string, i int) string {
	return Magenta(fmt.Sprintf("%s %d. ", timestamp, i))
}

func Bubble(err error, msg string) error {
	if err == nil || !flags.UseLogger {
		return errors.New(msg)
	}
	return fmt.Errorf("%s\n%w", msg, err)
}

func GreenBg(msg string) string {
	return WrapInColour(msg, bgGreen)
}

func WhiteBg(msg string) string {
	return WrapInColour(msg, bgWhite)
}

func YellowBg(msg string) string {
	return WrapInColour(msg, bgYellow)
}

func RedBg(msg string) string {
	return WrapInColour(msg, bgRed)
}

func BlueBg(msg string) string {
	return WrapInColour(msg, bgBlue)
}

func MagentaBg(msg string) string {
	return WrapInColour(msg, bgMagenta)
}

func CyanBg(msg string) string {
	return WrapInColour(msg, bgCyan)
}

func Green(msg string) string {
	return WrapInColour(msg, green)
}

func White(msg string) string {
	return WrapInColour(msg, white)
}

func Yellow(msg string) string {
	return WrapInColour(msg, yellow)
}

func Red(msg string) string {
	return WrapInColour(msg, red)
}

func Blue(msg string) string {
	return WrapInColour(msg, blue)
}

func Magenta(msg string) string {
	return WrapInColour(msg, magenta)
}

func Cyan(msg string) string {
	return WrapInColour(msg, cyan)
}

func LogError(err error, msg string) {
	if !flags.UseLogger {
		return
	}
	Log(Bubble(err, Red(msg)).Error())
}

func SampleColours() {
	fmt.Println(GreenBg("Green"))
	fmt.Println(WhiteBg("White"))
	fmt.Println(YellowBg("Yellow"))
	fmt.Println(RedBg("Red"))
	fmt.Println(BlueBg("Blue"))
	fmt.Println(MagentaBg("Magenta"))
	fmt.Println(CyanBg("Cyan"))
	fmt.Println(Green("Green"))
	fmt.Println(White("White"))
	fmt.Println(Yellow("Yellow"))
	fmt.Println(Red("Red"))
	fmt.Println(Blue("Blue"))
	fmt.Println(Magenta("Magenta"))
	fmt.Println(Cyan("Cyan"))
}

func Log(msg string) {
	if !flags.UseLogger {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	numLines := strings.Count(msg, "\n")
	for i := 0; i < numLines; i++ {
		prefix := GetLinePrefix(timestamp, i+1)
		msg = strings.Replace(msg, "\n", fmt.Sprintf("\n%s", prefix), 1)

		if i == numLines-1 {
			msg += "\n"
		}
	}
	fmt.Printf(GetLinePrefix(timestamp, 0) + msg + "\n")
}

func createStackMessage(msg string) string {
	// Get the callers in the stack up.
	callers := make([]uintptr, 100)
	numCallers := runtime.Callers(2, callers)
	callers = callers[:numCallers]
	frames := runtime.CallersFrames(callers)

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	lines := []string{}
	for i := numCallers - 1; i >= 0; i-- {
		frame, _ := frames.Next()
		file := frame.File
		line := frame.Line
		funcName := frame.Function

		prefix := GetLinePrefix(timestamp, i+1)
		currentLine := fmt.Sprintf("%s %s:%d %s", prefix, file, line, funcName)

		lines = append(lines, currentLine)
	}
	slices.Reverse(lines)
	output := strings.Join(lines, "\n")

	msgLines := strings.Split(msg, "\n")
	numLines := len(msgLines)
	for i := 0; i < numLines; i++ {
		msgLines[i] = GetLinePrefix(timestamp, numCallers+numLines-i) + msgLines[i]
	}
	slices.Reverse(msgLines)
	msg = strings.Join(msgLines, "\n") + "\n"

	if output != "" {
		output += "\n"
	}

	return output + msg
}

func LogErrorWithStack(err error, msg string) {
	if !flags.UseLogger {
		return
	}
	fmt.Print(createStackMessage(Bubble(err, Red(msg)).Error()))
}

func LogWithStack(msg string) {
	if !flags.UseLogger {
		return
	}
	fmt.Print(createStackMessage(msg))
}

func LogErrorToGin(err error, msg string) {
	if !flags.UseLogger {
		return
	}
	LogToGin(Bubble(err, msg).Error())
}

func LogToGin(msg string) {
	if !flags.UseLogger {
		return
	}
	fmt.Fprint(gin.DefaultWriter, createStackMessage(msg))
}

func GinError(ginContext *gin.Context, err error, msg string) {
	bubbled := Bubble(err, Red(msg))
	LogToGin(bubbled.Error())
	SendNotFound(ginContext, err, msg)
}

func WrapInColour(msg string, colour string) string {
	return fmt.Sprintf("%s%s%s", colour, msg, reset)
}

func SendNotFound(ginContext *gin.Context, err error, msg string) error {
	return ginContext.AbortWithError(404, Bubble(err, Red(msg)))
}

func LogWithRedirect(ginContext *gin.Context, err error, msg string) {
	LogError(err, msg)
	ginContext.Redirect(302, "/")
}
