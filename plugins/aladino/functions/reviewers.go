// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package plugins_aladino_functions

import (
	"github.com/reviewpad/reviewpad/v3/codehost/github/target"
	"github.com/reviewpad/reviewpad/v3/handler"
	"github.com/reviewpad/reviewpad/v3/lang/aladino"
)

func Reviewers() *aladino.BuiltInFunction {
	return &aladino.BuiltInFunction{
		Type:           aladino.BuildFunctionType([]aladino.Type{}, aladino.BuildArrayOfType(aladino.BuildStringType())),
		Code:           reviewersCode,
		SupportedKinds: []handler.TargetEntityKind{handler.PullRequest},
	}
}

func reviewersCode(e aladino.Env, _ []aladino.Value) (aladino.Value, error) {
	pullRequest := e.GetTarget().(*target.PullRequestTarget).PullRequest
	usersReviewers := pullRequest.RequestedReviewers
	teamReviewers := pullRequest.RequestedTeams
	totalReviewers := len(usersReviewers) + len(teamReviewers)
	reviewersLogin := make([]aladino.Value, totalReviewers)

	for i, userReviewer := range usersReviewers {
		reviewersLogin[i] = aladino.BuildStringValue(userReviewer.GetLogin())
	}

	for i, teamReviewer := range teamReviewers {
		reviewersLogin[i+len(usersReviewers)] = aladino.BuildStringValue(teamReviewer.GetSlug())
	}

	return aladino.BuildArrayValue(reviewersLogin), nil
}
