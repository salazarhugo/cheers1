import {Order as OrderGen} from "../../../../gen/ts/cheers/order/v1/order_service";

export class Order {
    id: string
    name: string
    price: number
}

export function toOrder(value: OrderGen): Order {
    const order = new Order()
    Object.assign(order, value)
    return order
}
