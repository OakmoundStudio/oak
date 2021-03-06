package render

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/oakmound/oak/v3/dlog"
	"github.com/oakmound/oak/v3/fileutil"
	"github.com/oakmound/oak/v3/oakerr"
)

var (
	regexpSingleNumber = regexp.MustCompile(`^\d+$`)
	regexpTwoNumbers   = regexp.MustCompile(`^\d+x\d+$`)
)

// BatchLoad loads subdirectories from the given base folder and imports all files,
// using alias rules to automatically determine the size of sprites and sheets in
// subfolders.
// A folder named 16x8 will have its images split into sheets where each sprite is
// 16x8, for example. 16 is a shorter way of writing 16x16.
// An alias.json file can be included that can indicate what dimensions named folders
// represent, so a "tiles": "32" field in the json would indicate that sprite sheets
// in the /tiles folder should be read as 32x32
func BatchLoad(baseFolder string) error {
	return BlankBatchLoad(baseFolder, 0)
}

// BlankBatchLoad acts like BatchLoad, but will not load and instead return a blank image
// of the appropriate dimensions for anything above maxFileSize.
func BlankBatchLoad(baseFolder string, maxFileSize int64) error {
	folders, err := fileutil.ReadDir(baseFolder)
	if err != nil {
		return err
	}
	aliases, err := parseAliasFile(baseFolder)
	if err != nil {
		return err
	}

	warnFiles := []string{}

	var wg sync.WaitGroup

	for _, folder := range folders {
		if folder.IsDir() {
			frameW, frameH, possibleSheet, err := parseLoadFolderName(aliases, folder.Name())
			if err != nil {
				return err
			}

			files, _ := fileutil.ReadDir(filepath.Join(baseFolder, folder.Name()))
			for _, file := range files {
				if !file.IsDir() {
					name := file.Name()
					if _, ok := fileDecoders[strings.ToLower(name[len(name)-4:])]; ok {
						lower := strings.ToLower(name)
						relativePath := filepath.Join(folder.Name(), name)
						if lower != name {
							warnFiles = append(warnFiles, relativePath)
						}
						wg.Add(1)
						go func(baseFolder, relativePath string, possibleSheet bool, frameW, frameH int) {
							defer wg.Done()
							buff, err := loadSprite(baseFolder, relativePath, maxFileSize)
							if err != nil {
								dlog.Error(err)
								return
							}

							if !possibleSheet {
								return
							}

							w := buff.Bounds().Max.X
							h := buff.Bounds().Max.Y

							if w < frameW || h < frameH {
								dlog.Error("File ", name, " in folder", folder.Name(),
									" is too small for folder dimensions", frameW, frameH)

								// Load this as a sheet if it is greater
								// than the folder size's frame size
							} else if w != frameW || h != frameH {
								_, err = LoadSheet(baseFolder, relativePath, frameW, frameH, defaultPad)
								dlog.ErrorCheck(err)
							}
						}(baseFolder, relativePath, possibleSheet, frameW, frameH)
					} else {
						dlog.Error("Unsupported file ending for batchLoad: ", name)
					}
				}
			}
		}
	}
	if len(warnFiles) != 0 {
		fileNames := strings.Join(warnFiles, ",")
		dlog.Info("The files", fileNames, "are not all lowercase. This may cause data to fail to load"+
			" when using tools like go-bindata.")
	}
	wg.Wait()
	return nil
}

func parseAliasFile(baseFolder string) (map[string]string, error) {
	aliasFile, err := fileutil.ReadFile(filepath.Join(baseFolder, "alias.json"))
	aliases := make(map[string]string)
	if err == nil {
		err = json.Unmarshal(aliasFile, &aliases)
		if err != nil {
			return nil, oakerr.InvalidInput{InputName: "alias.json"}
		}
	}
	return aliases, nil
}

func parseLoadFolderName(aliases map[string]string, name string) (int, int, bool, error) {
	var frameW, frameH int
	if result := regexpTwoNumbers.Find([]byte(name)); result != nil {
		vals := strings.Split(string(result), "x")
		frameW, _ = strconv.Atoi(vals[0])
		frameH, _ = strconv.Atoi(vals[1])
	} else if result := regexpSingleNumber.Find([]byte(name)); result != nil {
		val, _ := strconv.Atoi(string(result))
		frameW = val
		frameH = val
	} else {
		if aliased, ok := aliases[name]; ok {
			if result := regexpTwoNumbers.Find([]byte(aliased)); result != nil {
				vals := strings.Split(string(result), "x")
				frameW, _ = strconv.Atoi(vals[0])
				frameH, _ = strconv.Atoi(vals[1])
			} else if result := regexpSingleNumber.Find([]byte(aliased)); result != nil {
				val, _ := strconv.Atoi(string(result))
				frameW = val
				frameH = val
			} else {
				return 0, 0, false, errors.New("Alias value not parseable as a frame width and height pair")
			}
		} else {
			dlog.Info("Folder name", name, "parsed to 0x0 (unbound) dimensions.")
			frameW = 0
			frameH = 0
		}
	}
	return frameW, frameH, frameW != 0 && frameH != 0, nil
}
