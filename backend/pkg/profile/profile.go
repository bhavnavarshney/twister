package profile

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/afero"
)

type ID byte

// Point represents a torque sensor calibration point for the drill
type Point struct {
	AD     byte
	Torque byte
}

// Profile contains a set of calibration points for the drill
type Profile struct {
	Fields map[ID]Point
}

func LoadProfile(filepath string, fs afero.Fs) (*Profile, error) {
	file, err := afero.ReadFile(fs, filepath)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(strings.NewReader(string(file)))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	p := Profile{Fields: make(map[ID]Point)}

	for i, row := range rows {
		if i == 0 {
			err := ParseHeader(row)
			if err != nil {
				return nil, err
			}
			continue
		}
		id, point, err := ParseRow(row)
		if err != nil {
			return nil, err
		}
		p.Fields[id] = point
	}

	return &p, nil
}

const (
	FieldID = "ID"
	Torque  = "Torque"
	AD      = "TorqueAD"
)

// TODO: Add dynamic row ordering
// ParseHeader checks that the header row has 3 values and each value is as expected
func ParseHeader(row []string) error {
	if len(row) != 3 {
		return errors.New("header expected to be 3 items")
	}
	if row[0] != FieldID || row[1] != Torque || row[2] != AD {
		return errors.New("header expected to be " + strings.Join(BuildHeader(), ","))
	}
	return nil
}

func BuildHeader() []string {
	return []string{FieldID, Torque, AD}
}

func ParseRow(row []string) (ID, Point, error) {
	var conv []byte
	for _, val := range row {
		result, err := strconv.Atoi(val)
		if err != nil {
			return ID(0x00), Point{}, fmt.Errorf("error parsing profile: %w", err)
		}
		conv = append(conv, byte(result))
	}

	return ID(conv[0]), MakePoint(conv[1:]), nil
}

func MakePoint(row []byte) Point {
	return Point{
		AD:     row[0],
		Torque: row[1],
	}
}
