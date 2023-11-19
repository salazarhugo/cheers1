package repository

func (p *partyRepository) UpdateMinimumPrice(
	price int64,
	partyId string,
) error {
	party, err := p.GetPartyById(partyId)
	if err != nil {
		return err
	}

	//if price >= party.MinimumPrice {
	//	return nil
	//}
	//
	//party.MinimumPrice = price
	_, err = p.UpdateParty(party)
	if err != nil {
		return err
	}

	return nil
}
