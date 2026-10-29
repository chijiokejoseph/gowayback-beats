package env

import (
	"errors"
	"fmt"
	"os"
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

func BuildEnv() (*Env, error) {
	newEnv := Env{}
	envType := reflect.TypeOf(newEnv)
	numFields := envType.NumField()
	envVars := make([]any, numFields)
	for i := range numFields {
		envField := envType.Field(i)
		fieldTag := envField.Tag
		value := fieldTag.Get("env")
		if value == "" {
			envVars[i] = nil
			continue
		}
		if !strings.Contains(value, ",") {
			envVar, ok := os.LookupEnv(value)
			if !ok {
				envVars[i] = nil
				continue
			}
			envVars[i] = envVar
		}
		strPartitions := strings.Split(value, ",")
		value = strPartitions[0]
		keywords := strPartitions[1:len(strPartitions):len(strPartitions)]
		if slices.Contains(keywords, "required") {
			envVar, ok := os.LookupEnv(value)
			if !ok {
				msg := fmt.Sprintf("Struct Field %s of Env has tag %s with the required attribute. \nBut the corresponding environment variable for %s was not found.", envField.Name, fieldTag, value)
				return &newEnv, errors.New(msg)
			}
			envVars[i] = envVar
		}
	}

	for i, envVar := range envVars {
		field := envType.Field(i)
		switch field.Name {
		case "ClientID":
			value, ok := envVar.(string)
			if !ok {
				continue
			}
			newEnv.ClientID = value
		case "ClientSecret":
			value, ok := envVar.(string)
			if !ok {
				continue
			}
			newEnv.ClientSecret = value
		case "RedirectUri":
			value, ok := envVar.(string)
			if !ok {
				continue
			}
			newEnv.RedirectUri = value
		default:
			continue
		}
	}
	return &newEnv, nil
}

func NewEnv() *Env {
	return utils.Must(BuildEnv())
}

func LoadEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	envPath := fmt.Sprintf("%s/%s", dir, ".env")
	godotenv.Load(envPath)
	return nil
}
