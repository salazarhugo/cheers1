import {Party} from "./party.model";
import {Ticket as TicketPb} from "../../../../gen/ts/cheers/ticket/v1/ticket";


export class Ticket {
    id: string
    name: string
    partyId: string
    partyName: string
    price: number
    quantity: number
    currency: string = "EUR"
    description: string
    validated: boolean
    partyStartTime: number;
    partyEndTime: number;
    locationName: string;
    organizer: string;
    paymentIntentId: string;
}

export function toTicket(value: TicketPb): Ticket {
    const ticket = new Ticket()
    Object.assign(ticket, value)
    return ticket
}