// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"errors"
	"strings"

	"github.com/iotaledger/wasp/tools/schema/model"
)

const (
	KeyAccess      string = "access"
	KeyAuthor      string = "author"
	KeyCopyright   string = "copyright"
	KeyDescription string = "description"
	KeyEvents      string = "events"
	KeyFuncs       string = "funcs"
	KeyLicense     string = "license"
	KeyName        string = "name"
	KeyParams      string = "params"
	KeyRepository  string = "repository"
	KeyResults     string = "results"
	KeyState       string = "state"
	KeyStructs     string = "structs"
	KeyTypedefs    string = "typedefs"
	KeyVersion     string = "version"
	KeyViews       string = "views"
)

//nolint:gocyclo
func Convert(root *Node, def *model.SchemaDef) error {
	for _, key := range root.Contents {
		switch key.Val {
		case KeyCopyright:
			copyright := key.toStringElt()
			if copyright.Val != "" {
				def.Copyright = "// " + copyright.Val + "\n"
			} else {
				def.Copyright = key.HeadComment
			}
		case KeyName:
			def.Name = key.toStringElt()
		case KeyDescription:
			def.Description = key.toStringElt()
		case KeyAuthor:
			def.Author = key.toStringElt()
		case KeyLicense:
			def.License = key.toStringElt()
		case KeyRepository:
			def.Repository = key.toStringElt()
		case KeyVersion:
			def.Version = key.toStringElt()
		case KeyEvents:
			def.Events = key.toDefMapMap()
		case KeyStructs:
			def.Structs = key.toDefMapMap()
		case KeyTypedefs:
			def.Typedefs = key.toDefMap()
		case KeyState:
			def.State = key.toDefMap()
		case KeyFuncs:
			def.Funcs = key.toFuncDefMap()
		case KeyViews:
			def.Views = key.toFuncDefMap()
		default:
			return errors.New("unsupported key")
		}
	}
	return nil
}

func (n *Node) toStringElt() model.DefElt {
	var result model.DefElt
	if len(n.Contents) != 0 {
		result.Val = strings.TrimSpace(n.Contents[0].Val)
	}
	result.Line = n.Line
	return result
}

func (n *Node) toDefElt() *model.DefElt {
	comment := ""
	if len(n.HeadComment) > 0 {
		// remove trailing '\n'
		comment = n.HeadComment[:len(n.HeadComment)-1]
	} else if len(n.LineComment) > 0 {
		// remove trailing '\n'
		comment = n.LineComment[:len(n.LineComment)-1]
	}
	return &model.DefElt{
		Val:     n.Val,
		Comment: comment,
		Line:    n.Line,
	}
}

func (n *Node) toDefMap() model.DefMap {
	defs := make(model.DefMap)
	for _, yamlKey := range n.Contents {
		if strings.ReplaceAll(yamlKey.Val, " ", "") == "{}" {
			// treat "{}" as empty
			continue
		}
		key := *yamlKey.toDefElt()
		val := yamlKey.Contents[0].toDefElt()
		if val.Comment != "" && key.Comment == "" {
			key.Comment = val.Comment
		}
		val.Comment = ""
		defs[key] = val
	}
	return defs
}

func (n *Node) toDefMapMap() model.DefMapMap {
	defs := make(model.DefMapMap)
	for _, yamlKey := range n.Contents {
		// TODO better parsing
		if strings.ReplaceAll(yamlKey.Val, " ", "") == "{}" {
			// treat "{}" as empty
			continue
		}
		comment := ""
		if len(yamlKey.HeadComment) > 0 {
			comment = yamlKey.HeadComment[:len(yamlKey.HeadComment)-1] // remove trailing '\n'
		} else if len(yamlKey.LineComment) > 0 {
			comment = yamlKey.LineComment[:len(yamlKey.LineComment)-1] // remove trailing '\n'
		}

		key := model.DefElt{
			Val:     yamlKey.Val,
			Comment: comment,
			Line:    yamlKey.Line,
		}
		val := yamlKey.toDefMap()
		defs[key] = &val
	}
	return defs
}

func (n *Node) toFuncDef() model.FuncDef {
	def := model.FuncDef{}
	def.Line = n.Line
	if len(n.HeadComment) > 0 {
		def.Comment = n.HeadComment[:len(n.HeadComment)-1] // remove trailing '\n'
	} else if len(n.LineComment) > 0 {
		def.Comment = n.LineComment[:len(n.LineComment)-1] // remove trailing '\n'
	}

	for _, yamlKey := range n.Contents {
		if strings.ReplaceAll(yamlKey.Val, " ", "") == "{}" {
			// treat "{}" as empty
			continue
		}
		switch yamlKey.Val {
		case KeyAccess:
			def.Access = *yamlKey.Contents[0].toDefElt()
			if len(yamlKey.HeadComment) > 0 {
				def.Access.Comment = yamlKey.HeadComment[:len(yamlKey.HeadComment)-1] // remove trailing '\n'
			} else if len(yamlKey.LineComment) > 0 {
				def.Access.Comment = yamlKey.LineComment[:len(yamlKey.LineComment)-1] // remove trailing '\n'
			}
		case KeyParams:
			def.Params = yamlKey.toDefMap()
		case KeyResults:
			def.Results = yamlKey.toDefMap()
		default:
			return model.FuncDef{}
		}
	}
	return def
}

func (n *Node) toFuncDefMap() model.FuncDefMap {
	defs := make(model.FuncDefMap)
	for _, yamlKey := range n.Contents {
		if strings.ReplaceAll(yamlKey.Val, " ", "") == "{}" {
			// treat "{}" as empty
			continue
		}
		comment := ""
		if len(yamlKey.HeadComment) > 0 {
			comment = yamlKey.HeadComment[:len(yamlKey.HeadComment)-1] // remove trailing '\n'
		} else if len(yamlKey.LineComment) > 0 {
			comment = yamlKey.LineComment[:len(yamlKey.LineComment)-1] // remove trailing '\n'
		}
		key := model.DefElt{
			Val:     yamlKey.Val,
			Comment: comment,
			Line:    yamlKey.Line,
		}
		val := yamlKey.toFuncDef()
		defs[key] = &val
	}
	return defs
}
