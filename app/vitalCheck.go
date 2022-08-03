package app

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func vitalCheck(dir string) (string, string, error) {
	const MERGED_NAME string = "merged.tsv"
	const CUR_POST_NAME string = "curated-postprocessed.tsv"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", "", err
	}

	merged := ""
	cur_post := ""
	for _, file := range files {
		if file.Name() == MERGED_NAME {
			merged = filepath.Join(dir, MERGED_NAME)
			if cur_post != "" {
				return merged, cur_post, nil
			}
		} else if file.Name() == CUR_POST_NAME {
			cur_post = filepath.Join(dir, CUR_POST_NAME)
			if merged != "" {
				return merged, cur_post, nil
			}
		}
	}
	var empty_err error
	if merged == "" && cur_post == "" {
		empty_err = fmt.Errorf("could not find files: %s and %s", MERGED_NAME, CUR_POST_NAME)
	} else if merged == "" {
		empty_err = fmt.Errorf("could not find file: %s", MERGED_NAME)
	} else {
		empty_err = fmt.Errorf("could not find file: %s", CUR_POST_NAME)
	}

	return merged, cur_post, empty_err
}
