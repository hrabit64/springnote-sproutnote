package config

import (
	"fmt"
	"github.com/hrabit64/sproutnote/pkg/utils/regex"
	"github.com/spf13/viper"
	"regexp"
)

type Env struct {
	BackupPath           string `mapstructure:"BACKUP_PATH"`
	FileBackupTime       string `mapstructure:"FILE_BACKUP_TIME"`
	MaxFileBackupHistory int    `mapstructure:"MAX_FILE_BACKUP_HISTORY"`
	DbBackupTime         string `mapstructure:"DB_BACKUP_TIME"`
	MaxDbBackupHistory   int    `mapstructure:"MAX_DB_BACKUP_HISTORY"`
	DbItemSecret         string `mapstructure:"DB_ITEM_SECRET"`
}

var ProcessType string

var RootEnv *Env

func LoadEnv() (*Env, error) {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		return nil, err
	}

	return &env, nil
}

func ValidateEnv(env *Env) (bool, error) {
	valid, err := ValidateDbBackupTime(env.DbBackupTime)
	if err != nil {
		return false, err
	}

	if !valid {
		fmt.Println("Invalid DB backup time! Please check the time format.")
		return false, nil
	}

	valid, err = ValidateFileBackupTime(env.FileBackupTime)
	if err != nil {
		return false, err
	}

	if !valid {
		fmt.Println("Invalid File backup time! Please check the time format.")
		return false, nil
	}

	valid, err = ValidateMaxDbBackupHistory(env.MaxDbBackupHistory)
	if err != nil {
		return false, err
	}

	if !valid {
		fmt.Println("Invalid Max DB backup history! Please check the number.")
		return false, nil
	}

	valid, err = ValidateMaxFileBackupHistory(env.MaxFileBackupHistory)
	if err != nil {
		return false, err
	}

	if !valid {
		fmt.Println("Invalid Max File backup history! Please check the number.")
		return false, nil
	}

	return true, nil
}

func ValidateFileBackupTime(time string) (bool, error) {
	return regexp.MatchString(regex.DATE_REGEX, time)
}

func ValidateDbBackupTime(time string) (bool, error) {
	return regexp.MatchString(regex.DATE_REGEX, time)
}

func ValidateMaxFileBackupHistory(max int) (bool, error) {
	return max > 0 && max < 100, nil
}

func ValidateMaxDbBackupHistory(max int) (bool, error) {
	return max > 0 && max < 100, nil
}

func RewriteEnv(env *Env) error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.Set("BACKUP_PATH", env.BackupPath)
	viper.Set("FILE_BACKUP_TIME", env.FileBackupTime)
	viper.Set("MAX_FILE_BACKUP_HISTORY", env.MaxFileBackupHistory)
	viper.Set("DB_BACKUP_TIME", env.DbBackupTime)
	viper.Set("MAX_DB_BACKUP_HISTORY", env.MaxDbBackupHistory)
	viper.Set("DB_ITEM_SECRET", env.DbItemSecret)

	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}
