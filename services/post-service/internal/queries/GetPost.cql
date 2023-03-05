MATCH
  (post:Post {id: $postID})
RETURN
    PROPERTIES(post)