package utils

func GetLimitAndOffsetPagination(
	page int,
	pageSize int,
) (int, int) {
	if page == 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	skip := pageSize * (page - 1)

	return pageSize, skip
}
