package response

import (
	"api-gateway/app/domain/usercases/record/body/repo"
	"api-gateway/app/util"
)

func NewBodyRecordResponse(bodyRecord repo.UserBodyRecordRepo) UserBodyRecordResponse {
	return UserBodyRecordResponse{
		Id:         bodyRecord.Id,
		Percentage: bodyRecord.Percentage,
		Date:       util.FormatDateTime(*bodyRecord.CreatedAt),
	}
}

func NewBodyRecordResponses(bodyRecords []repo.UserBodyRecordRepo) []UserBodyRecordResponse {
	if len(bodyRecords) <= 0 {
		return []UserBodyRecordResponse{}
	}
	resp := make([]UserBodyRecordResponse, len(bodyRecords))
	for i := range bodyRecords {
		resp[i] = NewBodyRecordResponse(bodyRecords[i])
	}
	return resp
}
