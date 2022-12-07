package repository

import ticketpb "github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"

func CalculateAmount(
	tickets []*ticketpb.Ticket,
) (int64, error) {
	var amount int64 = 0

	for _, ticket := range tickets {
		amount += ticket.Price
	}

	return amount, nil
}
