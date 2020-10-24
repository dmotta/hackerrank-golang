package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"

)

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int32) int32 {
    var m map[int32]int
    m = make(map[int32]int)
    for i := 0; i<len(arr); i++{
        x, ok := m[arr[i]]
        if ok == true {
                x = x+1
                m[arr[i]]=x
        }else{
            m[arr[i]]=1
        }
    }
    var mostCommonBirdKey int32 = 0
    var mostCommonBirdVal int = 0
    for key, value := range m { 
        fmt.Println("Key:", key, "Value:", value)                
        if value > mostCommonBirdVal {

            mostCommonBirdVal = value
            mostCommonBirdKey = key
        } else if value == mostCommonBirdVal {
            if key < mostCommonBirdKey{
                mostCommonBirdVal = value
                mostCommonBirdKey = key
            }
        }
    }   

    fmt.Println("mostCommonBirdVal:", mostCommonBirdVal, "mostCommonBirdKey: ",mostCommonBirdKey)        

    return int32(mostCommonBirdKey)

}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(arrCount); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := migratoryBirds(arr)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
