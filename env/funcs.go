package env

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"

	"github.com/joho/godotenv"

	"github.com/chijiokejoseph/gowayback-beats/utils"

)

type Env struct {
	ClientID     string `env:"CLIENT_ID,required"`
	ClientSecret string `env:"CLIENT_SECRET,required"`
	RedirectUri  string `env:"REDIRECT_URI" json:"redirect_url"`
}

// BuildEnv reads the OS environment data and loads 
// the corresponding fields to an Env struct using a 
// series of field validations for better error rendering.
// 
// Returns:
// 	- *Env: a pointer to the `Env` struct if successful, nil
// 		otherwise
// 	- error: an error if the operation fails, nil otherwise
func BuildEnv() (*Env, error) {
	// instantiate Env struct
	newEnv := Env{}

	// get reflect.Type of Env struct
	envType := reflect.TypeOf(newEnv)

	// get number of struct fields
	numFields := envType.NumField()
	// get environment values
	envVars := make([]any, numFields)

	// repeat for each field of struct
	for i := range numFields {
		// get struct field
		envField := envType.Field(i)

		// get field tag
		fieldTag := envField.Tag

		// get env var name in field tag
		varname := fieldTag.Get("env")

		// if value is empty fill with nil
		if varname == "" {
			envVars[i] = nil
			continue
		}

		// if variable name doesn't contain comma i.e., it is plain
		if !strings.Contains(varname, ",") {
			// lookup environment variable name in the environment
			envVar, ok := os.LookupEnv(varname)
			// if variable name not found, fill `envVars` with empty(nil)
			if !ok {
				envVars[i] = nil
				continue
			}
			// else fill the `envVars` slice with variable's value
			envVars[i] = envVar
		}

		// if variable name contains comma i.e., it has a modifier
		// partition string across the commas
		strPartitions := strings.Split(varname, ",")

		// get variable name from the first part of the partition
		varname = strPartitions[0]

		// the rest of the partition slice contains the modifiers
		modifiers := strPartitions[1:len(strPartitions):len(strPartitions)]

		if slices.Contains(modifiers, "required") {
			// if one of the modifiers is "required"
			// look up variable name in the environment
			envVar, ok := os.LookupEnv(varname)

			// if variable not found, return an error since it is required
			if !ok {
				msg := fmt.Sprintf("Struct Field %s of Env has tag %s with the required attribute. \nBut the corresponding environment variable for %s was not found.", envField.Name, fieldTag, varname)
				return &newEnv, errors.New(msg)
			}

			// if found, append the value of the variable to the `envVars` slice
			envVars[i] = envVar
		}
	}

	// for each environment value
	for i, envVar := range envVars {
		// extract corresponding struct field
		field := envType.Field(i)

		// match against expected field names
		switch field.Name {
		case "ClientID":
			value, ok := envVar.(string)
			// if environment value is not a string skip current field
			if !ok {
				continue
			}
			// else set field of struct to the environment value
			newEnv.ClientID = value
		case "ClientSecret":
			value, ok := envVar.(string)
			// if environment value is not a string skip current field
			if !ok {
				continue
			}
			// else set field of struct to the environment value
			newEnv.ClientSecret = value
		case "RedirectUri":
			value, ok := envVar.(string)
			// if environment value is not a string skip current field
			if !ok {
				continue
			}
			// else set field of stsruct to the environment value
			newEnv.RedirectUri = value
		default:
			continue
		}
	}

	return &newEnv, nil
}

// NewEnv runs `BuildEnv` under the hood but forces the
// program to either return a functional `Env` struct pointer
// or crash otherwise
func NewEnv() *Env {
	// return fully built environment struct, crash on error
	return utils.Must(BuildEnv())
}

// LoadEnv reads a '.env' file and loads it into the 
// OS environment. Returns an error if it fails to 
// get current working directory.
func LoadEnv() error {
	// get current working dir
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	envPath := filepath.Join(dir, ".env")
	// load env file data into OS environment
	godotenv.Load(envPath)

	return nil
}
