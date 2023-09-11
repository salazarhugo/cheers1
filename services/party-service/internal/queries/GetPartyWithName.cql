MATCH (party:Party {name: $name})
RETURN PROPERTIES(party)
