package target

// takes an target and returns the target dto
func toTargetDto(target *target) targetDto {
	return targetDto{
		ID:           target.ID,
		Name:         target.Name,
		Description:  target.Description,
		TargetTypeID: target.TargetTypeID,
		AccountID:    target.AccountID,
		CreatedAt:    target.CreatedAt,
	}
}

func toRtmpDto(rtmp *rtmp) rtmpDto {
	return rtmpDto{
		TargetID:  rtmp.TargetID,
		URL:       rtmp.URL,
		StreamKey: rtmp.StreamKey,
		CreatedAt: rtmp.CreatedAt,
	}
}
