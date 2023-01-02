MATCH (post:Post {id: $postID})
SET
post.lastCommentText = $lastCommentText,
post.lastCommentUsername = $lastCommentUsername,
post.lastCommentCreateTime = $lastCommentCreateTime
