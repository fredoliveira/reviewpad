// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package plugins_aladino_functions

import (
	"github.com/reviewpad/reviewpad/v3/handler"
	"github.com/reviewpad/reviewpad/v3/lang/aladino"
)

func Description() *aladino.BuiltInFunction {
	return &aladino.BuiltInFunction{
		Type:           aladino.BuildFunctionType([]aladino.Type{}, aladino.BuildStringType()),
		Code:           descriptionCode,
		SupportedKinds: []handler.TargetEntityKind{handler.PullRequest, handler.Issue},
	}
}

func descriptionCode(e aladino.Env, args []aladino.Value) (aladino.Value, error) {
	description, err := e.GetTarget().GetDescription()
	if err != nil {
		return nil, err
	}

	return aladino.BuildStringValue(description), nil
}
