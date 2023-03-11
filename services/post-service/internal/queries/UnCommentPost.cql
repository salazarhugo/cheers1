MATCH
  (u:User {id: $userID})-[commented:COMMENTED]->(post:Post {id: $postID})
DELETE
	commented