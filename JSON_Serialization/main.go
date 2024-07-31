package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type MyTimedStruct struct {
	Name string 
	Timestamp time.Time
}

type MyStructCanMarshalRFC1123 struct {
	Name string 
	Timestamp time.Time
}

func (m *MyStructCanMarshalRFC1123) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	err := json.Unmarshal(data, &temp)

	if err != nil {
		return err
	}

	m.Name = temp["Name"].(string)

	timestamp_str := temp["Timestamp"].(string)

	t, err := time.Parse(time.RFC1123, timestamp_str)

	if err != nil {
		return err
	}

	m.Timestamp = t

	return nil
}

func (m MyStructCanMarshalRFC1123) MarshalJSON() ([]byte, error) {
	
	timestamp_str := m.Timestamp.Format(time.RFC1123)

	return json.Marshal(map[string]interface{}{"Name": m.Name, "Timestamp": timestamp_str})
}

func main(){

	//Json with a "Timestamp" ISO 8601 string
	//The string is in the format "2006-01-02T15:04:05Z07:00"
	//The "Z07:00" part is the time zone offset

	var jsonBlob [] byte = [] byte(`{"Timestamp":"2021-03-14T01:59:59-05:00"}`)

	//Unmarshal the JSON blob into a map

	var data map[string]interface{} = make(map[string]interface{})

	err := json.Unmarshal(jsonBlob, &data)

	if err != nil {
		panic(err)
	}

	//get the timestamp value from the map
	timestamp := data["Timestamp"].(string)

	//parse the timestamp string into time.Time assuming ISO 8601 format
	t, err := time.Parse(time.RFC3339, timestamp)

	if err != nil {
		panic(err)
	}

	//Print the parsed time
	fmt.Println(t)

	var myStruct MyTimedStruct

	//Unmarshal the JSON blob into a struct
	err = json.Unmarshal(jsonBlob, &myStruct)

	//test the struct
	fmt.Println(myStruct.Timestamp)

	//Print the struct
	fmt.Println(myStruct)

	// Print the time in RFCC3339 format
	fmt.Println(myStruct.Timestamp.Format(time.RFC3339))


	//What if you wrote JSON with a different serialization format

	var new_time time.Time = time.Date(2021, time.March, 14, 1, 59, 59, 0, time.FixedZone("PDT", -7*60*60))

	//new string using something not rfc3339
	formatted_new_time := new_time.Format(time.RFC1123)

	fmt.Println(formatted_new_time)

	// now create a new json string with the new time and Name
	new_json_blob := []byte(fmt.Sprintf(`{"Name":"MyName", "Timestamp":"%s"}`, formatted_new_time))

	//Unmarshal the JSON blob into a struct
	
	var new_struct MyTimedStruct

	err = json.Unmarshal(new_json_blob, &new_struct)

	if err != nil {
		//panic(err)
		fmt.Println(err)
	}


	//print the timestamp in time.RFC3339 format
	fmt.Println(new_struct.Timestamp.Format(time.RFC3339))


	var myStructCanMarshalRFC1123 MyStructCanMarshalRFC1123

	// ensure that the struct impliments the MarshalJSON and UnmarshalJSON interfaces
	
	//Fail out if the struct does not implement the MarshalJSON and UnmarshalJSON interfaces
	var _ json.Marshaler = &myStructCanMarshalRFC1123
	var _ json.Unmarshaler = &myStructCanMarshalRFC1123

	//Unmarshal the JSON blob into a struct
	err = json.Unmarshal(new_json_blob, &myStructCanMarshalRFC1123)

	if err != nil {
		panic(err)
	}

	//print the timestamp in time.RFC3339 format
	fmt.Println(myStructCanMarshalRFC1123.Timestamp.Format(time.RFC3339))
	//print timestamp out in native format
	fmt.Println(myStructCanMarshalRFC1123.Timestamp)
	
}



