MATCH
  (u:User {id: $userID})-[liked:LIKED]->(story:Story {id: $storyID})
DELETE
	liked