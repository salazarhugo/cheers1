MATCH
  (u:User {id: $userID})-[liked:LIKED]->(post:Post {id: $postID})
DELETE
	liked