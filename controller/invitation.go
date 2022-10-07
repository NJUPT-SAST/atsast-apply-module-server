package controller

import (
	"github.com/njupt-sast/atsast-apply-module-server/common/jwt"
	"github.com/njupt-sast/atsast-apply-module-server/controller/stdrsp"
	"github.com/njupt-sast/atsast-apply-module-server/dao"
	"github.com/njupt-sast/atsast-apply-module-server/entity"
	"github.com/njupt-sast/atsast-apply-module-server/service"
)

func readInvitation(identity *jwt.Identity, code *string) *stdrsp.Result {
	invitation, err := service.ReadInvitation(code)
	if err == service.DocumentNotFoundError {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.NotFoundErrMsg)
	}
	if err != nil {
		return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
	}

	if invitation.UserId == nil {
		err = dao.UpdateInvitationUserId(code, identity.Uid)
		if err != nil {
			return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
		}

		if *invitation.Type == "admin" {
			if err = service.UpdateUserRole(identity.Uid, &entity.AdminUserRole); err == service.DocumentNotFoundError {
				return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.NotFoundErrMsg)
			} else if err != nil {
				return stdrsp.Response(stdrsp.Failed).Msg(stdrsp.CallDatabaseErrMsg)
			}
		}
	} else if *invitation.UserId != *identity.Uid {
		return stdrsp.Response(stdrsp.BadRequest).Msg(stdrsp.PermissionDeniedErrMsg)
	}

	return stdrsp.Response(stdrsp.Success).Data(invitation)
}
