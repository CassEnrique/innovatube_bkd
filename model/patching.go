/*
@Author: eecass
@Date:
@Last Modified by:   eecass
@Last Modified time:
*/

package model

type Patching struct {
	Pk    int    `json:"pk,omitempty"`
	Key   string `json:"key,omitempty"`
	Value any    `json:"value,omitempty"`
}