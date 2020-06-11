package profile

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestLoadValidProfile(t *testing.T) {
	fs := afero.NewOsFs()
	profile, err := LoadProfile("../../config/default.csv", fs)
	assert.NoError(t, err)
	t.Log(profile)
}

func TestLoadNonExistentProfile(t *testing.T) {
	fs := afero.NewOsFs()
	profile, err := LoadProfile("nonexistentfile.csv", fs)
	assert.EqualError(t, err, "open nonexistentfile.csv: no such file or directory")
	t.Log(profile)
}

func TestLoadProfileUnevenRows(t *testing.T) {
	fileName := "test/invalid.csv"
	unevenFile := []byte(
		`a,b,c
	1,2,3,4
	`)
	fs := afero.NewMemMapFs()
	err := afero.WriteFile(fs, fileName, unevenFile, os.ModePerm)
	assert.NoError(t, err)
	profile, err := LoadProfile(fileName, fs)
	assert.EqualError(t, err, "record on line 2: wrong number of fields")
	assert.Nil(t, profile)
}

func TestLoadProfileNoHeader(t *testing.T) {
	fileName := "test/noheader.csv"
	noHeaderFile := []byte(
		`1,2,3`)
	fs := afero.NewMemMapFs()
	err := afero.WriteFile(fs, fileName, noHeaderFile, os.ModePerm)
	assert.NoError(t, err)
	profile, err := LoadProfile(fileName, fs)
	assert.EqualError(t, err, "header expected to be ID,Torque,TorqueAD")
	assert.Nil(t, profile)
}

func TestLoadProfileInvalidHeader(t *testing.T) {
	fileName := "test/noheader.csv"
	noHeaderFile := []byte(
		`a,b,c
		1,2,3`)
	fs := afero.NewMemMapFs()
	err := afero.WriteFile(fs, fileName, noHeaderFile, os.ModePerm)
	assert.NoError(t, err)
	profile, err := LoadProfile(fileName, fs)
	assert.EqualError(t, err, "header expected to be ID,Torque,TorqueAD")
	assert.Nil(t, profile)
}

func TestLoadProfileWrongHeaderLength(t *testing.T) {
	fileName := "test/noheader.csv"
	noHeaderFile := []byte(
		`ID,Torque,TorqueAD,Extra
		1,2,3,4`)
	fs := afero.NewMemMapFs()
	err := afero.WriteFile(fs, fileName, noHeaderFile, os.ModePerm)
	assert.NoError(t, err)
	profile, err := LoadProfile(fileName, fs)
	assert.EqualError(t, err, "header expected to be 3 items")
	assert.Nil(t, profile)
}

func TestLoadProfileInvalidChar(t *testing.T) {
	fileName := "test/noheader.csv"
	noHeaderFile := []byte(
		`ID,Torque,TorqueAD
		1,2,b`)
	fs := afero.NewMemMapFs()
	err := afero.WriteFile(fs, fileName, noHeaderFile, os.ModePerm)
	assert.NoError(t, err)
	profile, err := LoadProfile(fileName, fs)
	assert.EqualError(t, err, "error parsing profile: strconv.Atoi: parsing \"\\t\\t1\": invalid syntax")
	assert.Nil(t, profile)
}
