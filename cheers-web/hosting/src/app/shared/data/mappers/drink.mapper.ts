import {Drink} from "../../../../gen/ts/cheers/drink/v1/drink_service";
import {DrinkModel} from "../models/drink.model";


export function toDrinkModel(d: Drink): DrinkModel {
    const res = new DrinkModel()
    res.id = d.id
    res.name = d.name
    res.icon = d.icon

    return res
}