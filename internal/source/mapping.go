package source

// takes an source and returns the source dto
func toSourceDto(source *source) sourceDto {
	return sourceDto{
		UUID:         source.UUID,
		Name:         source.Name,
		Description:  source.Description,
		SourceTypeID: source.SourceTypeID,
		AccountID:    source.AccountID,
		CreatedAt:    source.CreatedAt,
	}
}

func toRtmpDto(rtmp *rtmp) rtmpDto {
	return rtmpDto{
		SourceID:  rtmp.SourceID,
		URL:       rtmp.URL,
		StreamKey: rtmp.StreamKey,
		CreatedAt: rtmp.CreatedAt,
	}
}
