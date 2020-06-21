package profile

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/afero"
)

const idMax = 31
const idMin = 0
const profileLen = 24

type ID byte

// Point represents a torque sensor calibration point for the drill
type Point struct {
	AD     uint16
	Torque uint16
}

func SaveProfile(p *Profile, fileName string, fs afero.Fs) error {
	csv, err := p.MarshalCSV()
	if err != nil {
		return err
	}
	err = afero.WriteFile(fs, fileName, csv, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Profile contains a set of calibration points for the drill
type Profile struct {
	Fields map[ID]Point
}

func (p *Profile) Validate() error {
	if len(p.Fields) > profileLen {
		return fmt.Errorf("profile should only have %d fields", profileLen)
	}
	for id := range p.Fields {
		if id < idMin || id > idMax {
			return fmt.Errorf("invalid profile parameter: %d should be between %d and %d", id, idMin, idMax)
		}
	}
	return nil
}

func (p *Profile) MarshalBytes() [24 * 2]uint16 {
	keys := make([]ID, len(p.Fields))
	for k := range p.Fields {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i int, j int) bool { return keys[i] < keys[j] })
	var output [24 * 2]uint16
	for k := range keys {
		output[int(k)] = p.Fields[ID(k)].Torque
		output[int(k)+24] = p.Fields[ID(k)].AD
	}
	return output
}

func (p *Profile) MarshalCSV() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	err := writer.Write(WriteHeader())
	if err != nil {
		return nil, err
	}

	for id, point := range p.Fields {
		err := writer.Write(WriteRow(id, point))
		if err != nil {
			return nil, err
		}
	}
	writer.Flush()
	return buf.Bytes(), nil
}

func WriteHeader() []string {
	return []string{FieldID, Torque, AD}
}

func WriteRow(id ID, point Point) []string {
	return []string{
		strconv.Itoa(int(id)),
		strconv.Itoa(int(point.AD)),
		strconv.Itoa(int(point.Torque)),
	}
}

// LoadProfile takes an input file path and returns a Profile struct
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
	var conv []uint16
	for _, val := range row {
		result, err := strconv.Atoi(val)
		if err != nil {
			return ID(0x00), Point{}, fmt.Errorf("error parsing profile: %w", err)
		}
		conv = append(conv, uint16(result))
	}

	return ID(conv[0]), MakePoint(conv[1:]), nil
}

func MakePoint(row []uint16) Point {
	return Point{
		AD:     row[0],
		Torque: row[1],
	}
}
