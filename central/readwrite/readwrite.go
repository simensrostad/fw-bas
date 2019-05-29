package readwrite

import  (
    "encoding/json"
    "io/ioutil"
     "fmt"
    ."../config"
    "os"
)

func Write_to_file(elev PATIENT){
  elevator_json, _ := json.MarshalIndent(elev, "","")
  ioutil.WriteFile("UUIDdir.json", elevator_json, 0644)
}

func Read_from_file(UUID_check string) (PATIENT){
    output := PATIENT{}
    jsonFile, err := os.Open("UUIDdir.json")
        if err != nil {
            fmt.Println(err)
        }

        raw_data, _ := ioutil.ReadAll(jsonFile)
        elev := Users{}

        json.Unmarshal(raw_data, &elev)

        for i := 0; i < len(elev.Users); i++ {
          if elev.Users[i].UUID == UUID_check {
              return elev.Users[i]
          }
        }
    return output
}
