package profile

import (
	"fmt"
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

func TestWriteProfileValid(t *testing.T) {
	profile := Profile{
		Fields: map[ID]Point{
			0x33: {
				Torque: 0x30,
				AD:     0x30,
			},
		},
	}
	result, err := profile.MarshalCSV()
	assert.NoError(t, err)
	assert.Equal(t, "ID,Torque,TorqueAD\n51,48,48\n", string(result))
}

func TestSaveProfile(t *testing.T) {
	fs := afero.NewMemMapFs()
	fileName := "testprofile.csv"
	profile := &Profile{
		Fields: map[ID]Point{
			0x33: {
				Torque: 0x30,
				AD:     0x30,
			},
		},
	}
	err := SaveProfile(profile, fileName, fs)
	assert.NoError(t, err)
	readBack, err := afero.ReadFile(fs, fileName)
	assert.NoError(t, err)
	assert.Equal(t, "ID,Torque,TorqueAD\n51,48,48\n", string(readBack))
}

func TestValidateProfileTooMany(t *testing.T) {
	profile := &Profile{}
	profile.Fields = make(map[ID]Point)
	for i := 0; i < 100; i++ {
		profile.Fields[ID(i)] = Point{
			AD:     0x30,
			Torque: 0x33,
		}
	}
	err := profile.Validate()
	assert.EqualError(t, err, fmt.Sprintf("profile should only have %d fields", profileLen))
}

func TestValidateProfileInvalidID(t *testing.T) {
	profile := &Profile{
		Fields: map[ID]Point{
			0xFF: {
				Torque: 0x30,
				AD:     0x30,
			},
		},
	}
	err := profile.Validate()
	assert.EqualError(t, err, fmt.Sprintf("invalid profile parameter: 255 should be between %d and %d", idMin, idMax))
}

func TestValidateProfileValid(t *testing.T) {
	profile := &Profile{
		Fields: map[ID]Point{
			0x00: {
				Torque: 0x30,
				AD:     0x30,
			},
		},
	}
	err := profile.Validate()
	assert.NoError(t, err)
}

func TestMarshalBytesOrdered(t *testing.T) {
	profile := &Profile{
		Fields: map[ID]Point{
			0x01: {
				Torque: 0x0030,
				AD:     0x00EF,
			},
			0x02: {
				Torque: 0xFFFF,
				AD:     0xFF00,
			},
		},
	}
	expected := []uint16{0x0030, 0xFFFF, 0x0030, 0xFF00}
	result := profile.MarshalBytes()
	assert.Equal(t, expected, result[0:4])
}

func TestMarshalBytesUnOrdered(t *testing.T) {
	profile := &Profile{
		Fields: map[ID]Point{
			0x08: {
				Torque: 0x0030,
				AD:     0x0030,
			},
			0x01: {
				Torque: 0xFFFF,
				AD:     0xFF00,
			},
		},
	}
	expected := []uint16{0x0030, 0xFFFF, 0x0030, 0xFF00}
	result := profile.MarshalBytes()
	assert.Equal(t, expected, result)
}
