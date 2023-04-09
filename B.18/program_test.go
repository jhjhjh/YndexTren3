package main
import( 
    "testing"
    "os"
    "bytes"
  // "fmt"
   "io"
    "io/ioutil"
    "strings"
)

func TestDo(t *testing.T){
    file, err := os.Open("03")
    if err != nil{
	panic(err)
    }
    fContent, err := ioutil.ReadFile("03.out")
    if err !=nil{
	panic(err)
    }
    result := bytes.NewBufferString("")
    Do(file,result)
    buf := new(strings.Builder)
    _, err = io.Copy(buf, result)
//    fmt.Println(string(fContent))
//    fmt.Println(buf)
    //str = new String(fContent.array(), "ASCII")
    //res  = new String(result.array(),"ASCII")
    if buf.String() != string(fContent) {
	t.Errorf("Result was incorrect, got: %s, want: %s.", buf.String(),string(fContent))
    }
    
}