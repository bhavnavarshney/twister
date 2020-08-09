// +build unit

package profile

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestLoadValidProfile(t *testing.T) {
	fs := afero.NewOsFs()
	file, err := afero.ReadFile(fs, "../../config/default.csv")
	profile, err := LoadProfile(string(file))
	assert.NoError(t, err)
	assert.Equal(t, uint16(60), profile.Fields[0].AD)
	assert.Equal(t, uint16(4), profile.Fields[0].Torque)
	assert.Equal(t, uint16(94), profile.Fields[23].AD)
	assert.Equal(t, uint16(7), profile.Fields[23].Torque)
}

func TestLoadProfileUnevenRows(t *testing.T) {
	unevenFile := `a,b,c
	1,2,3,4
	`
	profile, err := LoadProfile(unevenFile)
	assert.EqualError(t, err, "record on line 2: wrong number of fields")
	assert.Nil(t, profile)
}

func TestLoadProfileNoHeader(t *testing.T) {
	noHeaderFile := `1,2,3`
	profile, err := LoadProfile(noHeaderFile)
	assert.EqualError(t, err, "header expected to be ID,Torque,TorqueAD")
	assert.Nil(t, profile)
}

func TestLoadProfileInvalidHeader(t *testing.T) {
	invalidHeaderFile := `a,b,c
		1,2,3`
	profile, err := LoadProfile(invalidHeaderFile)
	assert.EqualError(t, err, "header expected to be ID,Torque,TorqueAD")
	assert.Nil(t, profile)
}

func TestLoadProfileWrongHeaderLength(t *testing.T) {
	noHeaderFile :=
		`ID,Torque,TorqueAD,Extra
		1,2,3,4`
	profile, err := LoadProfile(noHeaderFile)
	assert.EqualError(t, err, "header expected to be 3 items")
	assert.Nil(t, profile)
}

func TestLoadProfileInvalidChar(t *testing.T) {
	noHeaderFile := `ID,Torque,TorqueAD
		1,2,b`
	profile, err := LoadProfile(noHeaderFile)
	assert.EqualError(t, err, "error parsing profile: strconv.Atoi: parsing \"\\t\\t1\": invalid syntax")
	assert.Nil(t, profile)
}

func TestWriteProfileValid(t *testing.T) {
	profile := Profile{
		Fields: [24]Point{
			0: {
				Torque: 0x30,
				AD:     0x30,
			},
		},
	}
	result, err := profile.MarshalCSV()
	assert.NoError(t, err)
	assert.Equal(t, "ID,Torque,TorqueAD\n1,48,48\n2,0,0\n3,0,0\n4,0,0\n5,0,0\n6,0,0\n7,0,0\n8,0,0\n9,0,0\n10,0,0\n11,0,0\n12,0,0\n13,0,0\n14,0,0\n15,0,0\n16,0,0\n17,0,0\n18,0,0\n19,0,0\n20,0,0\n21,0,0\n22,0,0\n23,0,0\n24,0,0\n", string(result))
}

func TestSaveProfile(t *testing.T) {
	fs := afero.NewMemMapFs()
	fileName := "testprofile.csv"
	profile := &Profile{
		Fields: [24]Point{
			0: {
				Torque: 0x30,
				AD:     0x30,
			},
		},
	}
	err := SaveProfile(profile, fileName, fs)
	assert.NoError(t, err)
	readBack, err := afero.ReadFile(fs, fileName)
	assert.NoError(t, err)
	assert.Equal(t, "ID,Torque,TorqueAD\n1,48,48\n2,0,0\n3,0,0\n4,0,0\n5,0,0\n6,0,0\n7,0,0\n8,0,0\n9,0,0\n10,0,0\n11,0,0\n12,0,0\n13,0,0\n14,0,0\n15,0,0\n16,0,0\n17,0,0\n18,0,0\n19,0,0\n20,0,0\n21,0,0\n22,0,0\n23,0,0\n24,0,0\n", string(readBack))
}

func TestValidateProfileValid(t *testing.T) {
	profile := &Profile{
		Fields: [24]Point{
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
		Fields: [24]Point{
			0: {
				Torque: 0x0030,
				AD:     0x00EF,
			},
			1: {
				Torque: 0xFFFF,
				AD:     0xFF00,
			},
		},
	}
	expected := [48]uint16{0: 0x0030, 1: 0xFFFF, 24: 0x00EF, 25: 0xFF00}
	result := profile.MarshalBytes()
	assert.Equal(t, expected, result)
}

func TestMarshalRow(t *testing.T) {
	fs := afero.NewOsFs()
	file, err := afero.ReadFile(fs, "../../config/default.csv")
	profile, err := LoadProfile(string(file))
	assert.NoError(t, err)
	result := profile.MarshalRow()
	expected := []string{"4", "60", "5", "72", "6", "83", "7", "94", "8", "108", "9", "122", "10", "131", "11", "153", "12", "168", "13", "180", "14", "191", "30", "30", "30", "30", "30", "30", "30", "30", "30", "30", "30", "30", "30", "30", "30", "30", "3", "50", "4", "60", "5", "72", "6", "83", "7", "94"}
	assert.Equal(t, expected, result)
}
