MATCH
  (post:Post {id: $postID})
DETACH DELETE post
