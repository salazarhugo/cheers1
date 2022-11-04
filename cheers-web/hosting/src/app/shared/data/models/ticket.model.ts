export class Ticket {
    id: string
    name: string
    price: number
    quantity: number
    currency: string = "EUR"
    description: string
    validated: boolean
}