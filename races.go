package main 

import (
"encoding/json"
)


type Race {
 Name string `json="name"`
 City string `json="city"`
 Link string `json="link"`
 Departement int `json="dep"`
 Site string `json="site"`
 
}


func (r *Race) IsComplete() bool {
 if r.Name != "" && r.City != "" && r.Link != "" && r.Departement != 0 && r.Site != "" {
 return true

} 
 return false

}


func RacesToJson(dist *os.File, races []Race){
   
   
 JsonByte, err := json.Marshall(races)
 dist.WriteBytes(jsonByte)
 defer dist.Close()


}