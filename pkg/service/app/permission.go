// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Auto generated by 'go run gen_helper.go', DO NOT EDIT.

package app

import (
	"context"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/ctxutil"
)

func CheckAppsPermission(ctx context.Context, resourceIds []string) ([]*models.App, error) {
	if len(resourceIds) == 0 {
		return nil, nil
	}
	var sender = ctxutil.GetSender(ctx)
	var apps []*models.App
	_, err := pi.Global().DB(ctx).
		Select(models.AppColumns...).
		From(constants.TableApp).
		Where(db.Eq(constants.ColumnAppId, resourceIds)).Load(&apps)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil {
		for _, app := range apps {
			if !app.OwnerPath.CheckPermission(sender) && app.Owner != sender.UserId {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, app.AppId)
			}
		}
	}
	if len(apps) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceIds)
	}
	return apps, nil
}

func CheckAppPermission(ctx context.Context, resourceId string) (*models.App, error) {
	if len(resourceId) == 0 {
		return nil, nil
	}
	var sender = ctxutil.GetSender(ctx)
	var apps []*models.App
	_, err := pi.Global().DB(ctx).
		Select(models.AppColumns...).
		From(constants.TableApp).
		Where(db.Eq(constants.ColumnAppId, resourceId)).Load(&apps)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil {
		for _, app := range apps {
			if !app.OwnerPath.CheckPermission(sender) {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, app.AppId)
			}
		}
	}
	if len(apps) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceId)
	}
	return apps[0], nil
}

func CheckAppVersionsPermission(ctx context.Context, resourceIds []string) ([]*models.AppVersion, error) {
	if len(resourceIds) == 0 {
		return nil, nil
	}
	var sender = ctxutil.GetSender(ctx)
	var appversions []*models.AppVersion
	_, err := pi.Global().DB(ctx).
		Select(models.AppVersionColumns...).
		From(constants.TableAppVersion).
		Where(db.Eq(constants.ColumnVersionId, resourceIds)).Load(&appversions)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil {
		for _, appversion := range appversions {
			if !appversion.OwnerPath.CheckPermission(sender) && appversion.Owner != sender.UserId {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, appversion.VersionId)
			}
		}
	}
	if len(appversions) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceIds)
	}
	return appversions, nil
}

func CheckAppVersionPermission(ctx context.Context, resourceId string) (*models.AppVersion, error) {
	if len(resourceId) == 0 {
		return nil, nil
	}
	var sender = ctxutil.GetSender(ctx)
	var appversions []*models.AppVersion
	_, err := pi.Global().DB(ctx).
		Select(models.AppVersionColumns...).
		From(constants.TableAppVersion).
		Where(db.Eq(constants.ColumnVersionId, resourceId)).Load(&appversions)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil {
		for _, appversion := range appversions {
			if !appversion.OwnerPath.CheckPermission(sender) {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, appversion.VersionId)
			}
		}
	}
	if len(appversions) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceId)
	}
	return appversions[0], nil
}