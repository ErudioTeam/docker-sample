package util

// import (
// 	//"encoding/json"
// 	"fmt"
// 	"github.com/spf13/viper"
// )

type Student struct {
	ID			string `mapstructure:"id"`
	Name    	string `mapstructure:"name"`
	FatherName 	string `mapstructure:"father_name"`
}

type StudentCollection struct {
	StudentList []Student `mapstructure:"studentlist"`
}

//var StudentCollection *Students

/* func (i Collection) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
} */

// func LoadCollection(path string) (err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("collection")
// 	viper.SetConfigType("json")

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		fmt.Println("Error in reading config ", err)
// 		return
// 	}

// 	err = viper.Unmarshal(&BookCollection)
// 	if err != nil {
// 		fmt.Println("Unable to decode into struct, %v", err)
// 	}

// 	return
// }

// func LoadCollection(path string) (err error) {
// 	// viper.AddConfigPath(path)
// 	// viper.SetConfigName("collection")
// 	// viper.SetConfigType("json")

// 	// err = viper.ReadInConfig()
// 	// if err != nil {
// 	// 	fmt.Println("Error in reading config ", err)
// 	// 	return
// 	// }

// 	// err = viper.Unmarshal(&StudentCollection)
// 	// if err != nil {
// 	// 	fmt.Println("Unable to decode into struct, %v", err)
// 	// }

	




// 	return
// }