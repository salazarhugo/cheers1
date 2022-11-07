MATCH (party:Party {id: $partyId})
RETURN PROPERTIES(party)
