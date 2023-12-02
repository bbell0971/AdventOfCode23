package main
import (
	"bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

func main() {
    // open file
    f, err := os.Open("day1Input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    
    var calibrationValues = buildValues(f)
    var finalValue = addValuesOfSlice(calibrationValues)
    
    fmt.Println(finalValue)
    
}

func buildValues(f *os.File)[]int{
    scanner := bufio.NewScanner(f)
    var calibrationValues []int
    for scanner.Scan() {
        re:=regexp.MustCompile("[0-9]")
        integersInText := re.FindAllString(scanner.Text(), -1)
        lineValue := integersInText[0] + integersInText[len(integersInText)-1]
        intValueFromLine, lineErr := strconv.Atoi(lineValue)
        if lineErr != nil {
            log.Fatal(lineErr)
        }
        calibrationValues = append(calibrationValues, intValueFromLine )
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return calibrationValues
}

func addValuesOfSlice(calibrationValues []int) int{
    var finalValue = 0
    for _, v := range calibrationValues {
        finalValue += v
        fmt.Println(finalValue)
    }
    return finalValue
}